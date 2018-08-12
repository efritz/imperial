package imperial

import (
	"time"

	"github.com/efritz/glock"
)

type RiemannConfigFunc func(r *RiemannReporter)

func WithRiemannReportConfigs(configs ...ConfigFunc) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.configs = append(r.configs, configs...) }
}

func WithRiemannBatchSize(batchSize int) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.batchSize = batchSize }
}

func WithRiemannQueueSize(queueSize int) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.queueSize = queueSize }
}

func WithRiemannTickDuration(tickDuration time.Duration) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.tickDuration = tickDuration }
}

func WithRiemannConnectionTimeout(connectionTimeout time.Duration) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.connectionTimeout = connectionTimeout }
}

func WithRiemannTTL(ttl float32) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.ttl = ttl }
}

func WithRiemannLogger(logger Logger) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.logger = logger }
}

func WithRiemannClock(clock glock.Clock) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.clock = clock }
}

func WithRiemannDialer(dialer RiemannDialer) RiemannConfigFunc {
	return func(r *RiemannReporter) { r.dialer = dialer }
}
