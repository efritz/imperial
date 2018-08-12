package imperial

import (
	"time"

	"github.com/efritz/glock"
)

type RiemannConfigFunc func(r *riemannConfig)

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
