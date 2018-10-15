package base

//go:generate go-mockgen -f github.com/efritz/imperial/base -i SimpleReporter

import (
	"math"
	"sync"
	"sync/atomic"

	"github.com/prometheus/client_golang/prometheus"
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

func (r *SimpleReporterShim) RegisterCounter(name string, configs ...ConfigFunc)   {}
func (r *SimpleReporterShim) RegisterGauge(name string, configs ...ConfigFunc)     {}
func (r *SimpleReporterShim) RegisterHistogram(name string, configs ...ConfigFunc) {}
func (r *SimpleReporterShim) RegisterSummary(name string, configs ...ConfigFunc)   {}

func (r *SimpleReporterShim) AddCounter(name string, value float64, configs ...ConfigFunc) {
	if value < 0 {
		panic("counter cannot decrease in value")
	}

	key := makeKey(name, configs)
	r.ensureValue(r.counters, key)
	r.reporter.Report(name, r.counters[key].Add(value), configs...)
}

func (r *SimpleReporterShim) AddGauge(name string, value float64, configs ...ConfigFunc) {
	key := makeKey(name, configs)
	r.ensureValue(r.gauges, key)
	r.reporter.Report(name, r.gauges[key].Add(value), configs...)
}

func (r *SimpleReporterShim) SetGauge(name string, value float64, configs ...ConfigFunc) {
	key := makeKey(name, configs)
	r.ensureValue(r.gauges, key)
	r.reporter.Report(name, r.gauges[key].Set(value), configs...)
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

//
//

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

//
//

func makeKey(name string, configs []ConfigFunc) string {
	config := ApplyConfigs(configs, nil)

	return prometheus.BuildFQName(
		config.Namespace,
		config.Subsystem,
		name,
	)
}
