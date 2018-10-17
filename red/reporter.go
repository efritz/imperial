package red

import (
	"fmt"

	"github.com/efritz/imperial/base"
)

type (
	Reporter struct {
		reporter      base.Reporter
		logger        base.Logger
		prefixConfigs map[string]*PrefixConfig
	}
)

func NewReporter(reporter base.Reporter, configs ...ConfigFunc) *Reporter {
	config := newConfig()
	for _, f := range configs {
		f(config)
	}

	return &Reporter{
		reporter:      reporter,
		logger:        config.logger,
		prefixConfigs: config.prefixConfigs,
	}
}

func (r *Reporter) Register() {
	for prefix, config := range r.prefixConfigs {
		r.reporter.RegisterCounter(fmt.Sprintf("%s-request", prefix), r.requestMetricConfigs(config, nil)...)
		r.reporter.RegisterCounter(fmt.Sprintf("%s-error", prefix), r.errorMetricConfigs(config, nil)...)
		r.reporter.RegisterHistogram(fmt.Sprintf("%s-duration", prefix), r.durationMetricConfigs(config, nil)...)
	}
}

func (r *Reporter) ReportRequest(prefix string, configs ...base.ConfigFunc) {
	config, ok := r.prefixConfigs[prefix]
	if !ok {
		r.logger.Printf("No configuration registered for prefix '%s'", prefix)
		return
	}

	r.reporter.AddCounter(
		fmt.Sprintf("%s-request", prefix),
		1,
		r.requestMetricConfigs(config, configs)...,
	)
}

func (r *Reporter) ReportError(prefix string, err error, configs ...base.ConfigFunc) {
	config, ok := r.prefixConfigs[prefix]
	if !ok {
		r.logger.Printf("No configuration registered for prefix '%s'", prefix)
		return
	}

	code, ok := config.errorInterpreter(err)
	if !ok {
		return
	}

	attributes := map[string]string{
		"code": code,
	}

	configs = append(
		configs,
		base.WithAttributes(attributes),
	)

	r.reporter.AddCounter(
		fmt.Sprintf("%s-error", prefix),
		1,
		r.durationMetricConfigs(config, configs)...,
	)
}

func (r *Reporter) ReportDuration(prefix string, duration float64, configs ...base.ConfigFunc) {
	config, ok := r.prefixConfigs[prefix]
	if !ok {
		r.logger.Printf("No configuration registered for prefix '%s'", prefix)
		return
	}

	r.reporter.ObserveHistogram(
		fmt.Sprintf("%s-duration", prefix),
		duration,
		r.durationMetricConfigs(config, configs)...,
	)
}

//
//

func (r *Reporter) requestMetricConfigs(config *PrefixConfig, configs []base.ConfigFunc) []base.ConfigFunc {
	defaultConfigs := []base.ConfigFunc{
		base.WithHelp(""), // TODO
	}

	return append(append(defaultConfigs, config.configs...), configs...)
}

func (r *Reporter) errorMetricConfigs(config *PrefixConfig, configs []base.ConfigFunc) []base.ConfigFunc {
	defaultConfigs := []base.ConfigFunc{
		base.WithExposedAttributes("code"),
		base.WithHelp(""), // TODO
	}

	return append(append(defaultConfigs, config.configs...), configs...)
}

func (r *Reporter) durationMetricConfigs(config *PrefixConfig, configs []base.ConfigFunc) []base.ConfigFunc {
	defaultConfigs := []base.ConfigFunc{
		base.WithBuckets(config.buckets),
		base.WithHelp(""), // TODO
	}

	return append(append(defaultConfigs, config.configs...), configs...)
}
