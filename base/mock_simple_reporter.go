// DO NOT EDIT
// Code generated automatically by github.com/efritz/go-mockgen
// $ go-mockgen github.com/efritz/imperial/base -i SimpleReporter -o mock_simple_reporter.go -f

package base

import "sync"

type MockSimpleReporter struct {
	ReportFunc   func(string, float64, ...ConfigFunc)
	histReport   []SimpleReporterReportParamSet
	ShutdownFunc func()
	histShutdown []SimpleReporterShutdownParamSet
	mutex        sync.RWMutex
}
type SimpleReporterReportParamSet struct {
	Arg0 string
	Arg1 float64
	Arg2 []ConfigFunc
}
type SimpleReporterShutdownParamSet struct{}

func NewMockSimpleReporter() *MockSimpleReporter {
	m := &MockSimpleReporter{}
	m.ReportFunc = m.defaultReportFunc
	m.ShutdownFunc = m.defaultShutdownFunc
	return m
}
func (m *MockSimpleReporter) Report(v0 string, v1 float64, v2 ...ConfigFunc) {
	m.mutex.Lock()
	m.histReport = append(m.histReport, SimpleReporterReportParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.ReportFunc(v0, v1, v2...)
}
func (m *MockSimpleReporter) ReportFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histReport)
}
func (m *MockSimpleReporter) ReportFuncCallParams() []SimpleReporterReportParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histReport
}

func (m *MockSimpleReporter) Shutdown() {
	m.mutex.Lock()
	m.histShutdown = append(m.histShutdown, SimpleReporterShutdownParamSet{})
	m.mutex.Unlock()
	m.ShutdownFunc()
}
func (m *MockSimpleReporter) ShutdownFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histShutdown)
}
func (m *MockSimpleReporter) ShutdownFuncCallParams() []SimpleReporterShutdownParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histShutdown
}

func (m *MockSimpleReporter) defaultReportFunc(v0 string, v1 float64, v2 ...ConfigFunc) {
	return
}
func (m *MockSimpleReporter) defaultShutdownFunc() {
	return
}
