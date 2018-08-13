package multi

//go:generate go-mockgen github.com/efritz/imperial/base -i Reporter -o mock_reporter_test.go -f

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type MultiSuite struct{}

func (s *MultiSuite) TestAddCounterDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewReporter(r1, r2, r3).AddCounter("requests", 35)

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.AddCounterFuncCallCount()).To(Equal(1))
		Expect(r.AddCounterFuncCallParams()[0].Arg0).To(Equal("requests"))
		Expect(r.AddCounterFuncCallParams()[0].Arg1).To(Equal(float64(35)))
	}
}

func (s *MultiSuite) TestAddGaugeDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewReporter(r1, r2, r3).AddGauge("requests", 35)

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.AddGaugeFuncCallCount()).To(Equal(1))
		Expect(r.AddGaugeFuncCallParams()[0].Arg0).To(Equal("requests"))
		Expect(r.AddGaugeFuncCallParams()[0].Arg1).To(Equal(float64(35)))
	}
}

func (s *MultiSuite) TestSetGaugeDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewReporter(r1, r2, r3).SetGauge("requests", 35)

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.SetGaugeFuncCallCount()).To(Equal(1))
		Expect(r.SetGaugeFuncCallParams()[0].Arg0).To(Equal("requests"))
		Expect(r.SetGaugeFuncCallParams()[0].Arg1).To(Equal(float64(35)))
	}
}

func (s *MultiSuite) TestObserveHistogramDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewReporter(r1, r2, r3).ObserveHistogram("requests", 35)

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.ObserveHistogramFuncCallCount()).To(Equal(1))
		Expect(r.ObserveHistogramFuncCallParams()[0].Arg0).To(Equal("requests"))
		Expect(r.ObserveHistogramFuncCallParams()[0].Arg1).To(Equal(float64(35)))
	}
}

func (s *MultiSuite) TestObserveSummaryDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewReporter(r1, r2, r3).ObserveSummary("requests", 35)

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.ObserveSummaryFuncCallCount()).To(Equal(1))
		Expect(r.ObserveSummaryFuncCallParams()[0].Arg0).To(Equal("requests"))
		Expect(r.ObserveSummaryFuncCallParams()[0].Arg1).To(Equal(float64(35)))
	}
}

func (s *MultiSuite) TestShutdownDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewReporter(r1, r2, r3).Shutdown()

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.ShutdownFuncCallCount()).To(Equal(1))
	}
}
