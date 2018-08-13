package base

type (
	SimpleReporter interface {
		Report(name string, value float64, configs ...ConfigFunc)
		Shutdown()
	}

	SimpleReporterShim struct {
		reporter SimpleReporter
		counters map[string]float64
		gauges   map[string]float64
	}
)

func NewSimpleReporterShim(reporter SimpleReporter) *SimpleReporterShim {
	return &SimpleReporterShim{
		reporter: reporter,
		counters: map[string]float64{},
		gauges:   map[string]float64{},
	}
}

func (r *SimpleReporterShim) AddCounter(name string, value float64, configs ...ConfigFunc) {
	if value < 0 {
		// TODO - panic
	}

	if _, ok := r.counters[name]; !ok {
		r.counters[name] = 0
	}

	// TODO - with locks or sync/atomic
	r.counters[name] += value
	r.reporter.Report(name, r.counters[name], configs...)
}

func (r *SimpleReporterShim) AddGauge(name string, value float64, configs ...ConfigFunc) {
	if _, ok := r.gauges[name]; !ok {
		r.gauges[name] = 0
	}

	// TODO - with locks or sync/atomic
	r.gauges[name] += value
	r.reporter.Report(name, r.gauges[name], configs...)
}

func (r *SimpleReporterShim) SetGauge(name string, value float64, configs ...ConfigFunc) {
	if _, ok := r.gauges[name]; !ok {
		r.gauges[name] = 0
	}

	// TODO - with locks or sync/atomic
	r.gauges[name] = value
	r.reporter.Report(name, r.gauges[name], configs...)
}

func (r *SimpleReporterShim) ObserveHistogram(name string, value float64, configs ...ConfigFunc) {
	r.reporter.Report(name, value, configs...)
}

func (r *SimpleReporterShim) ObserveSummary(name string, value float64, configs ...ConfigFunc) {
	r.reporter.Report(name, value, configs...)
}

func (r *SimpleReporterShim) Shutdown() {
	r.reporter.Shutdown()
}
