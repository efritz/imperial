package base

type (
	ReportConfig struct {
		Namespace         string
		Subsystem         string
		Help              string
		Unit              Unit
		Attributes        map[string]string
		ExposedAttributes []string
		Buckets           []float64
		Quantiles         map[float64]float64
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

func WithSubsystem(subsystem string) ConfigFunc {
	return func(rc *ReportConfig) { rc.Subsystem = subsystem }
}

func WithHelp(help string) ConfigFunc {
	return func(rc *ReportConfig) { rc.Help = help }
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

func WithExposedAttributes(names ...string) ConfigFunc {
	return func(rc *ReportConfig) {
		rc.ExposedAttributes = append(rc.ExposedAttributes, names...)
	}
}

func WithBuckets(buckets []float64) ConfigFunc {
	return func(rc *ReportConfig) { rc.Buckets = buckets }
}

func WithQuantiles(quantiles map[float64]float64) ConfigFunc {
	return func(rc *ReportConfig) { rc.Quantiles = quantiles }
}

func ApplyConfigs(baseConfigs []ConfigFunc, configs []ConfigFunc) *ReportConfig {
	reportConfig := &ReportConfig{
		Help:              "<no help>",
		Unit:              UnitCount,
		Attributes:        map[string]string{},
		ExposedAttributes: []string{},
	}

	for _, f := range append(baseConfigs, configs...) {
		f(reportConfig)
	}

	return reportConfig
}
