package imperial

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/efritz/backoff"
	"github.com/efritz/glock"
	"github.com/efritz/imperial/proto"
	"github.com/efritz/watchdog"
	pb "github.com/golang/protobuf/proto"
)

type (
	RiemannReporter struct {
		logger   Logger
		clock    glock.Clock
		dialer   RiemannDialer
		configs  []ConfigFunc
		ttl      float32
		conn     io.ReadWriteCloser
		events   chan *riemannEvent
		messages chan []byte
		done     chan struct{}
		once     *sync.Once
		wg       *sync.WaitGroup
	}

	RiemannDialer func() (io.ReadWriteCloser, error)

	riemannEvent struct {
		service    string
		metric     int64
		time       int64
		attributes map[string]string
	}
)

func NewRiemannReporter(addr string, configs ...RiemannConfigFunc) *RiemannReporter {
	config := newRiemannConfig()
	for _, f := range configs {
		f(config)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	reporter := &RiemannReporter{
		dialer:   makeDialer(addr, config),
		logger:   config.logger,
		clock:    config.clock,
		configs:  config.configs,
		ttl:      config.ttl,
		events:   make(chan *riemannEvent),
		messages: make(chan []byte, config.queueSize),
		done:     make(chan struct{}),
		once:     &sync.Once{},
		wg:       wg,
	}

	go reporter.batch(config.batchSize, config.tickDuration)
	go reporter.publish()

	return reporter
}

func (r *RiemannReporter) Report(name string, value int, configs ...ConfigFunc) {
	options := applyConfigs(r.configs, configs)

	r.events <- &riemannEvent{
		service:    name,
		metric:     int64(value),
		time:       r.clock.Now().Unix(),
		attributes: options.attributes,
	}
}

func (r *RiemannReporter) Shutdown() {
	r.once.Do(func() {
		close(r.done)
		close(r.events)
	})

	r.wg.Wait()
}

func (r *RiemannReporter) batch(batchSize int, tickDuration time.Duration) {
	defer close(r.messages)

	var (
		hostname, _ = os.Hostname()
		ticker      = r.clock.NewTicker(tickDuration)
		batch       = make([]*proto.Event, 0, batchSize)
	)

	sendBatch := func() {
		if len(batch) == 0 {
			return
		}

		serialized, err := pb.Marshal(&proto.Msg{Events: batch})
		if err != nil {
			r.logger.Printf(
				"Failed to serialize Riemann message (%s)",
				err.Error(),
			)
		}

		r.sendToPublisher(serialized)
		batch = batch[:0]
	}

loop:
	for {
		select {
		case event, ok := <-r.events:
			if !ok {
				break loop
			}

			batch = append(batch, &proto.Event{
				Ttl:          &r.ttl,
				Host:         &hostname,
				Time:         &event.time,
				Service:      &event.service,
				MetricSint64: &event.metric,
				Attributes:   serializeRiemannAttributes(event.attributes),
			})

			if len(batch) < batchSize {
				continue
			}

		case <-ticker.Chan():
		}

		sendBatch()
	}

	sendBatch()
}

func (r *RiemannReporter) sendToPublisher(serialized []byte) {
	for {
		select {
		case r.messages <- serialized:
			return
		default:
		}

		select {
		case <-r.messages:
			r.logger.Printf("Riemann metric buffer full, dropping oldest batch")
		default:
		}
	}
}

func (r *RiemannReporter) publish() {
	defer r.wg.Done()

	for message := range r.messages {
		if err := r.publishMessage(message); err != nil {
			r.logger.Printf(
				"Failed to publish message to Riemann (%s)",
				err.Error(),
			)

			r.conn = nil
		}
	}

	if r.conn != nil {
		r.conn.Close()
	}
}

func (r *RiemannReporter) publishMessage(message []byte) error {
	if !r.ensureConnection() {
		return nil
	}

	if err := r.write(message); err != nil {
		return err
	}

	if err := r.read(); err != nil {
		return err
	}

	return nil
}

func (r *RiemannReporter) ensureConnection() bool {
	for r.conn != nil {
		return true
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-r.done:
		case <-ctx.Done():
		}

		cancel()
	}()

	return watchdog.BlockUntilSuccess(
		ctx,
		watchdog.RetryFunc(r.connect),
		backoff.NewExponentialBackoff(time.Millisecond*250, time.Second*30),
	)
}

func (r *RiemannReporter) connect() bool {
	conn, err := r.dialer()
	if err != nil {
		r.logger.Printf("Failed to connect to Riemann (%s)", err.Error())
		return false
	}

	r.conn = conn
	return true
}

func (r *RiemannReporter) write(message []byte) error {
	return writePrefixedMessage(r.conn, message)
}

func (r *RiemannReporter) read() error {
	data, err := readPrefixedMessage(r.conn)
	if err != nil {
		return err
	}

	response := &proto.Msg{}
	if err := pb.Unmarshal(data, response); err != nil {
		return err
	}

	if response.GetOk() != true {
		return fmt.Errorf("riemann returned a non-ok response")
	}

	return nil
}

//
// Helpers

func serializeRiemannAttributes(eventAttributes map[string]string) []*proto.Attribute {
	attributes := make([]*proto.Attribute, 0, len(eventAttributes))
	for key, value := range eventAttributes {
		attributes = append(attributes, &proto.Attribute{
			Key:   stringptr(key),
			Value: stringptr(value),
		})
	}

	return attributes
}

func readPrefixedMessage(r io.Reader) ([]byte, error) {
	var header uint32
	if err := binary.Read(r, binary.BigEndian, &header); err != nil {
		return nil, err
	}

	data := make([]byte, header)
	if err := readAll(r, data); err != nil {
		return nil, err
	}

	return data, nil
}

func writePrefixedMessage(w io.Writer, data []byte) error {
	prefix, err := serializeMessagePrefix(len(data))
	if err != nil {
		return nil
	}

	if err := writeAll(w, prefix); err != nil {
		return err
	}

	if err := writeAll(w, data); err != nil {
		return err
	}

	return nil
}

func serializeMessagePrefix(n int) ([]byte, error) {
	b := &bytes.Buffer{}
	if err := binary.Write(b, binary.BigEndian, uint32(n)); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func readAll(r io.Reader, p []byte) error {
	return bufio(r.Read, p)
}

func writeAll(w io.Writer, p []byte) error {
	return bufio(w.Write, p)
}

func bufio(f func([]byte) (int, error), p []byte) error {
	for len(p) > 0 {
		n, err := f(p)
		if err != nil {
			return err
		}

		p = p[n:]
	}

	return nil
}
