package imperial

type (
	reportConfig struct {
		attributes          map[string]string
		cloudwatchNamespace string
		cloudwatchUnit      CloudwatchUnit
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

func WithCloudwatchNamespace(namespace string) ConfigFunc {
	return func(rc *reportConfig) { rc.cloudwatchNamespace = namespace }
}

func WithCloudwatchUnit(unit CloudwatchUnit) ConfigFunc {
	return func(rc *reportConfig) { rc.cloudwatchUnit = unit }
}

func applyConfigs(baseConfigs []ConfigFunc, configs []ConfigFunc) *reportConfig {
	reportConfig := &reportConfig{
		attributes:     map[string]string{},
		cloudwatchUnit: CloudwatchUnitCount,
	}

	for _, f := range append(baseConfigs, configs...) {
		f(reportConfig)
	}

	return reportConfig
}
