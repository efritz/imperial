package red

import (
	"github.com/efritz/imperial/base"
)

type (
	config struct {
		logger        base.Logger
		prefixConfigs map[string]*PrefixConfig
	}

	ConfigFunc func(r *config)
)

func newConfig() *config {
	return &config{
		logger:        base.NilLogger,
		prefixConfigs: map[string]*PrefixConfig{},
	}
}

func WithLogger(logger base.Logger) ConfigFunc {
	return func(c *config) { c.logger = logger }
}

func WithPrefixConfig(prefix string, prefixConfig *PrefixConfig) ConfigFunc {
	return func(c *config) { c.prefixConfigs[prefix] = prefixConfig }
}

func WithPrefixConfigs(prefixConfigs map[string]*PrefixConfig) ConfigFunc {
	return func(c *config) {
		for prefix, config := range prefixConfigs {
			c.prefixConfigs[prefix] = config
		}
	}
}
