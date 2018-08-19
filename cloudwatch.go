package imperial

import "github.com/efritz/imperial/cloudwatch"

var (
	WithCloudwatchReportConfigs = cloudwatch.WithReportConfigs
	WithCloudwatchBatchSize     = cloudwatch.WithBatchSize
	WithCloudwatchBufferSize    = cloudwatch.WithBufferSize
	WithCloudwatchTickDuration  = cloudwatch.WithTickDuration
	WithCloudwatchLogger        = cloudwatch.WithLogger
	WithCloudwatchClock         = cloudwatch.WithClock
	WithCloudwatchSession       = cloudwatch.WithSession
	WithCloudwatchAPI           = cloudwatch.WithAPI
)

func NewCloudwatchReporter(addr string, configs ...cloudwatch.ConfigFunc) Reporter {
	return NewSimpleReporterShim(cloudwatch.NewReporter(addr, configs...))
}
