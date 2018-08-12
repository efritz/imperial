package imperial

import (
	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type MultiReporterSuite struct{}

func (s *MultiReporterSuite) TestReportDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewMultiReporter(r1, r2, r3).Report("requests", 35)

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.ReportFuncCallCount()).To(Equal(1))
		Expect(r.ReportFuncCallParams()[0].Arg0).To(Equal("requests"))
		Expect(r.ReportFuncCallParams()[0].Arg1).To(Equal(35))
	}
}

func (s *MultiReporterSuite) TestShutdownDelegates(t sweet.T) {
	var (
		r1 = NewMockReporter()
		r2 = NewMockReporter()
		r3 = NewMockReporter()
	)

	NewMultiReporter(r1, r2, r3).Shutdown()

	for _, r := range []*MockReporter{r1, r2, r3} {
		Expect(r.ShutdownFuncCallCount()).To(Equal(1))
	}
}
