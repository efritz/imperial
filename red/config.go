package red

import "github.com/efritz/imperial/base"

type (
	PrefixConfig struct {
		buckets          []float64
		errorInterpreter ErrorInterpreter
		configs          []base.ConfigFunc
	}

	PrefixConfigFunc func(r *PrefixConfig)
)

var defaultBuckets = []float64{0.01, 0.1, 0.5, 1}

func NewPrefixConfig(configs ...PrefixConfigFunc) *PrefixConfig {
	config := &PrefixConfig{
		buckets:          defaultBuckets,
		errorInterpreter: DefaultErrorInterpreter,
		configs:          []base.ConfigFunc{},
	}

	for _, f := range configs {
		f(config)
	}

	return config
}

func WithPrefixBuckets(buckets []float64) PrefixConfigFunc {
	return func(c *PrefixConfig) { c.buckets = buckets }
}

func WithPrefixErrorInterpreter(errorInterpreter ErrorInterpreter) PrefixConfigFunc {
	return func(c *PrefixConfig) { c.errorInterpreter = errorInterpreter }
}

func WithPrefixReportConfigs(configs ...base.ConfigFunc) PrefixConfigFunc {
	return func(c *PrefixConfig) { c.configs = append(c.configs, configs...) }
}
