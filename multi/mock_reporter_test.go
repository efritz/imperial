// DO NOT EDIT
// Code generated automatically by github.com/efritz/go-mockgen
// $ go-mockgen github.com/efritz/imperial/base -i Reporter -o mock_reporter_test.go -f

package multi

import (
	base "github.com/efritz/imperial/base"
	"sync"
)

type MockReporter struct {
	AddCounterFunc        func(string, float64, ...base.ConfigFunc)
	histAddCounter        []ReporterAddCounterParamSet
	AddGaugeFunc          func(string, float64, ...base.ConfigFunc)
	histAddGauge          []ReporterAddGaugeParamSet
	ObserveHistogramFunc  func(string, float64, ...base.ConfigFunc)
	histObserveHistogram  []ReporterObserveHistogramParamSet
	ObserveSummaryFunc    func(string, float64, ...base.ConfigFunc)
	histObserveSummary    []ReporterObserveSummaryParamSet
	RegisterCounterFunc   func(string, ...base.ConfigFunc)
	histRegisterCounter   []ReporterRegisterCounterParamSet
	RegisterGaugeFunc     func(string, ...base.ConfigFunc)
	histRegisterGauge     []ReporterRegisterGaugeParamSet
	RegisterHistogramFunc func(string, ...base.ConfigFunc)
	histRegisterHistogram []ReporterRegisterHistogramParamSet
	RegisterSummaryFunc   func(string, ...base.ConfigFunc)
	histRegisterSummary   []ReporterRegisterSummaryParamSet
	SetGaugeFunc          func(string, float64, ...base.ConfigFunc)
	histSetGauge          []ReporterSetGaugeParamSet
	ShutdownFunc          func()
	histShutdown          []ReporterShutdownParamSet
	mutex                 sync.RWMutex
}
type ReporterAddCounterParamSet struct {
	Arg0 string
	Arg1 float64
	Arg2 []base.ConfigFunc
}
type ReporterAddGaugeParamSet struct {
	Arg0 string
	Arg1 float64
	Arg2 []base.ConfigFunc
}
type ReporterObserveHistogramParamSet struct {
	Arg0 string
	Arg1 float64
	Arg2 []base.ConfigFunc
}
type ReporterObserveSummaryParamSet struct {
	Arg0 string
	Arg1 float64
	Arg2 []base.ConfigFunc
}
type ReporterRegisterCounterParamSet struct {
	Arg0 string
	Arg1 []base.ConfigFunc
}
type ReporterRegisterGaugeParamSet struct {
	Arg0 string
	Arg1 []base.ConfigFunc
}
type ReporterRegisterHistogramParamSet struct {
	Arg0 string
	Arg1 []base.ConfigFunc
}
type ReporterRegisterSummaryParamSet struct {
	Arg0 string
	Arg1 []base.ConfigFunc
}
type ReporterSetGaugeParamSet struct {
	Arg0 string
	Arg1 float64
	Arg2 []base.ConfigFunc
}
type ReporterShutdownParamSet struct{}

func NewMockReporter() *MockReporter {
	m := &MockReporter{}
	m.AddCounterFunc = m.defaultAddCounterFunc
	m.AddGaugeFunc = m.defaultAddGaugeFunc
	m.ObserveHistogramFunc = m.defaultObserveHistogramFunc
	m.ObserveSummaryFunc = m.defaultObserveSummaryFunc
	m.RegisterCounterFunc = m.defaultRegisterCounterFunc
	m.RegisterGaugeFunc = m.defaultRegisterGaugeFunc
	m.RegisterHistogramFunc = m.defaultRegisterHistogramFunc
	m.RegisterSummaryFunc = m.defaultRegisterSummaryFunc
	m.SetGaugeFunc = m.defaultSetGaugeFunc
	m.ShutdownFunc = m.defaultShutdownFunc
	return m
}
func (m *MockReporter) AddCounter(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histAddCounter = append(m.histAddCounter, ReporterAddCounterParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.AddCounterFunc(v0, v1, v2...)
}
func (m *MockReporter) AddCounterFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histAddCounter)
}
func (m *MockReporter) AddCounterFuncCallParams() []ReporterAddCounterParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histAddCounter
}

func (m *MockReporter) AddGauge(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histAddGauge = append(m.histAddGauge, ReporterAddGaugeParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.AddGaugeFunc(v0, v1, v2...)
}
func (m *MockReporter) AddGaugeFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histAddGauge)
}
func (m *MockReporter) AddGaugeFuncCallParams() []ReporterAddGaugeParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histAddGauge
}

func (m *MockReporter) ObserveHistogram(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histObserveHistogram = append(m.histObserveHistogram, ReporterObserveHistogramParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.ObserveHistogramFunc(v0, v1, v2...)
}
func (m *MockReporter) ObserveHistogramFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histObserveHistogram)
}
func (m *MockReporter) ObserveHistogramFuncCallParams() []ReporterObserveHistogramParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histObserveHistogram
}

func (m *MockReporter) ObserveSummary(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histObserveSummary = append(m.histObserveSummary, ReporterObserveSummaryParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.ObserveSummaryFunc(v0, v1, v2...)
}
func (m *MockReporter) ObserveSummaryFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histObserveSummary)
}
func (m *MockReporter) ObserveSummaryFuncCallParams() []ReporterObserveSummaryParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histObserveSummary
}

func (m *MockReporter) RegisterCounter(v0 string, v1 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histRegisterCounter = append(m.histRegisterCounter, ReporterRegisterCounterParamSet{v0, v1})
	m.mutex.Unlock()
	m.RegisterCounterFunc(v0, v1...)
}
func (m *MockReporter) RegisterCounterFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histRegisterCounter)
}
func (m *MockReporter) RegisterCounterFuncCallParams() []ReporterRegisterCounterParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histRegisterCounter
}

func (m *MockReporter) RegisterGauge(v0 string, v1 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histRegisterGauge = append(m.histRegisterGauge, ReporterRegisterGaugeParamSet{v0, v1})
	m.mutex.Unlock()
	m.RegisterGaugeFunc(v0, v1...)
}
func (m *MockReporter) RegisterGaugeFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histRegisterGauge)
}
func (m *MockReporter) RegisterGaugeFuncCallParams() []ReporterRegisterGaugeParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histRegisterGauge
}

func (m *MockReporter) RegisterHistogram(v0 string, v1 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histRegisterHistogram = append(m.histRegisterHistogram, ReporterRegisterHistogramParamSet{v0, v1})
	m.mutex.Unlock()
	m.RegisterHistogramFunc(v0, v1...)
}
func (m *MockReporter) RegisterHistogramFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histRegisterHistogram)
}
func (m *MockReporter) RegisterHistogramFuncCallParams() []ReporterRegisterHistogramParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histRegisterHistogram
}

func (m *MockReporter) RegisterSummary(v0 string, v1 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histRegisterSummary = append(m.histRegisterSummary, ReporterRegisterSummaryParamSet{v0, v1})
	m.mutex.Unlock()
	m.RegisterSummaryFunc(v0, v1...)
}
func (m *MockReporter) RegisterSummaryFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histRegisterSummary)
}
func (m *MockReporter) RegisterSummaryFuncCallParams() []ReporterRegisterSummaryParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histRegisterSummary
}

func (m *MockReporter) SetGauge(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	m.mutex.Lock()
	m.histSetGauge = append(m.histSetGauge, ReporterSetGaugeParamSet{v0, v1, v2})
	m.mutex.Unlock()
	m.SetGaugeFunc(v0, v1, v2...)
}
func (m *MockReporter) SetGaugeFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histSetGauge)
}
func (m *MockReporter) SetGaugeFuncCallParams() []ReporterSetGaugeParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histSetGauge
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

func (m *MockReporter) defaultAddCounterFunc(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultAddGaugeFunc(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultObserveHistogramFunc(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultObserveSummaryFunc(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultRegisterCounterFunc(v0 string, v1 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultRegisterGaugeFunc(v0 string, v1 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultRegisterHistogramFunc(v0 string, v1 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultRegisterSummaryFunc(v0 string, v1 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultSetGaugeFunc(v0 string, v1 float64, v2 ...base.ConfigFunc) {
	return
}
func (m *MockReporter) defaultShutdownFunc() {
	return
}
