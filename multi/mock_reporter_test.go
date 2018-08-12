// DO NOT EDIT
// Code generated automatically by github.com/efritz/go-mockgen
// $ go-mockgen github.com/efritz/imperial/base -i Reporter -o mock_reporter_test.go -f

package multi

import (
	base "github.com/efritz/imperial/base"
	"sync"
)

type MockReporter struct {
	ReportFunc   func(string, int, ...base.ConfigFunc)
	histReport   []ReporterReportParamSet
	ShutdownFunc func()
	histShutdown []ReporterShutdownParamSet
	mutex        sync.RWMutex
}
type ReporterReportParamSet struct {
	Arg0 string
	Arg1 int
	Arg2 []base.ConfigFunc
}
type ReporterShutdownParamSet struct{}

func NewMockReporter() *MockReporter {
	m := &MockReporter{}
	m.ReportFunc = m.defaultReportFunc
	m.ShutdownFunc = m.defaultShutdownFunc
	return m
}
func (m *MockReporter) Report(v0 string, v1 int, v2 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histReport = append(m.histReport, ReporterReportParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.ReportFunc(v0, v1, v2...)
}
func (m *MockReporter) ReportFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histReport)
}
func (m *MockReporter) ReportFuncCallParams() []ReporterReportParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histReport
}

func (m *MockReporter) Shutdown() {
	m.mutex.Lock()
	m.histShutdown = append(m.histShutdown, ReporterShutdownParamSet{})
	m.mutex.Unlock()
	m.ShutdownFunc()
}
func (m *MockReporter) ShutdownFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histShutdown)
}
func (m *MockReporter) ShutdownFuncCallParams() []ReporterShutdownParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histShutdown
}

func (m *MockReporter) defaultReportFunc(v0 string, v1 int, v2 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultShutdownFunc() {
	return
}
