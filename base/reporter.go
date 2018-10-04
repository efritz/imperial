package base

type Reporter interface {
	RegisterCounter(name string, configs ...ConfigFunc)
	RegisterGauge(name string, configs ...ConfigFunc)
	RegisterHistogram(name string, configs ...ConfigFunc)
	RegisterSummary(name string, configs ...ConfigFunc)
	AddCounter(name string, value float64, configs ...ConfigFunc)
	AddGauge(name string, value float64, configs ...ConfigFunc)
	SetGauge(name string, value float64, configs ...ConfigFunc)
	ObserveHistogram(name string, value float64, configs ...ConfigFunc)
	ObserveSummary(name string, value float64, configs ...ConfigFunc)
	Shutdown()
}
