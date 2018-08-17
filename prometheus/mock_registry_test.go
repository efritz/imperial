// DO NOT EDIT
// Code generated automatically by github.com/efritz/go-mockgen
// $ go-mockgen github.com/efritz/imperial/prometheus -i Registry -o mock_registry_test.go -f

package prometheus

import (
	"sync"

	prometheus "github.com/prometheus/client_golang/prometheus"
	go1 "github.com/prometheus/client_model/go"
)

type MockRegistry struct {
	GatherFunc       func() ([]*go1.MetricFamily, error)
	histGather       []RegistryGatherParamSet
	MustRegisterFunc func(...prometheus.Collector)
	histMustRegister []RegistryMustRegisterParamSet
	RegisterFunc     func(prometheus.Collector) error
	histRegister     []RegistryRegisterParamSet
	UnregisterFunc   func(prometheus.Collector) bool
	histUnregister   []RegistryUnregisterParamSet
	mutex            sync.RWMutex
}
type RegistryGatherParamSet struct{}
type RegistryMustRegisterParamSet struct {
	Arg0 []prometheus.Collector
}
type RegistryRegisterParamSet struct {
	Arg0 prometheus.Collector
}
type RegistryUnregisterParamSet struct {
	Arg0 prometheus.Collector
}

func NewMockRegistry() *MockRegistry {
	m := &MockRegistry{}
	m.GatherFunc = m.defaultGatherFunc
	m.MustRegisterFunc = m.defaultMustRegisterFunc
	m.RegisterFunc = m.defaultRegisterFunc
	m.UnregisterFunc = m.defaultUnregisterFunc
	return m
}
func (m *MockRegistry) Gather() ([]*go1.MetricFamily, error) {
	m.mutex.Lock()
	m.histGather = append(m.histGather, RegistryGatherParamSet{})
	m.mutex.Unlock()
	return m.GatherFunc()
}
func (m *MockRegistry) GatherFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGather)
}
func (m *MockRegistry) GatherFuncCallParams() []RegistryGatherParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGather
}

func (m *MockRegistry) MustRegister(v0 ...prometheus.Collector) {
	m.mutex.Lock()
	m.histMustRegister = append(m.histMustRegister, RegistryMustRegisterParamSet{v0})
	m.mutex.Unlock()
	m.MustRegisterFunc(v0...)
}
func (m *MockRegistry) MustRegisterFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histMustRegister)
}
func (m *MockRegistry) MustRegisterFuncCallParams() []RegistryMustRegisterParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histMustRegister
}

func (m *MockRegistry) Register(v0 prometheus.Collector) error {
	m.mutex.Lock()
	m.histRegister = append(m.histRegister, RegistryRegisterParamSet{v0})
	m.mutex.Unlock()
	return m.RegisterFunc(v0)
}
func (m *MockRegistry) RegisterFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histRegister)
}
func (m *MockRegistry) RegisterFuncCallParams() []RegistryRegisterParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histRegister
}

func (m *MockRegistry) Unregister(v0 prometheus.Collector) bool {
	m.mutex.Lock()
	m.histUnregister = append(m.histUnregister, RegistryUnregisterParamSet{v0})
	m.mutex.Unlock()
	return m.UnregisterFunc(v0)
}
func (m *MockRegistry) UnregisterFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histUnregister)
}
func (m *MockRegistry) UnregisterFuncCallParams() []RegistryUnregisterParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histUnregister
}

func (m *MockRegistry) defaultGatherFunc() ([]*go1.MetricFamily, error) {
	return nil, nil
}
func (m *MockRegistry) defaultMustRegisterFunc(v0 ...prometheus.Collector) {
	return
}
func (m *MockRegistry) defaultRegisterFunc(v0 prometheus.Collector) error {
	return nil
}
func (m *MockRegistry) defaultUnregisterFunc(v0 prometheus.Collector) bool {
	return false
}
