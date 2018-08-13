package base

import (
	"math"
	"sync"
	"sync/atomic"
)

type (
	SimpleReporter interface {
		Report(name string, value float64, configs ...ConfigFunc)
		Shutdown()
	}

	SimpleReporterShim struct {
		reporter SimpleReporter
		counters map[string]*value
		gauges   map[string]*value
		mutex    *sync.RWMutex
	}

	value struct {
		val uint64
	}
)

func NewSimpleReporterShim(reporter SimpleReporter) *SimpleReporterShim {
	return &SimpleReporterShim{
		reporter: reporter,
		counters: map[string]*value{},
		gauges:   map[string]*value{},
		mutex:    &sync.RWMutex{},
	}
}

func (r *SimpleReporterShim) AddCounter(name string, value float64, configs ...ConfigFunc) {
	if value < 0 {
		panic("counter cannot decrease in value")
	}

	r.ensureValue(r.counters, name)
	r.reporter.Report(name, r.counters[name].Add(value), configs...)
}

func (r *SimpleReporterShim) AddGauge(name string, value float64, configs ...ConfigFunc) {
	r.ensureValue(r.gauges, name)
	r.reporter.Report(name, r.gauges[name].Add(value), configs...)
}

func (r *SimpleReporterShim) SetGauge(name string, value float64, configs ...ConfigFunc) {
	r.ensureValue(r.gauges, name)
	r.reporter.Report(name, r.gauges[name].Set(value), configs...)
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

func (r *SimpleReporterShim) ensureValue(m map[string]*value, name string) {
	r.mutex.RLock()
	if _, ok := m[name]; ok {
		r.mutex.RUnlock()
		return
	}

	r.mutex.RUnlock()
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := m[name]; ok {
		return
	}

	m[name] = &value{}
}

//
//

func (v *value) Add(value float64) float64 {
	for {
		var (
			oldBits  = atomic.LoadUint64(&v.val)
			newValue = math.Float64frombits(oldBits) + value
			newBits  = math.Float64bits(newValue)
		)

		if atomic.CompareAndSwapUint64(&v.val, oldBits, newBits) {
			return newValue
		}
	}
}

func (v *value) Set(value float64) float64 {
	for {
		var (
			oldBits = atomic.LoadUint64(&v.val)
			newBits = math.Float64bits(value)
		)

		if atomic.CompareAndSwapUint64(&v.val, oldBits, newBits) {
			return value
		}
	}
}
