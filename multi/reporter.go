package multi

import "github.com/efritz/imperial/base"

type Reporter struct {
	reporters []base.Reporter
}

var _ base.Reporter = &Reporter{}

func NewReporter(reporters ...base.Reporter) *Reporter {
	return &Reporter{
		reporters: reporters,
	}
}

func (r *Reporter) AddCounter(name string, value float64, configs ...base.ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.AddCounter(name, value, configs...)
	}
}

func (r *Reporter) AddGauge(name string, value float64, configs ...base.ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.AddGauge(name, value, configs...)
	}
}

func (r *Reporter) SetGauge(name string, value float64, configs ...base.ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.SetGauge(name, value, configs...)
	}
}

func (r *Reporter) ObserveHistogram(name string, value float64, configs ...base.ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.ObserveHistogram(name, value, configs...)
	}
}

func (r *Reporter) ObserveSummary(name string, value float64, configs ...base.ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.ObserveSummary(name, value, configs...)
	}
}

func (r *Reporter) Shutdown() {
	for _, reporter := range r.reporters {
		reporter.Shutdown()
	}
}
