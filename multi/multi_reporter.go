package multi

import "github.com/efritz/imperial/base"

type Reporter struct {
	reporters []base.Reporter
}

func NewReporter(reporters ...base.Reporter) *Reporter {
	return &Reporter{
		reporters: reporters,
	}

}
func (r *Reporter) Report(name string, value int, configs ...base.ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.Report(name, value, configs...)
	}
}

func (r *Reporter) Shutdown() {
	for _, reporter := range r.reporters {
		reporter.Shutdown()
	}
}
