package base

type (
	ReportConfig struct {
		Namespace  string
		Unit       Unit
		Attributes map[string]string
	}

	Unit       string
	ConfigFunc func(rc *ReportConfig)
)

const (
	UnitCount        Unit = "Count"
	UnitPercent      Unit = "Percent"
	UnitSeconds      Unit = "Seconds"
	UnitMicroseconds Unit = "Microseconds"
	UnitMilliseconds Unit = "Milliseconds"
	UnitBytes        Unit = "Bytes"
	UnitKilobytes    Unit = "Kilobytes"
	UnitMegabytes    Unit = "Megabytes"
	UnitGigabytes    Unit = "Gigabytes"
	UnitTerabytes    Unit = "Terabytes"
	UnitBits         Unit = "Bits"
	UnitKilobits     Unit = "Kilobits"
	UnitMegabits     Unit = "Megabits"
	UnitGigabits     Unit = "Gigabits"
	UnitTerabits     Unit = "Terabits"
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
