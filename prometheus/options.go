package prometheus

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/efritz/imperial/base"
)

type (
	config struct {
		logger                     base.Logger
		configs                    []base.ConfigFunc
		handlerMaxRequestsInFlight int
		handlerTimeout             time.Duration
		registry                   Registry
	}

	ConfigFunc func(r *config)
)

func newConfig() *config {
	return &config{
		logger:         base.NewNilLogger(),
		configs:        []base.ConfigFunc{},
		handlerTimeout: time.Second * 5,
		registry:       prometheus.NewRegistry(),
	}
}

func WithReportConfigs(configs ...base.ConfigFunc) ConfigFunc {
	return func(c *config) { c.configs = append(c.configs, configs...) }
}

func WithHandlerMaxRequestsInFlight(maxRequestsInFlight int) ConfigFunc {
	return func(c *config) { c.handlerMaxRequestsInFlight = maxRequestsInFlight }
}

func WithHandlerTimeout(handlerTimeout time.Duration) ConfigFunc {
	return func(c *config) { c.handlerTimeout = handlerTimeout }
}

func WithLogger(logger base.Logger) ConfigFunc {
	return func(c *config) { c.logger = logger }
}

func WithRegistry(registry Registry) ConfigFunc {
	return func(c *config) { c.registry = registry }
}
