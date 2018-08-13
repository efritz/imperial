package riemann

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
	"github.com/efritz/watchdog"
	pb "github.com/golang/protobuf/proto"

	"github.com/efritz/imperial/base"
	"github.com/efritz/imperial/proto"
)

type (
	Reporter struct {
		logger   base.Logger
		clock    glock.Clock
		dialer   Dialer
		configs  []base.ConfigFunc
		ttl      float32
		conn     io.ReadWriteCloser
		events   chan *event
		messages chan []byte
		done     chan struct{}
		once     *sync.Once
		wg       *sync.WaitGroup
	}

	Dialer func() (io.ReadWriteCloser, error)

	event struct {
		service    string
		metric     float64
		time       int64
		attributes map[string]string
	}
)

var _ base.SimpleReporter = &Reporter{}

func NewReporter(addr string, configs ...ConfigFunc) *Reporter {
	config := newConfig()
	for _, f := range configs {
		f(config)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	reporter := &Reporter{
		dialer:   makeDialer(addr, config),
		logger:   config.logger,
		clock:    config.clock,
		configs:  config.configs,
		ttl:      config.ttl,
		events:   make(chan *event),
		messages: make(chan []byte, config.queueSize),
		done:     make(chan struct{}),
		once:     &sync.Once{},
		wg:       wg,
	}

	go reporter.batch(config.batchSize, config.tickDuration)
	go reporter.publish()

	return reporter
}

func (r *Reporter) Report(name string, value float64, configs ...base.ConfigFunc) {
	options := base.ApplyConfigs(r.configs, configs)

	r.events <- &event{
		service:    name,
		metric:     value,
		time:       r.clock.Now().Unix(),
		attributes: options.Attributes,
	}
}

func (r *Reporter) Shutdown() {
	r.once.Do(func() {
		close(r.done)
		close(r.events)
	})

	r.wg.Wait()
}

func (r *Reporter) batch(batchSize int, tickDuration time.Duration) {
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
				Ttl:        &r.ttl,
				Host:       &hostname,
				Time:       &event.time,
				Service:    &event.service,
				MetricD:    &event.metric,
				Attributes: serializeAttributes(event.attributes),
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

func (r *Reporter) sendToPublisher(serialized []byte) {
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

func (r *Reporter) publish() {
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

func (r *Reporter) publishMessage(message []byte) error {
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

func (r *Reporter) ensureConnection() bool {
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

func (r *Reporter) connect() bool {
	conn, err := r.dialer()
	if err != nil {
		r.logger.Printf("Failed to connect to Riemann (%s)", err.Error())
		return false
	}

	r.conn = conn
	return true
}

func (r *Reporter) write(message []byte) error {
	return writePrefixedMessage(r.conn, message)
}

func (r *Reporter) read() error {
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
// Serialization Helpers

func serializeAttributes(eventAttributes map[string]string) []*proto.Attribute {
	attributes := make([]*proto.Attribute, 0, len(eventAttributes))
	for key, value := range eventAttributes {
		attributes = append(attributes, &proto.Attribute{
			Key:   strptr(key),
			Value: strptr(value),
		})
	}

	return attributes
}

func strptr(val string) *string {
	return &val
}

//
// IO Helpers

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
