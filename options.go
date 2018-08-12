package imperial

type (
	reportConfig struct {
		attributes map[string]string
	}

	ConfigFunc func(rc *reportConfig)
)

func WithAttributes(attributes map[string]string) ConfigFunc {
	return func(rc *reportConfig) {
		for k, v := range attributes {
			rc.attributes[k] = v
		}
	}
}

func applyConfigs(baseConfigs []ConfigFunc, configs []ConfigFunc) *reportConfig {
	reportConfig := &reportConfig{
		attributes: map[string]string{},
	}

	for _, f := range append(baseConfigs, configs...) {
		f(reportConfig)
	}

	return reportConfig
}
