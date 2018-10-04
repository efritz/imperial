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

func (r *Reporter) RegisterCounter(name string, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.RegisterCounter(name, configs...) })
}

func (r *Reporter) RegisterGauge(name string, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.RegisterGauge(name, configs...) })
}

func (r *Reporter) RegisterHistogram(name string, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.RegisterHistogram(name, configs...) })
}

func (r *Reporter) RegisterSummary(name string, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.RegisterSummary(name, configs...) })
}

func (r *Reporter) AddCounter(name string, value float64, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.AddCounter(name, value, configs...) })
}

func (r *Reporter) AddGauge(name string, value float64, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.AddGauge(name, value, configs...) })
}

func (r *Reporter) SetGauge(name string, value float64, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.SetGauge(name, value, configs...) })
}

func (r *Reporter) ObserveHistogram(name string, value float64, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.ObserveHistogram(name, value, configs...) })
}

func (r *Reporter) ObserveSummary(name string, value float64, configs ...base.ConfigFunc) {
	r.each(func(r base.Reporter) { r.ObserveSummary(name, value, configs...) })
}

func (r *Reporter) Shutdown() {
	r.each(func(r base.Reporter) { r.Shutdown() })
}

func (r *Reporter) each(f func(r base.Reporter)) {
	for _, reporter := range r.reporters {
		f(reporter)
	}
}
