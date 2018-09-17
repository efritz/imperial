package prometheus

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/efritz/imperial/base"
)

type (
	Reporter struct {
		logger                     base.Logger
		configs                    []base.ConfigFunc
		registry                   Registry
		handlerMaxRequestsInFlight int
		handlerTimeout             time.Duration
		counters                   map[string]prometheus.Collector
		gauges                     map[string]prometheus.Collector
		histograms                 map[string]prometheus.Collector
		summaries                  map[string]prometheus.Collector
		mutex                      *sync.RWMutex
	}

	Registry interface {
		prometheus.Registerer
		prometheus.Gatherer
	}

	CollectorFactory func(string, *base.ReportConfig) prometheus.Collector
)

var _ base.Reporter = &Reporter{}

func NewReporter(configs ...ConfigFunc) *Reporter {
	config := newConfig()
	for _, f := range configs {
		f(config)
	}

	return &Reporter{
		logger:                     config.logger,
		configs:                    config.configs,
		registry:                   config.registry,
		handlerMaxRequestsInFlight: config.handlerMaxRequestsInFlight,
		handlerTimeout:             config.handlerTimeout,
		counters:                   map[string]prometheus.Collector{},
		gauges:                     map[string]prometheus.Collector{},
		histograms:                 map[string]prometheus.Collector{},
		summaries:                  map[string]prometheus.Collector{},
		mutex:                      &sync.RWMutex{},
	}
}

func (r *Reporter) Handler() http.Handler {
	return promhttp.HandlerFor(r.registry, promhttp.HandlerOpts{
		ErrorLog:            &loggerShim{r.logger},
		MaxRequestsInFlight: r.handlerMaxRequestsInFlight,
		Timeout:             r.handlerTimeout,
	})
}

func (r *Reporter) AddCounter(name string, value float64, configs ...base.ConfigFunc) {
	counter, err := r.ensureCounter(name, base.ApplyConfigs(r.configs, configs))
	if err != nil {
		r.logger.Printf("Error registering counter (%s)", err.Error())
		return
	}

	counter.Add(value)
}

func (r *Reporter) AddGauge(name string, value float64, configs ...base.ConfigFunc) {
	gauge, err := r.ensureGauge(name, base.ApplyConfigs(r.configs, configs))
	if err != nil {
		r.logger.Printf("Error registering gauge (%s)", err.Error())
		return
	}

	gauge.Add(value)
}

func (r *Reporter) SetGauge(name string, value float64, configs ...base.ConfigFunc) {
	gauge, err := r.ensureGauge(name, base.ApplyConfigs(r.configs, configs))
	if err != nil {
		r.logger.Printf("Error registering gauge (%s)", err.Error())
		return
	}

	gauge.Set(value)
}

func (r *Reporter) ObserveHistogram(name string, value float64, configs ...base.ConfigFunc) {
	histogram, err := r.ensureHistogram(name, base.ApplyConfigs(r.configs, configs))
	if err != nil {
		r.logger.Printf("Error registering histogram (%s)", err.Error())
		return
	}

	histogram.Observe(value)
}

func (r *Reporter) ObserveSummary(name string, value float64, configs ...base.ConfigFunc) {
	summary, err := r.ensureSummary(name, base.ApplyConfigs(r.configs, configs))
	if err != nil {
		r.logger.Printf("Error registering summary (%s)", err.Error())
		return
	}

	summary.Observe(value)
}

func (r *Reporter) Shutdown() {
	// nothing to do
}

//
//

func (r *Reporter) ensureCounter(name string, config *base.ReportConfig) (prometheus.Counter, error) {
	vec, err := r.ensureVec(r.counters, name, config, makeCounterVec)
	if err != nil {
		return nil, err
	}

	return vec.(*prometheus.CounterVec).GetMetricWith(makeLabelsFromConfig(config))
}

func (r *Reporter) ensureGauge(name string, config *base.ReportConfig) (prometheus.Gauge, error) {
	vec, err := r.ensureVec(r.gauges, name, config, makeGaugeVec)
	if err != nil {
		return nil, err
	}

	return vec.(*prometheus.GaugeVec).GetMetricWith(makeLabelsFromConfig(config))
}

func (r *Reporter) ensureHistogram(name string, config *base.ReportConfig) (prometheus.Observer, error) {
	vec, err := r.ensureVec(r.histograms, name, config, makeHistogramVec)
	if err != nil {
		return nil, err
	}

	return vec.(*prometheus.HistogramVec).GetMetricWith(makeLabelsFromConfig(config))
}

func (r *Reporter) ensureSummary(name string, config *base.ReportConfig) (prometheus.Observer, error) {
	vec, err := r.ensureVec(r.summaries, name, config, makeSummaryVec)
	if err != nil {
		return nil, err
	}

	return vec.(*prometheus.SummaryVec).GetMetricWith(makeLabelsFromConfig(config))
}

func (r *Reporter) ensureVec(
	m map[string]prometheus.Collector,
	name string,
	config *base.ReportConfig,
	factory CollectorFactory,
) (interface{}, error) {
	key := prometheus.BuildFQName(
		config.Namespace,
		config.Subsystem,
		name,
	)

	r.mutex.RLock()
	if vec, ok := m[key]; ok {
		r.mutex.RUnlock()
		return vec, nil
	}

	r.mutex.RUnlock()
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if vec, ok := m[key]; ok {
		return vec, nil
	}

	vec := factory(name, config)
	m[key] = vec

	if err := r.registry.Register(vec); err != nil {
		return nil, err
	}

	return vec, nil
}

//
//

func makeCounterVec(name string, config *base.ReportConfig) prometheus.Collector {
	opts := prometheus.CounterOpts{
		Name:      name,
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Help:      config.Help,
	}

	return prometheus.NewCounterVec(opts, config.ExposedAttributes)
}

func makeGaugeVec(name string, config *base.ReportConfig) prometheus.Collector {
	opts := prometheus.GaugeOpts{
		Name:      name,
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Help:      config.Help,
	}

	return prometheus.NewGaugeVec(opts, config.ExposedAttributes)
}

func makeHistogramVec(name string, config *base.ReportConfig) prometheus.Collector {
	opts := prometheus.HistogramOpts{
		Name:      name,
		Namespace: config.Namespace,
		Subsystem: config.Subsystem,
		Help:      config.Help,
		Buckets:   config.Buckets,
	}

	return prometheus.NewHistogramVec(opts, config.ExposedAttributes)
}

func makeSummaryVec(name string, config *base.ReportConfig) prometheus.Collector {
	opts := prometheus.SummaryOpts{
		Name:       name,
		Namespace:  config.Namespace,
		Subsystem:  config.Subsystem,
		Help:       config.Help,
		Objectives: config.Quantiles,
	}

	return prometheus.NewSummaryVec(opts, config.ExposedAttributes)
}

func makeLabelsFromConfig(config *base.ReportConfig) prometheus.Labels {
	return makeLabels(config.ExposedAttributes, config.Attributes)
}

func makeLabels(exposedAttributes []string, attributes map[string]string) prometheus.Labels {
	labels := prometheus.Labels{}
	for _, name := range exposedAttributes {
		labels[name] = attributes[name]
	}

	return labels
}
