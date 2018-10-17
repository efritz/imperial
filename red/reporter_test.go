package red

import (
	"fmt"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"

	"github.com/efritz/imperial/base"
)

type ReporterSuite struct{}

func (s *ReporterSuite) TestRegister(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(
		reporter,
		WithPrefixConfig("foo", NewPrefixConfig()),
		WithPrefixConfig("bar", NewPrefixConfig(WithPrefixBuckets([]float64{1, 2, 3}))),
		WithPrefixConfig("baz", NewPrefixConfig(WithPrefixBuckets([]float64{1, 2, 4}))),
	)

	redReporter.Register()

	counterNames := []string{}
	for _, param := range reporter.RegisterCounterFuncCallParams() {
		counterNames = append(counterNames, param.Arg0)
	}

	histogramNames := []string{}
	for _, param := range reporter.RegisterHistogramFuncCallParams() {
		histogramNames = append(histogramNames, param.Arg0)
	}

	Expect(counterNames).To(ConsistOf("foo", "foo-error", "bar", "bar-error", "baz", "baz-error"))
	Expect(histogramNames).To(ConsistOf("foo-duration", "bar-duration", "baz-duration"))
}

func (s *ReporterSuite) TestReportRequest(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter, WithPrefixConfig("foo", NewPrefixConfig()))
	redReporter.ReportRequest("foo")

	Expect(reporter.AddCounterFuncCallCount()).To(Equal(1))
	Expect(reporter.AddCounterFuncCallParams()[0].Arg0).To(Equal("foo-request"))
	Expect(reporter.AddCounterFuncCallParams()[0].Arg1).To(Equal(float64(1)))
}

func (s *ReporterSuite) TestReportRequestUnknownPrefix(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter)
	redReporter.ReportRequest("foo")
	Expect(reporter.AddCounterFuncCallCount()).To(Equal(0))
}

func (s *ReporterSuite) TestReportError(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter, WithPrefixConfig("foo", NewPrefixConfig()))
	redReporter.ReportError("foo", fmt.Errorf("utoh"))

	Expect(reporter.AddCounterFuncCallCount()).To(Equal(1))
	Expect(reporter.AddCounterFuncCallParams()[0].Arg0).To(Equal("foo-error"))
	Expect(reporter.AddCounterFuncCallParams()[0].Arg1).To(Equal(float64(1)))
}

func (s *ReporterSuite) TestReportErrorUnknownPrefix(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter)
	redReporter.ReportError("foo", fmt.Errorf("utoh"))
	Expect(reporter.AddCounterFuncCallCount()).To(Equal(0))
}

func (s *ReporterSuite) TestReportNoMatch(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter, WithPrefixConfig("foo", NewPrefixConfig()))
	redReporter.ReportError("foo", nil)
	Expect(reporter.AddCounterFuncCallCount()).To(Equal(0))
}

func (s *ReporterSuite) TestReportDuration(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter, WithPrefixConfig("foo", NewPrefixConfig()))
	redReporter.ReportDuration("foo", 3.1234)

	Expect(reporter.ObserveHistogramFuncCallCount()).To(Equal(1))
	Expect(reporter.ObserveHistogramFuncCallParams()[0].Arg0).To(Equal("foo-duration"))
	Expect(reporter.ObserveHistogramFuncCallParams()[0].Arg1).To(Equal(float64(3.1234)))
}

func (s *ReporterSuite) TestReportDurationUnknownPRefix(t sweet.T) {
	reporter := base.NewMockReporter()
	redReporter := NewReporter(reporter)
	redReporter.ReportDuration("foo", 3.1234)
	Expect(reporter.ObserveHistogramFuncCallCount()).To(Equal(0))
}
