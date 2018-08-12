package imperial

import (
	"io"
	"net"
	"time"

	"github.com/efritz/glock"
)

type (
	riemannConfig struct {
		logger            Logger
		clock             glock.Clock
		dialer            RiemannDialer
		configs           []ConfigFunc
		ttl               float32
		batchSize         int
		queueSize         int
		tickDuration      time.Duration
		connectionTimeout time.Duration
	}

	RiemannConfigFunc func(r *riemannConfig)
)

func newRiemannConfig() *riemannConfig {
	return &riemannConfig{
		logger:            NewNilLogger(),
		clock:             glock.NewRealClock(),
		configs:           []ConfigFunc{},
		ttl:               60,
		batchSize:         5000,
		queueSize:         360,
		tickDuration:      time.Second * 5,
		connectionTimeout: time.Second * 5,
	}
}

func WithRiemannReportConfigs(configs ...ConfigFunc) RiemannConfigFunc {
	return func(r *riemannConfig) { r.configs = append(r.configs, configs...) }
}

func WithRiemannBatchSize(batchSize int) RiemannConfigFunc {
	return func(r *riemannConfig) { r.batchSize = batchSize }
}

func WithRiemannQueueSize(queueSize int) RiemannConfigFunc {
	return func(r *riemannConfig) { r.queueSize = queueSize }
}

func WithRiemannTickDuration(tickDuration time.Duration) RiemannConfigFunc {
	return func(r *riemannConfig) { r.tickDuration = tickDuration }
}

func WithRiemannConnectionTimeout(connectionTimeout time.Duration) RiemannConfigFunc {
	return func(r *riemannConfig) { r.connectionTimeout = connectionTimeout }
}

func WithRiemannTTL(ttl float32) RiemannConfigFunc {
	return func(r *riemannConfig) { r.ttl = ttl }
}

func WithRiemannLogger(logger Logger) RiemannConfigFunc {
	return func(r *riemannConfig) { r.logger = logger }
}

func WithRiemannClock(clock glock.Clock) RiemannConfigFunc {
	return func(r *riemannConfig) { r.clock = clock }
}

func WithRiemannDialer(dialer RiemannDialer) RiemannConfigFunc {
	return func(r *riemannConfig) { r.dialer = dialer }
}

//
//

func makeDialer(addr string, config *riemannConfig) RiemannDialer {
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
