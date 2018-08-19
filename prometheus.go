package imperial

import "github.com/efritz/imperial/prometheus"

var (
	NewPrometheusReporter                    = prometheus.NewReporter
	WithPrometheusReportConfigs              = prometheus.WithReportConfigs
	WithPrometheusHandlerMaxRequestsInFlight = prometheus.WithHandlerMaxRequestsInFlight
	WithPrometheusHandlerTimeout             = prometheus.WithHandlerTimeout
	WithPrometheusLogger                     = prometheus.WithLogger
	WithPrometheusRegistry                   = prometheus.WithRegistry
)
