package imperial

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/aphistic/sweet"
	"github.com/efritz/glock"
	"github.com/efritz/imperial/proto"
	pb "github.com/golang/protobuf/proto"
	. "github.com/onsi/gomega"
)

type RiemannSuite struct{}

func (s *RiemannSuite) TestReport(t sweet.T) {
	reporter, w, clock := makeRiemannReporter()

	t1 := time.Now()
	t2 := time.Now().Add(time.Minute)

	clock.SetCurrent(t1)
	reporter.Report("a", 1)
	reporter.Report("b", 2)
	reporter.Report("c", 3)
	clock.SetCurrent(t2)
	reporter.Report("d", 4)
	reporter.Report("e", 5)

	Eventually(func() ([]*riemannEvent, error) {
		return deserializeRiemannBatch(w.Bytes())
	}).Should(ConsistOf(
		&riemannEvent{"a", 1, t1.Unix(), map[string]string{}},
		&riemannEvent{"b", 2, t1.Unix(), map[string]string{}},
		&riemannEvent{"c", 3, t1.Unix(), map[string]string{}},
		&riemannEvent{"d", 4, t2.Unix(), map[string]string{}},
		&riemannEvent{"e", 5, t2.Unix(), map[string]string{}},
	))
}

func (s *RiemannSuite) TestReportWithAttributes(t sweet.T) {
	reporter, w, clock := makeRiemannReporter(
		WithRiemannReportConfigs(WithAttributes(map[string]string{
			"x": "xv",
			"y": "xy",
		})),
	)

	reporter.Report("a", 1)
	reporter.Report("b", 2, WithAttributes(map[string]string{"z": "z1"}))
	reporter.Report("c", 3)
	reporter.Report("d", 4, WithAttributes(map[string]string{"z": "z2"}))
	reporter.Report("e", 5)

	Eventually(func() ([]*riemannEvent, error) {
		return deserializeRiemannBatch(w.Bytes())
	}).Should(ConsistOf(
		&riemannEvent{"a", 1, clock.Now().Unix(), map[string]string{"x": "xv", "y": "xy"}},
		&riemannEvent{"b", 2, clock.Now().Unix(), map[string]string{"x": "xv", "y": "xy", "z": "z1"}},
		&riemannEvent{"c", 3, clock.Now().Unix(), map[string]string{"x": "xv", "y": "xy"}},
		&riemannEvent{"d", 4, clock.Now().Unix(), map[string]string{"x": "xv", "y": "xy", "z": "z2"}},
		&riemannEvent{"e", 5, clock.Now().Unix(), map[string]string{"x": "xv", "y": "xy"}},
	))
}

func (s *RiemannSuite) TestMultipleBatches(t sweet.T) {
	reporter, w, clock := makeRiemannReporter()

	for i := 0; i < 3; i++ {
		w.Reset()
		clock.Advance(time.Second)
		t1 := clock.Now().Unix()

		reporter.Report("a", 10*i+1)
		reporter.Report("b", 10*i+2)
		reporter.Report("c", 10*i+3)
		reporter.Report("d", 10*i+4)
		reporter.Report("e", 10*i+5)

		Eventually(func() ([]*riemannEvent, error) {
			return deserializeRiemannBatch(w.Bytes())
		}).Should(ConsistOf(
			&riemannEvent{"a", 10*int64(i) + 1, t1, map[string]string{}},
			&riemannEvent{"b", 10*int64(i) + 2, t1, map[string]string{}},
			&riemannEvent{"c", 10*int64(i) + 3, t1, map[string]string{}},
			&riemannEvent{"d", 10*int64(i) + 4, t1, map[string]string{}},
			&riemannEvent{"e", 10*int64(i) + 5, t1, map[string]string{}},
		))
	}

	w.Reset()
	reporter.Report("a", 41)
	reporter.Report("b", 42)
	Consistently(w.Bytes).Should(BeEmpty())
}

func (s *RiemannSuite) TestPartialBatchTick(t sweet.T) {
	reporter, w, clock := makeRiemannReporter()

	for i := 0; i < 3; i++ {
		w.Reset()
		reporter.Report("a", 10*i+1)
		Consistently(w.Bytes).Should(BeEmpty())

		t1 := clock.Now().Unix()
		clock.Advance(time.Second * 5)

		Eventually(func() ([]*riemannEvent, error) {
			return deserializeRiemannBatch(w.Bytes())
		}).Should(ConsistOf(
			&riemannEvent{"a", 10*int64(i) + 1, t1, map[string]string{}},
		))
	}
}

func (s *RiemannSuite) TestFullBuffer(t sweet.T) {
	var (
		conn, w = makeRiemannConn()
		clock   = glock.NewMockClock()
		block   = make(chan struct{})
	)

	dialer := func() (io.ReadWriteCloser, error) {
		<-block
		return conn, nil
	}

	reporter := NewRiemannReporter(
		"localhost:5555",
		WithRiemannClock(clock),
		WithRiemannDialer(dialer),
		WithRiemannBatchSize(5),
		WithRiemannQueueSize(25),
	)

	for i := 0; i < 500; i++ {
		if i == 450 {
			// Wait for writes to propagate
			<-time.After(time.Millisecond * 50)
			close(block)
			<-time.After(time.Millisecond * 50)
		}

		reporter.Report("a", i)
	}

	Eventually(func() (int, error) {
		messages, err := deserializeRiemannBatches(w.Bytes())
		return len(messages), err
	}).Should(BeNumerically("~", 200, 25))
}

func (s *RiemannSuite) TestReconnect(t sweet.T) {
	var (
		w           = &bytes.Buffer{}
		conn        = NewMockConn(makeRiemannReader(), w)
		failingConn = NewMockConn(makeRiemannReader(), &failingWriter{})
		clock       = glock.NewMockClock()
		dials       = 0
	)

	dialer := func() (io.ReadWriteCloser, error) {
		dials++

		if dials < 5 {
			return failingConn, nil
		}

		return conn, nil
	}

	reporter := NewRiemannReporter(
		"localhost:5555",
		WithRiemannClock(clock),
		WithRiemannDialer(dialer),
		WithRiemannBatchSize(5),
	)

	for i := 0; i < 30; i++ {
		reporter.Report("a", i)
	}

	Eventually(func() ([]*riemannEvent, error) {
		return deserializeRiemannBatches(w.Bytes())
	}).Should(ConsistOf(
		&riemannEvent{"a", 20, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 21, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 22, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 23, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 24, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 25, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 26, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 27, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 28, clock.Now().Unix(), map[string]string{}},
		&riemannEvent{"a", 29, clock.Now().Unix(), map[string]string{}},
	))

	Expect(dials).To(Equal(5))
}

func (s *RiemannSuite) TestShutdownClosesConnection(t sweet.T) {
	reporter := NewRiemannReporter(
		"localhost:5555",
		WithRiemannDialer(func() (io.ReadWriteCloser, error) {
			select {}
		}),
	)

	reporter.Report("a", 1)
	reporter.Report("b", 2)
	reporter.Report("c", 3)
	reporter.Report("d", 4)
	reporter.Report("e", 5)

	// Should unblock
	reporter.Shutdown()
}

func (s *RiemannSuite) TestShutdownDuringConnection(t sweet.T) {
	conn, _ := makeRiemannConn()

	reporter := NewRiemannReporter(
		"localhost:5555",
		WithRiemannDialer(func() (io.ReadWriteCloser, error) {
			return conn, nil
		}),
	)

	reporter.Report("a", 1)
	reporter.Report("b", 2)
	reporter.Report("c", 3)
	reporter.Report("d", 4)
	reporter.Report("e", 5)
	reporter.Report("f", 6)

	reporter.Shutdown()
	Eventually(func() bool { return conn.closed }).Should(BeTrue())
}

//
// Constructors

func makeRiemannReporter(configs ...RiemannConfigFunc) (Reporter, *bytes.Buffer, *glock.MockClock) {
	var (
		conn, w = makeRiemannConn()
		clock   = glock.NewMockClock()
	)

	dialer := func() (io.ReadWriteCloser, error) {
		return conn, nil
	}

	reporter := NewRiemannReporter(
		"localhost:5555",
		append(
			[]RiemannConfigFunc{
				WithRiemannClock(clock),
				WithRiemannDialer(dialer),
				WithRiemannBatchSize(5),
			},
			configs...,
		)...,
	)

	return reporter, w, clock
}

func makeRiemannConn() (*mockConn, *bytes.Buffer) {
	var (
		w = &bytes.Buffer{}
		r = makeRiemannReader()
		c = NewMockConn(r, w)
	)

	return c, w
}

func makeRiemannReader() io.Reader {
	data, _ := pb.Marshal(&proto.Msg{Ok: boolptr(true)})
	prefix, _ := serializeMessagePrefix(len(data))

	return bytes.NewReader(append(prefix, data...))
}

//
// Mock Riemann Connection

type mockConn struct {
	r      io.Reader
	w      io.Writer
	closed bool
}

func NewMockConn(r io.Reader, w io.Writer) *mockConn {
	return &mockConn{r: r, w: w}
}

func (c *mockConn) Read(p []byte) (int, error)  { return c.r.Read(p) }
func (c *mockConn) Write(p []byte) (int, error) { return c.w.Write(p) }
func (c *mockConn) Close() error                { c.closed = true; return nil }

type failingWriter struct{}

func (w *failingWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("bad write")
}

//
// Payload Deserialization Helpers

func deserializeRiemannBatch(data []byte) ([]*riemannEvent, error) {
	reader := bytes.NewReader(data)
	events, err := riemannEventsFromReader(reader)
	if err != nil {
		return nil, err
	}

	if _, err := reader.ReadByte(); err != io.EOF {
		return nil, fmt.Errorf("multiple writes occurred")
	}

	return events, nil
}

func deserializeRiemannBatches(data []byte) ([]*riemannEvent, error) {
	reader := bytes.NewReader(data)

	allEvents := []*riemannEvent{}

	for {
		events, err := riemannEventsFromReader(reader)
		if err != nil {
			if err == io.EOF {
				return allEvents, nil
			}

			return nil, err
		}

		allEvents = append(allEvents, events...)
	}
}

func riemannEventsFromReader(r io.Reader) ([]*riemannEvent, error) {
	raw, err := readPrefixedMessage(r)
	if err != nil {
		return nil, err
	}

	payload := &proto.Msg{}
	if err := pb.Unmarshal(raw, payload); err != nil {
		return nil, err
	}

	events := []*riemannEvent{}
	for _, event := range payload.Events {
		attributes := map[string]string{}
		for _, pair := range event.GetAttributes() {
			attributes[pair.GetKey()] = pair.GetValue()
		}

		events = append(events, &riemannEvent{
			service:    event.GetService(),
			metric:     event.GetMetricSint64(),
			time:       event.GetTime(),
			attributes: attributes,
		})
	}

	return events, nil
}
