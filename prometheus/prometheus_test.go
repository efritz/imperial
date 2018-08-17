package prometheus

//go:generate go-mockgen github.com/efritz/imperial/prometheus -i Registry -o mock_registry_test.go -f

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
	model "github.com/prometheus/client_model/go"

	"github.com/efritz/imperial/base"
)

type PrometheusSuite struct{}

func (s *PrometheusSuite) TestAddCounter(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns"),
		base.WithSubsystem("ss"),
	))

	reporter.AddCounter("foo", 1)
	reporter.AddCounter("bar", 2, base.WithHelp("h2"))
	reporter.AddCounter("foo", 3, base.WithHelp("h1"))
	reporter.AddCounter("foo", 4)
	reporter.AddCounter("baz", 5)

	Expect(registry.RegisterFuncCallCount()).To(Equal(3))

	counter1 := registry.RegisterFuncCallParams()[0].Arg0
	assertDesc(counter1, nil, `fqName: "ns_ss_foo"`)
	assertDesc(counter1, nil, `help: "<no help>"`)
	assertValue(counter1, nil, 8)

	counter2 := registry.RegisterFuncCallParams()[1].Arg0
	assertDesc(counter2, nil, `fqName: "ns_ss_bar"`)
	assertDesc(counter2, nil, `help: "h2"`)
	assertValue(counter2, nil, 2)

	counter3 := registry.RegisterFuncCallParams()[2].Arg0
	assertDesc(counter3, nil, `fqName: "ns_ss_baz"`)
	assertDesc(counter3, nil, `help: "<no help>"`)
	assertValue(counter3, nil, 5)
}

func (s *PrometheusSuite) TestAddCounterNamespaceSubsystem(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
	))

	reporter.AddCounter("foo", 1)
	reporter.AddCounter("foo", 2)
	reporter.AddCounter("foo", 3, base.WithNamespace("ns3"))
	reporter.AddCounter("foo", 4, base.WithSubsystem("ss4"))
	reporter.AddCounter("foo", 5, base.WithNamespace("ns5"), base.WithSubsystem("ss5"))

	Expect(registry.RegisterFuncCallCount()).To(Equal(4))
	assertValue(registry.RegisterFuncCallParams()[0].Arg0, nil, 3)
	assertValue(registry.RegisterFuncCallParams()[1].Arg0, nil, 3)
	assertValue(registry.RegisterFuncCallParams()[2].Arg0, nil, 4)
	assertValue(registry.RegisterFuncCallParams()[3].Arg0, nil, 5)
}

func (s *PrometheusSuite) TestAddCounterWithAttributes(t sweet.T) {
	var (
		exposed = []string{"status_code", "method"}
		attrs1  = map[string]string{"status_code": "200", "method": "GET"}
		attrs2  = map[string]string{"status_code": "200", "method": "POST"}
		attrs3  = map[string]string{"status_code": "404", "method": "GET"}
	)

	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
		base.WithExposedAttributes(exposed...),
	))

	reporter.AddCounter("foo", 1)
	reporter.AddCounter("foo", 2)
	reporter.AddCounter("foo", 3, base.WithAttributes(attrs1))
	reporter.AddCounter("foo", 4, base.WithAttributes(attrs2))
	reporter.AddCounter("foo", 5, base.WithAttributes(attrs3))

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	counter := registry.RegisterFuncCallParams()[0].Arg0
	assertValue(counter, makeLabels(exposed, nil), 3)
	assertValue(counter, makeLabels(exposed, attrs1), 3)
	assertValue(counter, makeLabels(exposed, attrs2), 4)
	assertValue(counter, makeLabels(exposed, attrs3), 5)
}

func (s *PrometheusSuite) TestAddCounterNegative(t sweet.T) {
	reporter, _ := makeReporter()

	Expect(func() {
		reporter.AddCounter("foo", -1)
	}).To(Panic())
}

func (s *PrometheusSuite) TestAddGauge(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns"),
		base.WithSubsystem("ss"),
	))

	reporter.AddGauge("foo", 1)
	reporter.AddGauge("bar", 2, base.WithHelp("h2"))
	reporter.AddGauge("foo", 3, base.WithHelp("h1"))
	reporter.AddGauge("foo", -6)
	reporter.AddGauge("baz", 5)

	Expect(registry.RegisterFuncCallCount()).To(Equal(3))

	gauge1 := registry.RegisterFuncCallParams()[0].Arg0
	assertValue(gauge1, nil, -2)
	assertDesc(gauge1, nil, `fqName: "ns_ss_foo"`)
	assertDesc(gauge1, nil, `help: "<no help>"`)

	gauge2 := registry.RegisterFuncCallParams()[1].Arg0
	assertValue(gauge2, nil, 2)
	assertDesc(gauge2, nil, `fqName: "ns_ss_bar"`)
	assertDesc(gauge2, nil, `help: "h2"`)

	gauge3 := registry.RegisterFuncCallParams()[2].Arg0
	assertValue(gauge3, nil, 5)
	assertDesc(gauge3, nil, `fqName: "ns_ss_baz"`)
	assertDesc(gauge3, nil, `help: "<no help>"`)
}

func (s *PrometheusSuite) TestAddGaugeNamespaceSubsystem(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
	))

	reporter.AddGauge("foo", 1)
	reporter.AddGauge("foo", 2)
	reporter.AddGauge("foo", 3, base.WithNamespace("ns3"))
	reporter.AddGauge("foo", 4, base.WithSubsystem("ss4"))
	reporter.AddGauge("foo", 5, base.WithNamespace("ns5"), base.WithSubsystem("ss5"))

	Expect(registry.RegisterFuncCallCount()).To(Equal(4))
	assertValue(registry.RegisterFuncCallParams()[0].Arg0, nil, 3)
	assertValue(registry.RegisterFuncCallParams()[1].Arg0, nil, 3)
	assertValue(registry.RegisterFuncCallParams()[2].Arg0, nil, 4)
	assertValue(registry.RegisterFuncCallParams()[3].Arg0, nil, 5)
}

func (s *PrometheusSuite) TestAddGaugeWithAttributes(t sweet.T) {
	var (
		exposed = []string{"status_code", "method"}
		attrs1  = map[string]string{"status_code": "200", "method": "GET"}
		attrs2  = map[string]string{"status_code": "200", "method": "POST"}
		attrs3  = map[string]string{"status_code": "404", "method": "GET"}
	)

	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
		base.WithExposedAttributes(exposed...),
	))

	reporter.AddGauge("foo", 1)
	reporter.AddGauge("foo", 2)
	reporter.AddGauge("foo", 3, base.WithAttributes(attrs1))
	reporter.AddGauge("foo", 4, base.WithAttributes(attrs2))
	reporter.AddGauge("foo", 5, base.WithAttributes(attrs3))

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	gauge := registry.RegisterFuncCallParams()[0].Arg0
	assertValue(gauge, makeLabels(exposed, nil), 3)
	assertValue(gauge, makeLabels(exposed, attrs1), 3)
	assertValue(gauge, makeLabels(exposed, attrs2), 4)
	assertValue(gauge, makeLabels(exposed, attrs3), 5)
}

func (s *PrometheusSuite) TestSetGauge(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns"),
		base.WithSubsystem("ss"),
	))

	reporter.SetGauge("foo", 1)
	reporter.SetGauge("bar", 2, base.WithHelp("h2"))
	reporter.SetGauge("foo", 3, base.WithHelp("h1"))
	reporter.SetGauge("foo", -6)
	reporter.SetGauge("baz", 5)

	Expect(registry.RegisterFuncCallCount()).To(Equal(3))

	gauge1 := registry.RegisterFuncCallParams()[0].Arg0
	assertValue(gauge1, nil, -6)
	assertDesc(gauge1, nil, `fqName: "ns_ss_foo"`)
	assertDesc(gauge1, nil, `help: "<no help>"`)

	gauge2 := registry.RegisterFuncCallParams()[1].Arg0
	assertValue(gauge2, nil, 2)
	assertDesc(gauge2, nil, `fqName: "ns_ss_bar"`)
	assertDesc(gauge2, nil, `help: "h2"`)

	gauge3 := registry.RegisterFuncCallParams()[2].Arg0
	assertValue(gauge3, nil, 5)
	assertDesc(gauge3, nil, `fqName: "ns_ss_baz"`)
	assertDesc(gauge3, nil, `help: "<no help>"`)
}

func (s *PrometheusSuite) TestSetGaugeNamespaceSubsystem(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
	))

	reporter.SetGauge("foo", 1)
	reporter.SetGauge("foo", 2)
	reporter.SetGauge("foo", 3, base.WithNamespace("ns3"))
	reporter.SetGauge("foo", 4, base.WithSubsystem("ss4"))
	reporter.SetGauge("foo", 5, base.WithNamespace("ns5"), base.WithSubsystem("ss5"))

	Expect(registry.RegisterFuncCallCount()).To(Equal(4))
	assertValue(registry.RegisterFuncCallParams()[0].Arg0, nil, 2)
	assertValue(registry.RegisterFuncCallParams()[1].Arg0, nil, 3)
	assertValue(registry.RegisterFuncCallParams()[2].Arg0, nil, 4)
	assertValue(registry.RegisterFuncCallParams()[3].Arg0, nil, 5)
}

func (s *PrometheusSuite) TestSetGaugeWithAttributes(t sweet.T) {
	var (
		exposed = []string{"status_code", "method"}
		attrs1  = map[string]string{"status_code": "200", "method": "GET"}
		attrs2  = map[string]string{"status_code": "200", "method": "POST"}
		attrs3  = map[string]string{"status_code": "404", "method": "GET"}
	)

	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
		base.WithExposedAttributes(exposed...),
	))

	reporter.SetGauge("foo", 1)
	reporter.SetGauge("foo", 2)
	reporter.SetGauge("foo", 3, base.WithAttributes(attrs1))
	reporter.SetGauge("foo", 4, base.WithAttributes(attrs2))
	reporter.SetGauge("foo", 5, base.WithAttributes(attrs3))

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	gauge := registry.RegisterFuncCallParams()[0].Arg0
	assertValue(gauge, makeLabels(exposed, nil), 2)
	assertValue(gauge, makeLabels(exposed, attrs1), 3)
	assertValue(gauge, makeLabels(exposed, attrs2), 4)
	assertValue(gauge, makeLabels(exposed, attrs3), 5)
}

func (s *PrometheusSuite) TestObserveHistogram(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns"),
		base.WithSubsystem("ss"),
	))

	reporter.ObserveHistogram("foo", 5)
	reporter.ObserveHistogram("bar", 4, base.WithHelp("h2"))
	reporter.ObserveHistogram("foo", 3, base.WithHelp("h1"))
	reporter.ObserveHistogram("foo", 2)

	for _, x := range []float64{0.005, 0.005, 0.005, 0.01, 0.02, 1, 8} {
		reporter.ObserveHistogram("baz", x)
	}

	Expect(registry.RegisterFuncCallCount()).To(Equal(3))

	histogram1 := registry.RegisterFuncCallParams()[0].Arg0
	assertDesc(histogram1, nil, `fqName: "ns_ss_foo"`)
	assertDesc(histogram1, nil, `help: "<no help>"`)
	assertBucketCount(histogram1, nil, 1, 0)
	assertBucketCount(histogram1, nil, 2.5, 1)
	assertBucketCount(histogram1, nil, 5, 3)

	histogram2 := registry.RegisterFuncCallParams()[1].Arg0
	assertDesc(histogram2, nil, `fqName: "ns_ss_bar"`)
	assertDesc(histogram2, nil, `help: "h2"`)
	assertBucketCount(histogram2, nil, 2.5, 0)
	assertBucketCount(histogram2, nil, 5, 1)

	histogram3 := registry.RegisterFuncCallParams()[2].Arg0
	assertDesc(histogram3, nil, `fqName: "ns_ss_baz"`)
	assertDesc(histogram3, nil, `help: "<no help>"`)
	assertBucketCount(histogram3, nil, 0.005, 3)
	assertBucketCount(histogram3, nil, 0.01, 4)
	assertBucketCount(histogram3, nil, 0.025, 5)
	assertBucketCount(histogram3, nil, 0.05, 5)
	assertBucketCount(histogram3, nil, 0.1, 5)
	assertBucketCount(histogram3, nil, 0.25, 5)
	assertBucketCount(histogram3, nil, 0.5, 5)
	assertBucketCount(histogram3, nil, 1, 6)
	assertBucketCount(histogram3, nil, 2.5, 6)
	assertBucketCount(histogram3, nil, 5, 6)
	assertBucketCount(histogram3, nil, 10, 7)
}

func (s *PrometheusSuite) TestObserveHistogramNamespaceSubsystem(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
	))

	reporter.ObserveHistogram("foo", 10)
	reporter.ObserveHistogram("foo", 10)
	reporter.ObserveHistogram("foo", 10, base.WithNamespace("ns3"))
	reporter.ObserveHistogram("foo", 10, base.WithSubsystem("ss4"))
	reporter.ObserveHistogram("foo", 10, base.WithNamespace("ns5"), base.WithSubsystem("ss5"))

	Expect(registry.RegisterFuncCallCount()).To(Equal(4))
	assertBucketCount(registry.RegisterFuncCallParams()[0].Arg0, nil, 10, 2)
	assertBucketCount(registry.RegisterFuncCallParams()[1].Arg0, nil, 10, 1)
	assertBucketCount(registry.RegisterFuncCallParams()[2].Arg0, nil, 10, 1)
	assertBucketCount(registry.RegisterFuncCallParams()[3].Arg0, nil, 10, 1)
}

func (s *PrometheusSuite) TestObserveHistogramWithLabels(t sweet.T) {
	var (
		exposed = []string{"status_code", "method"}
		attrs1  = map[string]string{"status_code": "200", "method": "GET"}
		attrs2  = map[string]string{"status_code": "200", "method": "POST"}
		attrs3  = map[string]string{"status_code": "404", "method": "GET"}
	)

	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
		base.WithExposedAttributes(exposed...),
	))

	reporter.ObserveHistogram("foo", 0.005)
	reporter.ObserveHistogram("foo", 0.02, base.WithAttributes(attrs1))
	reporter.ObserveHistogram("foo", 1, base.WithAttributes(attrs2))
	reporter.ObserveHistogram("foo", 8, base.WithAttributes(attrs3))

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	histogram := registry.RegisterFuncCallParams()[0].Arg0
	assertBucketCount(histogram, makeLabels(exposed, nil), 0.005, 1)
	assertBucketCount(histogram, makeLabels(exposed, attrs1), 0.005, 0)
	assertBucketCount(histogram, makeLabels(exposed, attrs1), 0.02, 1)
	assertBucketCount(histogram, makeLabels(exposed, attrs2), 0.02, 0)
	assertBucketCount(histogram, makeLabels(exposed, attrs2), 1, 1)
	assertBucketCount(histogram, makeLabels(exposed, attrs3), 1, 0)
	assertBucketCount(histogram, makeLabels(exposed, attrs3), 10, 1)
}

func (s *PrometheusSuite) TestObserveHistogramWithCustomBuckets(t sweet.T) {
	reporter, registry := makeReporter()

	var (
		buckets = []float64{1, 2, 4, 8, 16}
		samples = []float64{1, 2, 3, 4, 5, 6, 7, 8, 10}
	)

	for _, x := range samples {
		reporter.ObserveHistogram("foo", x, base.WithBuckets(buckets))
	}

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	histogram := registry.RegisterFuncCallParams()[0].Arg0
	assertBucketCount(histogram, nil, 1, 1)
	assertBucketCount(histogram, nil, 2, 2)
	assertBucketCount(histogram, nil, 4, 4)
	assertBucketCount(histogram, nil, 8, 8)
	assertBucketCount(histogram, nil, 16, 9)
}

func (s *PrometheusSuite) TestObserveSummary(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns"),
		base.WithSubsystem("ss"),
	))

	reporter.ObserveSummary("foo", 1)
	reporter.ObserveSummary("foo", 2, base.WithHelp("h1"))
	reporter.ObserveSummary("foo", 3)
	reporter.ObserveSummary("bar", 2, base.WithHelp("h2"))

	for i := 100; i < 400; i++ {
		reporter.ObserveSummary("baz", float64(i))
	}

	Expect(registry.RegisterFuncCallCount()).To(Equal(3))

	summary1 := registry.RegisterFuncCallParams()[0].Arg0
	assertQuantileValue(summary1, nil, 0.5, 2)
	assertQuantileValue(summary1, nil, 0.9, 3)
	assertQuantileValue(summary1, nil, 0.99, 3)
	assertDesc(summary1, nil, `fqName: "ns_ss_foo"`)
	assertDesc(summary1, nil, `help: "<no help>"`)

	summary2 := registry.RegisterFuncCallParams()[1].Arg0
	assertDesc(summary2, nil, `fqName: "ns_ss_bar"`)
	assertDesc(summary2, nil, `help: "h2"`)
	assertQuantileValue(summary2, nil, 0.5, 2)
	assertQuantileValue(summary2, nil, 0.9, 2)
	assertQuantileValue(summary2, nil, 0.99, 2)

	summary3 := registry.RegisterFuncCallParams()[2].Arg0
	assertDesc(summary3, nil, `fqName: "ns_ss_baz"`)
	assertDesc(summary3, nil, `help: "<no help>"`)
	assertQuantileValue(summary3, nil, 0.5, 249)
	assertQuantileValue(summary3, nil, 0.9, 369)
	assertQuantileValue(summary3, nil, 0.99, 396)
}

func (s *PrometheusSuite) TestObserveSummaryNamespaceSubsystem(t sweet.T) {
	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
	))

	reporter.ObserveSummary("foo", 10)
	reporter.ObserveSummary("foo", 20)
	reporter.ObserveSummary("foo", 30)
	reporter.ObserveSummary("foo", 25, base.WithNamespace("ns3"))
	reporter.ObserveSummary("foo", 35, base.WithSubsystem("ss4"))
	reporter.ObserveSummary("foo", 45, base.WithNamespace("ns5"), base.WithSubsystem("ss5"))

	Expect(registry.RegisterFuncCallCount()).To(Equal(4))
	assertQuantileValue(registry.RegisterFuncCallParams()[0].Arg0, nil, 0.5, 20)
	assertQuantileValue(registry.RegisterFuncCallParams()[1].Arg0, nil, 0.5, 25)
	assertQuantileValue(registry.RegisterFuncCallParams()[2].Arg0, nil, 0.5, 35)
	assertQuantileValue(registry.RegisterFuncCallParams()[3].Arg0, nil, 0.5, 45)
}

func (s *PrometheusSuite) TestObserveSummaryWithLabels(t sweet.T) {
	var (
		exposed = []string{"status_code", "method"}
		attrs1  = map[string]string{"status_code": "200", "method": "GET"}
		attrs2  = map[string]string{"status_code": "200", "method": "POST"}
		attrs3  = map[string]string{"status_code": "404", "method": "GET"}
	)

	reporter, registry := makeReporter(WithReportConfigs(
		base.WithNamespace("ns1"),
		base.WithSubsystem("ss1"),
		base.WithExposedAttributes(exposed...),
	))

	reporter.ObserveSummary("foo", 20)
	reporter.ObserveSummary("foo", 30, base.WithAttributes(attrs1))
	reporter.ObserveSummary("foo", 40, base.WithAttributes(attrs2))
	reporter.ObserveSummary("foo", 50, base.WithAttributes(attrs3))

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	summary := registry.RegisterFuncCallParams()[0].Arg0
	assertQuantileValue(summary, makeLabels(exposed, nil), 0.5, 20)
	assertQuantileValue(summary, makeLabels(exposed, attrs1), 0.5, 30)
	assertQuantileValue(summary, makeLabels(exposed, attrs2), 0.5, 40)
	assertQuantileValue(summary, makeLabels(exposed, attrs3), 0.5, 50)
}

func (s *PrometheusSuite) TestObserveSummaryWithCustomQuantiles(t sweet.T) {
	reporter, registry := makeReporter()

	for _, x := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16} {
		reporter.ObserveSummary("foo", x, base.WithQuantiles(map[float64]float64{
			0.25:  0.01,
			0.50:  0.01,
			0.75:  0.01,
		}))
	}

	Expect(registry.RegisterFuncCallCount()).To(Equal(1))
	histogram := registry.RegisterFuncCallParams()[0].Arg0
	assertQuantileValue(histogram, nil, 0.25, 4)
	assertQuantileValue(histogram, nil, 0.50, 8)
	assertQuantileValue(histogram, nil, 0.75, 12)
}

func (s *PrometheusSuite) TestMakeLabels(t sweet.T) {
	Expect(makeLabelsFromConfig(&base.ReportConfig{
		ExposedAttributes: []string{"a", "b"},
		Attributes:        map[string]string{"a": "foo", "b": "bar"},
	})).To(Equal(prometheus.Labels{
		"a": "foo",
		"b": "bar",
	}))

	Expect(makeLabelsFromConfig(&base.ReportConfig{
		ExposedAttributes: []string{"a", "c"},
		Attributes:        map[string]string{"a": "foo", "b": "bar"},
	})).To(Equal(prometheus.Labels{
		"a": "foo",
		"c": "",
	}))

	Expect(makeLabelsFromConfig(&base.ReportConfig{
		ExposedAttributes: []string{"a", "b"},
		Attributes:        map[string]string{"a": "foo"},
	})).To(Equal(prometheus.Labels{
		"a": "foo",
		"b": "",
	}))
}

//
// Constructors

func makeReporter(configs ...ConfigFunc) (*Reporter, *MockRegistry) {
	var (
		registry = NewMockRegistry()
		reporter = NewReporter(append(
			[]ConfigFunc{
				WithRegistry(registry),
			},
			configs...,
		)...)
	)

	return reporter, registry
}

//
// Checkers

func assertDesc(collector prometheus.Collector, labels prometheus.Labels, substring string) {
	Expect(extractMetric(collector, labels).Desc().String()).To(ContainSubstring(substring))
}

func assertValue(collector prometheus.Collector, labels prometheus.Labels, expected interface{}) {
	metric := &model.Metric{}
	Expect(extractMetric(collector, labels).Write(metric)).To(BeNil())
	Expect(metric.String()).To(ContainSubstring("value:%v", expected))
}

func assertBucketCount(collector prometheus.Collector, labels prometheus.Labels, bound float64, expected interface{}) {
	metric := &model.Metric{}
	Expect(extractMetric(collector, labels).Write(metric)).To(BeNil())
	Expect(metric.String()).To(ContainSubstring("cumulative_count:%v upper_bound:%v", expected, bound))
}

func assertQuantileValue(collector prometheus.Collector, labels prometheus.Labels, quantile float64, expected interface{}) {
	metric := &model.Metric{}
	Expect(extractMetric(collector, labels).Write(metric)).To(BeNil())
	Expect(metric.String()).To(ContainSubstring("quantile:%v value:%v", quantile, expected))
}

func extractMetric(collector prometheus.Collector, labels prometheus.Labels) prometheus.Metric {
	switch v := collector.(type) {
	case *prometheus.CounterVec:
		metric, _ := v.GetMetricWith(labels)
		return metric

	case *prometheus.GaugeVec:
		metric, _ := v.GetMetricWith(labels)
		return metric

	case *prometheus.HistogramVec:
		metric, _ := v.GetMetricWith(labels)
		return metric.(prometheus.Metric)

	case *prometheus.SummaryVec:
		metric, _ := v.GetMetricWith(labels)
		return metric.(prometheus.Metric)
	}

	return nil
}
