package base

import (
	"math"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type SimpleReporterSuite struct{}

func (s *SimpleReporterSuite) TestAddCounter(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	reporter.AddCounter("foo", 1)
	reporter.AddCounter("foo", 2)
	reporter.AddCounter("bar", 3)
	reporter.AddCounter("foo", 4)

	Expect(simple.ReportFuncCallCount()).To(Equal(4))
	Expect(simple.ReportFuncCallParams()[0].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[0].Arg1).To(Equal(float64(1)))
	Expect(simple.ReportFuncCallParams()[1].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[1].Arg1).To(Equal(float64(3)))
	Expect(simple.ReportFuncCallParams()[2].Arg0).To(Equal("bar"))
	Expect(simple.ReportFuncCallParams()[2].Arg1).To(Equal(float64(3)))
	Expect(simple.ReportFuncCallParams()[3].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[3].Arg1).To(Equal(float64(7)))
}

func (s *SimpleReporterSuite) TestAddCounterNegative(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	Expect(func() { reporter.AddCounter("name", -1) }).To(Panic())
}

func (s *SimpleReporterSuite) TestAddGauge(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	reporter.AddGauge("foo", 1)
	reporter.AddGauge("foo", 2)
	reporter.AddGauge("bar", 3)
	reporter.AddGauge("foo", -1)

	Expect(simple.ReportFuncCallCount()).To(Equal(4))
	Expect(simple.ReportFuncCallParams()[0].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[0].Arg1).To(Equal(float64(1)))
	Expect(simple.ReportFuncCallParams()[1].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[1].Arg1).To(Equal(float64(3)))
	Expect(simple.ReportFuncCallParams()[2].Arg0).To(Equal("bar"))
	Expect(simple.ReportFuncCallParams()[2].Arg1).To(Equal(float64(3)))
	Expect(simple.ReportFuncCallParams()[3].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[3].Arg1).To(Equal(float64(2)))
}

func (s *SimpleReporterSuite) TestSetGauge(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	reporter.AddGauge("foo", 1)
	reporter.AddGauge("foo", 2)
	reporter.SetGauge("foo", 20)
	reporter.AddGauge("foo", 3)

	Expect(simple.ReportFuncCallCount()).To(Equal(4))
	Expect(simple.ReportFuncCallParams()[0].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[0].Arg1).To(Equal(float64(1)))
	Expect(simple.ReportFuncCallParams()[1].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[1].Arg1).To(Equal(float64(3)))
	Expect(simple.ReportFuncCallParams()[2].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[2].Arg1).To(Equal(float64(20)))
	Expect(simple.ReportFuncCallParams()[3].Arg0).To(Equal("foo"))
	Expect(simple.ReportFuncCallParams()[3].Arg1).To(Equal(float64(23)))
}

func (s *SimpleReporterSuite) TestObserveHistogram(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	for i := 0; i < 5; i++ {
		reporter.ObserveHistogram("name", float64(i*10+1))
		Expect(simple.ReportFuncCallCount()).To(Equal(i + 1))
		Expect(simple.ReportFuncCallParams()[i].Arg0).To(Equal("name"))
		Expect(simple.ReportFuncCallParams()[i].Arg1).To(Equal(float64(i*10 + 1)))
	}
}

func (s *SimpleReporterSuite) TestObserveSummary(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	for i := 0; i < 5; i++ {
		reporter.ObserveSummary("name", float64(i*10+1))
		Expect(simple.ReportFuncCallCount()).To(Equal(i + 1))
		Expect(simple.ReportFuncCallParams()[i].Arg0).To(Equal("name"))
		Expect(simple.ReportFuncCallParams()[i].Arg1).To(Equal(float64(i*10 + 1)))
	}
}

func (s *SimpleReporterSuite) TestShutdown(t sweet.T) {
	var (
		simple   = NewMockSimpleReporter()
		reporter = NewSimpleReporterShim(simple)
	)

	reporter.Shutdown()
	Expect(simple.ShutdownFuncCallCount()).To(Equal(1))
}

func (s *SimpleReporterSuite) TestValueAdd(t sweet.T) {
	v := &value{}
	Expect(math.Float64frombits(v.val)).To(Equal(float64(0)))
	v.Add(3.5)
	Expect(math.Float64frombits(v.val)).To(Equal(float64(3.5)))
	v.Add(3.5)
	Expect(math.Float64frombits(v.val)).To(Equal(float64(7.0)))
	v.Add(-1.25)
	Expect(math.Float64frombits(v.val)).To(Equal(float64(5.75)))
}

func (s *SimpleReporterSuite) TestValueSet(t sweet.T) {
	v := &value{}
	Expect(math.Float64frombits(v.val)).To(Equal(float64(0)))
	v.Set(4.25)
	Expect(math.Float64frombits(v.val)).To(Equal(float64(4.25)))
	v.Set(12.5)
	Expect(math.Float64frombits(v.val)).To(Equal(float64(12.5)))
}
