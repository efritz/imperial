package riemann

import (
	"io"
	"net"
	"time"

	"github.com/efritz/glock"

	"github.com/efritz/imperial/base"
)

type (
	config struct {
		logger            base.Logger
		clock             glock.Clock
		dialer            Dialer
		configs           []base.ConfigFunc
		ttl               float32
		batchSize         int
		queueSize         int
		tickDuration      time.Duration
		connectionTimeout time.Duration
	}

	ConfigFunc func(r *config)
)

func newConfig() *config {
	return &config{
		logger:            base.NewNilLogger(),
		clock:             glock.NewRealClock(),
		configs:           []base.ConfigFunc{},
		ttl:               60,
		batchSize:         5000,
		queueSize:         360,
		tickDuration:      time.Second * 5,
		connectionTimeout: time.Second * 5,
	}
}

func WithReportConfigs(configs ...base.ConfigFunc) ConfigFunc {
	return func(r *config) { r.configs = append(r.configs, configs...) }
}

func WithBatchSize(batchSize int) ConfigFunc {
	return func(r *config) { r.batchSize = batchSize }
}

func WithQueueSize(queueSize int) ConfigFunc {
	return func(r *config) { r.queueSize = queueSize }
}

func WithTickDuration(tickDuration time.Duration) ConfigFunc {
	return func(r *config) { r.tickDuration = tickDuration }
}

func WithConnectionTimeout(connectionTimeout time.Duration) ConfigFunc {
	return func(r *config) { r.connectionTimeout = connectionTimeout }
}

func WithTTL(ttl float32) ConfigFunc {
	return func(r *config) { r.ttl = ttl }
}

func WithLogger(logger base.Logger) ConfigFunc {
	return func(r *config) { r.logger = logger }
}

func WithClock(clock glock.Clock) ConfigFunc {
	return func(r *config) { r.clock = clock }
}

func WithDialer(dialer Dialer) ConfigFunc {
	return func(r *config) { r.dialer = dialer }
}

//
//

func makeDialer(addr string, config *config) Dialer {
	if config.dialer != nil {
		return config.dialer
	}

	return func() (io.ReadWriteCloser, error) {
		return net.DialTimeout(
			"tcp",
			addr,
			config.connectionTimeout,
		)
	}
}
