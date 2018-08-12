package base

type (
	ReportConfig struct {
		Namespace  string
		Unit       Unit
		Attributes map[string]string
	}

	ConfigFunc func(rc *ReportConfig)
)

func WithNamespace(namespace string) ConfigFunc {
	return func(rc *ReportConfig) { rc.Namespace = namespace }
}

func WithUnit(unit Unit) ConfigFunc {
	return func(rc *ReportConfig) { rc.Unit = unit }
}

func WithAttributes(attributes map[string]string) ConfigFunc {
	return func(rc *ReportConfig) {
		for k, v := range attributes {
			rc.Attributes[k] = v
		}
	}
}

func ApplyConfigs(baseConfigs []ConfigFunc, configs []ConfigFunc) *ReportConfig {
	reportConfig := &ReportConfig{
		Unit:       UnitCount,
		Attributes: map[string]string{},
	}

	for _, f := range append(baseConfigs, configs...) {
		f(reportConfig)
	}

	return reportConfig
}
