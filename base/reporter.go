package base

type Reporter interface {
	AddCounter(name string, value float64, configs ...ConfigFunc)
	AddGauge(name string, value float64, configs ...ConfigFunc)
	SetGauge(name string, value float64, configs ...ConfigFunc)
	ObserveHistogram(name string, value float64, configs ...ConfigFunc)
	ObserveSummary(name string, value float64, configs ...ConfigFunc)
	Shutdown()
}
