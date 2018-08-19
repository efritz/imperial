package imperial

import "github.com/efritz/imperial/riemann"

type RiemannDialer = riemann.Dialer

var (
	WithRiemannReportConfigs     = riemann.WithReportConfigs
	WithRiemannBatchSize         = riemann.WithBatchSize
	WithRiemannQueueSize         = riemann.WithQueueSize
	WithRiemannTickDuration      = riemann.WithTickDuration
	WithRiemannConnectionTimeout = riemann.WithConnectionTimeout
	WithRiemannTTL               = riemann.WithTTL
	WithRiemannLogger            = riemann.WithLogger
	WithRiemannClock             = riemann.WithClock
	WithRiemannDialer            = riemann.WithDialer
)

func NewRiemannReporter(addr string, configs ...riemann.ConfigFunc) Reporter {
	return NewSimpleReporterShim(riemann.NewReporter(addr, configs...))
}
