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
		r.reporter.RegisterCounter(prefix, r.requestMetricConfigs(config)...)
		r.reporter.RegisterCounter(fmt.Sprintf("%s-error", prefix), r.errorMetricConfigs(config)...)
		r.reporter.RegisterHistogram(fmt.Sprintf("%s-duration", prefix), r.durationMetricConfigs(config)...)
	}
}

func (r *Reporter) ReportRequest(prefix string) {
	config, ok := r.prefixConfigs[prefix]
	if !ok {
		r.logger.Printf("No configuration registered for prefix '%s'", prefix)
		return
	}

	r.reporter.AddCounter(fmt.Sprintf("%s-request",prefix), 1, r.requestMetricConfigs(config)...)
}

func (r *Reporter) ReportError(prefix string, err error) {
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

	configs := append(
		[]base.ConfigFunc{base.WithAttributes(attributes)},
		r.errorMetricConfigs(config)...,
	)

	r.reporter.AddCounter(fmt.Sprintf("%s-error", prefix), 1, configs...)
}

func (r *Reporter) ReportDuration(prefix string, duration float64) {
	config, ok := r.prefixConfigs[prefix]
	if !ok {
		r.logger.Printf("No configuration registered for prefix '%s'", prefix)
		return
	}

	r.reporter.ObserveHistogram(
		fmt.Sprintf("%s-duration", prefix),
		duration,
		r.durationMetricConfigs(config)...,
	)
}

//
//

func (r *Reporter) requestMetricConfigs(config *PrefixConfig) []base.ConfigFunc {
	return []base.ConfigFunc{
		base.WithHelp(""), // TODO
	}
}

func (r *Reporter) errorMetricConfigs(config *PrefixConfig) []base.ConfigFunc {
	return []base.ConfigFunc{
		base.WithExposedAttributes("code"),
		base.WithHelp(""), // TODO
	}
}

func (r *Reporter) durationMetricConfigs(config *PrefixConfig) []base.ConfigFunc {
	return []base.ConfigFunc{
		base.WithBuckets(config.buckets),
		base.WithHelp(""), // TODO
	}
}
