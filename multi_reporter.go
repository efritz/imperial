package imperial

type MultiReporter struct {
	reporters []Reporter
}

func NewMultiReporter(reporters ...Reporter) Reporter {
	return &MultiReporter{
		reporters: reporters,
	}

}
func (r *MultiReporter) Report(name string, value int, configs ...ConfigFunc) {
	for _, reporter := range r.reporters {
		reporter.Report(name, value, configs...)
	}
}

func (r *MultiReporter) Shutdown() {
	for _, reporter := range r.reporters {
		reporter.Shutdown()
	}
}
