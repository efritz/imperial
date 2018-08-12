// DO NOT EDIT
// Code generated automatically by github.com/efritz/go-mockgen
// $ go-mockgen github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface -i CloudWatchAPI -o mock_cloudwatch_api_test.go -f

package cloudwatch

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	cloudwatch "github.com/aws/aws-sdk-go/service/cloudwatch"
	"sync"
)

type MockCloudWatchAPI struct {
	DeleteAlarmsFunc                         func(*cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error)
	histDeleteAlarms                         []CloudWatchAPIDeleteAlarmsParamSet
	DeleteAlarmsRequestFunc                  func(*cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput)
	histDeleteAlarmsRequest                  []CloudWatchAPIDeleteAlarmsRequestParamSet
	DeleteAlarmsWithContextFunc              func(aws.Context, *cloudwatch.DeleteAlarmsInput, ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error)
	histDeleteAlarmsWithContext              []CloudWatchAPIDeleteAlarmsWithContextParamSet
	DeleteDashboardsFunc                     func(*cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error)
	histDeleteDashboards                     []CloudWatchAPIDeleteDashboardsParamSet
	DeleteDashboardsRequestFunc              func(*cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput)
	histDeleteDashboardsRequest              []CloudWatchAPIDeleteDashboardsRequestParamSet
	DeleteDashboardsWithContextFunc          func(aws.Context, *cloudwatch.DeleteDashboardsInput, ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error)
	histDeleteDashboardsWithContext          []CloudWatchAPIDeleteDashboardsWithContextParamSet
	DescribeAlarmHistoryFunc                 func(*cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	histDescribeAlarmHistory                 []CloudWatchAPIDescribeAlarmHistoryParamSet
	DescribeAlarmHistoryPagesFunc            func(*cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error
	histDescribeAlarmHistoryPages            []CloudWatchAPIDescribeAlarmHistoryPagesParamSet
	DescribeAlarmHistoryPagesWithContextFunc func(aws.Context, *cloudwatch.DescribeAlarmHistoryInput, func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, ...request.Option) error
	histDescribeAlarmHistoryPagesWithContext []CloudWatchAPIDescribeAlarmHistoryPagesWithContextParamSet
	DescribeAlarmHistoryRequestFunc          func(*cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput)
	histDescribeAlarmHistoryRequest          []CloudWatchAPIDescribeAlarmHistoryRequestParamSet
	DescribeAlarmHistoryWithContextFunc      func(aws.Context, *cloudwatch.DescribeAlarmHistoryInput, ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error)
	histDescribeAlarmHistoryWithContext      []CloudWatchAPIDescribeAlarmHistoryWithContextParamSet
	DescribeAlarmsFunc                       func(*cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error)
	histDescribeAlarms                       []CloudWatchAPIDescribeAlarmsParamSet
	DescribeAlarmsForMetricFunc              func(*cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	histDescribeAlarmsForMetric              []CloudWatchAPIDescribeAlarmsForMetricParamSet
	DescribeAlarmsForMetricRequestFunc       func(*cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput)
	histDescribeAlarmsForMetricRequest       []CloudWatchAPIDescribeAlarmsForMetricRequestParamSet
	DescribeAlarmsForMetricWithContextFunc   func(aws.Context, *cloudwatch.DescribeAlarmsForMetricInput, ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error)
	histDescribeAlarmsForMetricWithContext   []CloudWatchAPIDescribeAlarmsForMetricWithContextParamSet
	DescribeAlarmsPagesFunc                  func(*cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error
	histDescribeAlarmsPages                  []CloudWatchAPIDescribeAlarmsPagesParamSet
	DescribeAlarmsPagesWithContextFunc       func(aws.Context, *cloudwatch.DescribeAlarmsInput, func(*cloudwatch.DescribeAlarmsOutput, bool) bool, ...request.Option) error
	histDescribeAlarmsPagesWithContext       []CloudWatchAPIDescribeAlarmsPagesWithContextParamSet
	DescribeAlarmsRequestFunc                func(*cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput)
	histDescribeAlarmsRequest                []CloudWatchAPIDescribeAlarmsRequestParamSet
	DescribeAlarmsWithContextFunc            func(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error)
	histDescribeAlarmsWithContext            []CloudWatchAPIDescribeAlarmsWithContextParamSet
	DisableAlarmActionsFunc                  func(*cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error)
	histDisableAlarmActions                  []CloudWatchAPIDisableAlarmActionsParamSet
	DisableAlarmActionsRequestFunc           func(*cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput)
	histDisableAlarmActionsRequest           []CloudWatchAPIDisableAlarmActionsRequestParamSet
	DisableAlarmActionsWithContextFunc       func(aws.Context, *cloudwatch.DisableAlarmActionsInput, ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error)
	histDisableAlarmActionsWithContext       []CloudWatchAPIDisableAlarmActionsWithContextParamSet
	EnableAlarmActionsFunc                   func(*cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error)
	histEnableAlarmActions                   []CloudWatchAPIEnableAlarmActionsParamSet
	EnableAlarmActionsRequestFunc            func(*cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput)
	histEnableAlarmActionsRequest            []CloudWatchAPIEnableAlarmActionsRequestParamSet
	EnableAlarmActionsWithContextFunc        func(aws.Context, *cloudwatch.EnableAlarmActionsInput, ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error)
	histEnableAlarmActionsWithContext        []CloudWatchAPIEnableAlarmActionsWithContextParamSet
	GetDashboardFunc                         func(*cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error)
	histGetDashboard                         []CloudWatchAPIGetDashboardParamSet
	GetDashboardRequestFunc                  func(*cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput)
	histGetDashboardRequest                  []CloudWatchAPIGetDashboardRequestParamSet
	GetDashboardWithContextFunc              func(aws.Context, *cloudwatch.GetDashboardInput, ...request.Option) (*cloudwatch.GetDashboardOutput, error)
	histGetDashboardWithContext              []CloudWatchAPIGetDashboardWithContextParamSet
	GetMetricDataFunc                        func(*cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error)
	histGetMetricData                        []CloudWatchAPIGetMetricDataParamSet
	GetMetricDataRequestFunc                 func(*cloudwatch.GetMetricDataInput) (*request.Request, *cloudwatch.GetMetricDataOutput)
	histGetMetricDataRequest                 []CloudWatchAPIGetMetricDataRequestParamSet
	GetMetricDataWithContextFunc             func(aws.Context, *cloudwatch.GetMetricDataInput, ...request.Option) (*cloudwatch.GetMetricDataOutput, error)
	histGetMetricDataWithContext             []CloudWatchAPIGetMetricDataWithContextParamSet
	GetMetricStatisticsFunc                  func(*cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error)
	histGetMetricStatistics                  []CloudWatchAPIGetMetricStatisticsParamSet
	GetMetricStatisticsRequestFunc           func(*cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput)
	histGetMetricStatisticsRequest           []CloudWatchAPIGetMetricStatisticsRequestParamSet
	GetMetricStatisticsWithContextFunc       func(aws.Context, *cloudwatch.GetMetricStatisticsInput, ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error)
	histGetMetricStatisticsWithContext       []CloudWatchAPIGetMetricStatisticsWithContextParamSet
	ListDashboardsFunc                       func(*cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error)
	histListDashboards                       []CloudWatchAPIListDashboardsParamSet
	ListDashboardsRequestFunc                func(*cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput)
	histListDashboardsRequest                []CloudWatchAPIListDashboardsRequestParamSet
	ListDashboardsWithContextFunc            func(aws.Context, *cloudwatch.ListDashboardsInput, ...request.Option) (*cloudwatch.ListDashboardsOutput, error)
	histListDashboardsWithContext            []CloudWatchAPIListDashboardsWithContextParamSet
	ListMetricsFunc                          func(*cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error)
	histListMetrics                          []CloudWatchAPIListMetricsParamSet
	ListMetricsPagesFunc                     func(*cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool) error
	histListMetricsPages                     []CloudWatchAPIListMetricsPagesParamSet
	ListMetricsPagesWithContextFunc          func(aws.Context, *cloudwatch.ListMetricsInput, func(*cloudwatch.ListMetricsOutput, bool) bool, ...request.Option) error
	histListMetricsPagesWithContext          []CloudWatchAPIListMetricsPagesWithContextParamSet
	ListMetricsRequestFunc                   func(*cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput)
	histListMetricsRequest                   []CloudWatchAPIListMetricsRequestParamSet
	ListMetricsWithContextFunc               func(aws.Context, *cloudwatch.ListMetricsInput, ...request.Option) (*cloudwatch.ListMetricsOutput, error)
	histListMetricsWithContext               []CloudWatchAPIListMetricsWithContextParamSet
	PutDashboardFunc                         func(*cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error)
	histPutDashboard                         []CloudWatchAPIPutDashboardParamSet
	PutDashboardRequestFunc                  func(*cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput)
	histPutDashboardRequest                  []CloudWatchAPIPutDashboardRequestParamSet
	PutDashboardWithContextFunc              func(aws.Context, *cloudwatch.PutDashboardInput, ...request.Option) (*cloudwatch.PutDashboardOutput, error)
	histPutDashboardWithContext              []CloudWatchAPIPutDashboardWithContextParamSet
	PutMetricAlarmFunc                       func(*cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error)
	histPutMetricAlarm                       []CloudWatchAPIPutMetricAlarmParamSet
	PutMetricAlarmRequestFunc                func(*cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput)
	histPutMetricAlarmRequest                []CloudWatchAPIPutMetricAlarmRequestParamSet
	PutMetricAlarmWithContextFunc            func(aws.Context, *cloudwatch.PutMetricAlarmInput, ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error)
	histPutMetricAlarmWithContext            []CloudWatchAPIPutMetricAlarmWithContextParamSet
	PutMetricDataFunc                        func(*cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error)
	histPutMetricData                        []CloudWatchAPIPutMetricDataParamSet
	PutMetricDataRequestFunc                 func(*cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput)
	histPutMetricDataRequest                 []CloudWatchAPIPutMetricDataRequestParamSet
	PutMetricDataWithContextFunc             func(aws.Context, *cloudwatch.PutMetricDataInput, ...request.Option) (*cloudwatch.PutMetricDataOutput, error)
	histPutMetricDataWithContext             []CloudWatchAPIPutMetricDataWithContextParamSet
	SetAlarmStateFunc                        func(*cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error)
	histSetAlarmState                        []CloudWatchAPISetAlarmStateParamSet
	SetAlarmStateRequestFunc                 func(*cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput)
	histSetAlarmStateRequest                 []CloudWatchAPISetAlarmStateRequestParamSet
	SetAlarmStateWithContextFunc             func(aws.Context, *cloudwatch.SetAlarmStateInput, ...request.Option) (*cloudwatch.SetAlarmStateOutput, error)
	histSetAlarmStateWithContext             []CloudWatchAPISetAlarmStateWithContextParamSet
	WaitUntilAlarmExistsFunc                 func(*cloudwatch.DescribeAlarmsInput) error
	histWaitUntilAlarmExists                 []CloudWatchAPIWaitUntilAlarmExistsParamSet
	WaitUntilAlarmExistsWithContextFunc      func(aws.Context, *cloudwatch.DescribeAlarmsInput, ...request.WaiterOption) error
	histWaitUntilAlarmExistsWithContext      []CloudWatchAPIWaitUntilAlarmExistsWithContextParamSet
	mutex                                    sync.RWMutex
}
type CloudWatchAPIDeleteAlarmsParamSet struct {
	Arg0 *cloudwatch.DeleteAlarmsInput
}
type CloudWatchAPIDeleteAlarmsRequestParamSet struct {
	Arg0 *cloudwatch.DeleteAlarmsInput
}
type CloudWatchAPIDeleteAlarmsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DeleteAlarmsInput
	Arg2 []request.Option
}
type CloudWatchAPIDeleteDashboardsParamSet struct {
	Arg0 *cloudwatch.DeleteDashboardsInput
}
type CloudWatchAPIDeleteDashboardsRequestParamSet struct {
	Arg0 *cloudwatch.DeleteDashboardsInput
}
type CloudWatchAPIDeleteDashboardsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DeleteDashboardsInput
	Arg2 []request.Option
}
type CloudWatchAPIDescribeAlarmHistoryParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmHistoryInput
}
type CloudWatchAPIDescribeAlarmHistoryPagesParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmHistoryInput
	Arg1 func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool
}
type CloudWatchAPIDescribeAlarmHistoryPagesWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DescribeAlarmHistoryInput
	Arg2 func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool
	Arg3 []request.Option
}
type CloudWatchAPIDescribeAlarmHistoryRequestParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmHistoryInput
}
type CloudWatchAPIDescribeAlarmHistoryWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DescribeAlarmHistoryInput
	Arg2 []request.Option
}
type CloudWatchAPIDescribeAlarmsParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmsInput
}
type CloudWatchAPIDescribeAlarmsForMetricParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmsForMetricInput
}
type CloudWatchAPIDescribeAlarmsForMetricRequestParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmsForMetricInput
}
type CloudWatchAPIDescribeAlarmsForMetricWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DescribeAlarmsForMetricInput
	Arg2 []request.Option
}
type CloudWatchAPIDescribeAlarmsPagesParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmsInput
	Arg1 func(*cloudwatch.DescribeAlarmsOutput, bool) bool
}
type CloudWatchAPIDescribeAlarmsPagesWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DescribeAlarmsInput
	Arg2 func(*cloudwatch.DescribeAlarmsOutput, bool) bool
	Arg3 []request.Option
}
type CloudWatchAPIDescribeAlarmsRequestParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmsInput
}
type CloudWatchAPIDescribeAlarmsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DescribeAlarmsInput
	Arg2 []request.Option
}
type CloudWatchAPIDisableAlarmActionsParamSet struct {
	Arg0 *cloudwatch.DisableAlarmActionsInput
}
type CloudWatchAPIDisableAlarmActionsRequestParamSet struct {
	Arg0 *cloudwatch.DisableAlarmActionsInput
}
type CloudWatchAPIDisableAlarmActionsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DisableAlarmActionsInput
	Arg2 []request.Option
}
type CloudWatchAPIEnableAlarmActionsParamSet struct {
	Arg0 *cloudwatch.EnableAlarmActionsInput
}
type CloudWatchAPIEnableAlarmActionsRequestParamSet struct {
	Arg0 *cloudwatch.EnableAlarmActionsInput
}
type CloudWatchAPIEnableAlarmActionsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.EnableAlarmActionsInput
	Arg2 []request.Option
}
type CloudWatchAPIGetDashboardParamSet struct {
	Arg0 *cloudwatch.GetDashboardInput
}
type CloudWatchAPIGetDashboardRequestParamSet struct {
	Arg0 *cloudwatch.GetDashboardInput
}
type CloudWatchAPIGetDashboardWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.GetDashboardInput
	Arg2 []request.Option
}
type CloudWatchAPIGetMetricDataParamSet struct {
	Arg0 *cloudwatch.GetMetricDataInput
}
type CloudWatchAPIGetMetricDataRequestParamSet struct {
	Arg0 *cloudwatch.GetMetricDataInput
}
type CloudWatchAPIGetMetricDataWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.GetMetricDataInput
	Arg2 []request.Option
}
type CloudWatchAPIGetMetricStatisticsParamSet struct {
	Arg0 *cloudwatch.GetMetricStatisticsInput
}
type CloudWatchAPIGetMetricStatisticsRequestParamSet struct {
	Arg0 *cloudwatch.GetMetricStatisticsInput
}
type CloudWatchAPIGetMetricStatisticsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.GetMetricStatisticsInput
	Arg2 []request.Option
}
type CloudWatchAPIListDashboardsParamSet struct {
	Arg0 *cloudwatch.ListDashboardsInput
}
type CloudWatchAPIListDashboardsRequestParamSet struct {
	Arg0 *cloudwatch.ListDashboardsInput
}
type CloudWatchAPIListDashboardsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.ListDashboardsInput
	Arg2 []request.Option
}
type CloudWatchAPIListMetricsParamSet struct {
	Arg0 *cloudwatch.ListMetricsInput
}
type CloudWatchAPIListMetricsPagesParamSet struct {
	Arg0 *cloudwatch.ListMetricsInput
	Arg1 func(*cloudwatch.ListMetricsOutput, bool) bool
}
type CloudWatchAPIListMetricsPagesWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.ListMetricsInput
	Arg2 func(*cloudwatch.ListMetricsOutput, bool) bool
	Arg3 []request.Option
}
type CloudWatchAPIListMetricsRequestParamSet struct {
	Arg0 *cloudwatch.ListMetricsInput
}
type CloudWatchAPIListMetricsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.ListMetricsInput
	Arg2 []request.Option
}
type CloudWatchAPIPutDashboardParamSet struct {
	Arg0 *cloudwatch.PutDashboardInput
}
type CloudWatchAPIPutDashboardRequestParamSet struct {
	Arg0 *cloudwatch.PutDashboardInput
}
type CloudWatchAPIPutDashboardWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.PutDashboardInput
	Arg2 []request.Option
}
type CloudWatchAPIPutMetricAlarmParamSet struct {
	Arg0 *cloudwatch.PutMetricAlarmInput
}
type CloudWatchAPIPutMetricAlarmRequestParamSet struct {
	Arg0 *cloudwatch.PutMetricAlarmInput
}
type CloudWatchAPIPutMetricAlarmWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.PutMetricAlarmInput
	Arg2 []request.Option
}
type CloudWatchAPIPutMetricDataParamSet struct {
	Arg0 *cloudwatch.PutMetricDataInput
}
type CloudWatchAPIPutMetricDataRequestParamSet struct {
	Arg0 *cloudwatch.PutMetricDataInput
}
type CloudWatchAPIPutMetricDataWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.PutMetricDataInput
	Arg2 []request.Option
}
type CloudWatchAPISetAlarmStateParamSet struct {
	Arg0 *cloudwatch.SetAlarmStateInput
}
type CloudWatchAPISetAlarmStateRequestParamSet struct {
	Arg0 *cloudwatch.SetAlarmStateInput
}
type CloudWatchAPISetAlarmStateWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.SetAlarmStateInput
	Arg2 []request.Option
}
type CloudWatchAPIWaitUntilAlarmExistsParamSet struct {
	Arg0 *cloudwatch.DescribeAlarmsInput
}
type CloudWatchAPIWaitUntilAlarmExistsWithContextParamSet struct {
	Arg0 aws.Context
	Arg1 *cloudwatch.DescribeAlarmsInput
	Arg2 []request.WaiterOption
}

func NewMockCloudWatchAPI() *MockCloudWatchAPI {
	m := &MockCloudWatchAPI{}
	m.DeleteAlarmsFunc = m.defaultDeleteAlarmsFunc
	m.DeleteAlarmsRequestFunc = m.defaultDeleteAlarmsRequestFunc
	m.DeleteAlarmsWithContextFunc = m.defaultDeleteAlarmsWithContextFunc
	m.DeleteDashboardsFunc = m.defaultDeleteDashboardsFunc
	m.DeleteDashboardsRequestFunc = m.defaultDeleteDashboardsRequestFunc
	m.DeleteDashboardsWithContextFunc = m.defaultDeleteDashboardsWithContextFunc
	m.DescribeAlarmHistoryFunc = m.defaultDescribeAlarmHistoryFunc
	m.DescribeAlarmHistoryPagesFunc = m.defaultDescribeAlarmHistoryPagesFunc
	m.DescribeAlarmHistoryPagesWithContextFunc = m.defaultDescribeAlarmHistoryPagesWithContextFunc
	m.DescribeAlarmHistoryRequestFunc = m.defaultDescribeAlarmHistoryRequestFunc
	m.DescribeAlarmHistoryWithContextFunc = m.defaultDescribeAlarmHistoryWithContextFunc
	m.DescribeAlarmsFunc = m.defaultDescribeAlarmsFunc
	m.DescribeAlarmsForMetricFunc = m.defaultDescribeAlarmsForMetricFunc
	m.DescribeAlarmsForMetricRequestFunc = m.defaultDescribeAlarmsForMetricRequestFunc
	m.DescribeAlarmsForMetricWithContextFunc = m.defaultDescribeAlarmsForMetricWithContextFunc
	m.DescribeAlarmsPagesFunc = m.defaultDescribeAlarmsPagesFunc
	m.DescribeAlarmsPagesWithContextFunc = m.defaultDescribeAlarmsPagesWithContextFunc
	m.DescribeAlarmsRequestFunc = m.defaultDescribeAlarmsRequestFunc
	m.DescribeAlarmsWithContextFunc = m.defaultDescribeAlarmsWithContextFunc
	m.DisableAlarmActionsFunc = m.defaultDisableAlarmActionsFunc
	m.DisableAlarmActionsRequestFunc = m.defaultDisableAlarmActionsRequestFunc
	m.DisableAlarmActionsWithContextFunc = m.defaultDisableAlarmActionsWithContextFunc
	m.EnableAlarmActionsFunc = m.defaultEnableAlarmActionsFunc
	m.EnableAlarmActionsRequestFunc = m.defaultEnableAlarmActionsRequestFunc
	m.EnableAlarmActionsWithContextFunc = m.defaultEnableAlarmActionsWithContextFunc
	m.GetDashboardFunc = m.defaultGetDashboardFunc
	m.GetDashboardRequestFunc = m.defaultGetDashboardRequestFunc
	m.GetDashboardWithContextFunc = m.defaultGetDashboardWithContextFunc
	m.GetMetricDataFunc = m.defaultGetMetricDataFunc
	m.GetMetricDataRequestFunc = m.defaultGetMetricDataRequestFunc
	m.GetMetricDataWithContextFunc = m.defaultGetMetricDataWithContextFunc
	m.GetMetricStatisticsFunc = m.defaultGetMetricStatisticsFunc
	m.GetMetricStatisticsRequestFunc = m.defaultGetMetricStatisticsRequestFunc
	m.GetMetricStatisticsWithContextFunc = m.defaultGetMetricStatisticsWithContextFunc
	m.ListDashboardsFunc = m.defaultListDashboardsFunc
	m.ListDashboardsRequestFunc = m.defaultListDashboardsRequestFunc
	m.ListDashboardsWithContextFunc = m.defaultListDashboardsWithContextFunc
	m.ListMetricsFunc = m.defaultListMetricsFunc
	m.ListMetricsPagesFunc = m.defaultListMetricsPagesFunc
	m.ListMetricsPagesWithContextFunc = m.defaultListMetricsPagesWithContextFunc
	m.ListMetricsRequestFunc = m.defaultListMetricsRequestFunc
	m.ListMetricsWithContextFunc = m.defaultListMetricsWithContextFunc
	m.PutDashboardFunc = m.defaultPutDashboardFunc
	m.PutDashboardRequestFunc = m.defaultPutDashboardRequestFunc
	m.PutDashboardWithContextFunc = m.defaultPutDashboardWithContextFunc
	m.PutMetricAlarmFunc = m.defaultPutMetricAlarmFunc
	m.PutMetricAlarmRequestFunc = m.defaultPutMetricAlarmRequestFunc
	m.PutMetricAlarmWithContextFunc = m.defaultPutMetricAlarmWithContextFunc
	m.PutMetricDataFunc = m.defaultPutMetricDataFunc
	m.PutMetricDataRequestFunc = m.defaultPutMetricDataRequestFunc
	m.PutMetricDataWithContextFunc = m.defaultPutMetricDataWithContextFunc
	m.SetAlarmStateFunc = m.defaultSetAlarmStateFunc
	m.SetAlarmStateRequestFunc = m.defaultSetAlarmStateRequestFunc
	m.SetAlarmStateWithContextFunc = m.defaultSetAlarmStateWithContextFunc
	m.WaitUntilAlarmExistsFunc = m.defaultWaitUntilAlarmExistsFunc
	m.WaitUntilAlarmExistsWithContextFunc = m.defaultWaitUntilAlarmExistsWithContextFunc
	return m
}
func (m *MockCloudWatchAPI) DeleteAlarms(v0 *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	m.mutex.Lock()
	m.histDeleteAlarms = append(m.histDeleteAlarms, CloudWatchAPIDeleteAlarmsParamSet{v0})
	m.mutex.Unlock()
	return m.DeleteAlarmsFunc(v0)
}
func (m *MockCloudWatchAPI) DeleteAlarmsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDeleteAlarms)
}
func (m *MockCloudWatchAPI) DeleteAlarmsFuncCallParams() []CloudWatchAPIDeleteAlarmsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDeleteAlarms
}

func (m *MockCloudWatchAPI) DeleteAlarmsRequest(v0 *cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput) {
	m.mutex.Lock()
	m.histDeleteAlarmsRequest = append(m.histDeleteAlarmsRequest, CloudWatchAPIDeleteAlarmsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.DeleteAlarmsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) DeleteAlarmsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDeleteAlarmsRequest)
}
func (m *MockCloudWatchAPI) DeleteAlarmsRequestFuncCallParams() []CloudWatchAPIDeleteAlarmsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDeleteAlarmsRequest
}

func (m *MockCloudWatchAPI) DeleteAlarmsWithContext(v0 aws.Context, v1 *cloudwatch.DeleteAlarmsInput, v2 ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error) {
	m.mutex.Lock()
	m.histDeleteAlarmsWithContext = append(m.histDeleteAlarmsWithContext, CloudWatchAPIDeleteAlarmsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.DeleteAlarmsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) DeleteAlarmsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDeleteAlarmsWithContext)
}
func (m *MockCloudWatchAPI) DeleteAlarmsWithContextFuncCallParams() []CloudWatchAPIDeleteAlarmsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDeleteAlarmsWithContext
}

func (m *MockCloudWatchAPI) DeleteDashboards(v0 *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	m.mutex.Lock()
	m.histDeleteDashboards = append(m.histDeleteDashboards, CloudWatchAPIDeleteDashboardsParamSet{v0})
	m.mutex.Unlock()
	return m.DeleteDashboardsFunc(v0)
}
func (m *MockCloudWatchAPI) DeleteDashboardsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDeleteDashboards)
}
func (m *MockCloudWatchAPI) DeleteDashboardsFuncCallParams() []CloudWatchAPIDeleteDashboardsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDeleteDashboards
}

func (m *MockCloudWatchAPI) DeleteDashboardsRequest(v0 *cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput) {
	m.mutex.Lock()
	m.histDeleteDashboardsRequest = append(m.histDeleteDashboardsRequest, CloudWatchAPIDeleteDashboardsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.DeleteDashboardsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) DeleteDashboardsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDeleteDashboardsRequest)
}
func (m *MockCloudWatchAPI) DeleteDashboardsRequestFuncCallParams() []CloudWatchAPIDeleteDashboardsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDeleteDashboardsRequest
}

func (m *MockCloudWatchAPI) DeleteDashboardsWithContext(v0 aws.Context, v1 *cloudwatch.DeleteDashboardsInput, v2 ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error) {
	m.mutex.Lock()
	m.histDeleteDashboardsWithContext = append(m.histDeleteDashboardsWithContext, CloudWatchAPIDeleteDashboardsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.DeleteDashboardsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) DeleteDashboardsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDeleteDashboardsWithContext)
}
func (m *MockCloudWatchAPI) DeleteDashboardsWithContextFuncCallParams() []CloudWatchAPIDeleteDashboardsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDeleteDashboardsWithContext
}

func (m *MockCloudWatchAPI) DescribeAlarmHistory(v0 *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	m.mutex.Lock()
	m.histDescribeAlarmHistory = append(m.histDescribeAlarmHistory, CloudWatchAPIDescribeAlarmHistoryParamSet{v0})
	m.mutex.Unlock()
	return m.DescribeAlarmHistoryFunc(v0)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmHistory)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryFuncCallParams() []CloudWatchAPIDescribeAlarmHistoryParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmHistory
}

func (m *MockCloudWatchAPI) DescribeAlarmHistoryPages(v0 *cloudwatch.DescribeAlarmHistoryInput, v1 func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error {
	m.mutex.Lock()
	m.histDescribeAlarmHistoryPages = append(m.histDescribeAlarmHistoryPages, CloudWatchAPIDescribeAlarmHistoryPagesParamSet{v0, v1})
	m.mutex.Unlock()
	return m.DescribeAlarmHistoryPagesFunc(v0, v1)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryPagesFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmHistoryPages)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryPagesFuncCallParams() []CloudWatchAPIDescribeAlarmHistoryPagesParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmHistoryPages
}

func (m *MockCloudWatchAPI) DescribeAlarmHistoryPagesWithContext(v0 aws.Context, v1 *cloudwatch.DescribeAlarmHistoryInput, v2 func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, v3 ...request.Option) error {
	m.mutex.Lock()
	m.histDescribeAlarmHistoryPagesWithContext = append(m.histDescribeAlarmHistoryPagesWithContext, CloudWatchAPIDescribeAlarmHistoryPagesWithContextParamSet{v0, v1, v2, v3})
	m.mutex.Unlock()
	return m.DescribeAlarmHistoryPagesWithContextFunc(v0, v1, v2, v3...)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryPagesWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmHistoryPagesWithContext)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryPagesWithContextFuncCallParams() []CloudWatchAPIDescribeAlarmHistoryPagesWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmHistoryPagesWithContext
}

func (m *MockCloudWatchAPI) DescribeAlarmHistoryRequest(v0 *cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput) {
	m.mutex.Lock()
	m.histDescribeAlarmHistoryRequest = append(m.histDescribeAlarmHistoryRequest, CloudWatchAPIDescribeAlarmHistoryRequestParamSet{v0})
	m.mutex.Unlock()
	return m.DescribeAlarmHistoryRequestFunc(v0)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmHistoryRequest)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryRequestFuncCallParams() []CloudWatchAPIDescribeAlarmHistoryRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmHistoryRequest
}

func (m *MockCloudWatchAPI) DescribeAlarmHistoryWithContext(v0 aws.Context, v1 *cloudwatch.DescribeAlarmHistoryInput, v2 ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	m.mutex.Lock()
	m.histDescribeAlarmHistoryWithContext = append(m.histDescribeAlarmHistoryWithContext, CloudWatchAPIDescribeAlarmHistoryWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.DescribeAlarmHistoryWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmHistoryWithContext)
}
func (m *MockCloudWatchAPI) DescribeAlarmHistoryWithContextFuncCallParams() []CloudWatchAPIDescribeAlarmHistoryWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmHistoryWithContext
}

func (m *MockCloudWatchAPI) DescribeAlarms(v0 *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	m.mutex.Lock()
	m.histDescribeAlarms = append(m.histDescribeAlarms, CloudWatchAPIDescribeAlarmsParamSet{v0})
	m.mutex.Unlock()
	return m.DescribeAlarmsFunc(v0)
}
func (m *MockCloudWatchAPI) DescribeAlarmsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarms)
}
func (m *MockCloudWatchAPI) DescribeAlarmsFuncCallParams() []CloudWatchAPIDescribeAlarmsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarms
}

func (m *MockCloudWatchAPI) DescribeAlarmsForMetric(v0 *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	m.mutex.Lock()
	m.histDescribeAlarmsForMetric = append(m.histDescribeAlarmsForMetric, CloudWatchAPIDescribeAlarmsForMetricParamSet{v0})
	m.mutex.Unlock()
	return m.DescribeAlarmsForMetricFunc(v0)
}
func (m *MockCloudWatchAPI) DescribeAlarmsForMetricFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsForMetric)
}
func (m *MockCloudWatchAPI) DescribeAlarmsForMetricFuncCallParams() []CloudWatchAPIDescribeAlarmsForMetricParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsForMetric
}

func (m *MockCloudWatchAPI) DescribeAlarmsForMetricRequest(v0 *cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput) {
	m.mutex.Lock()
	m.histDescribeAlarmsForMetricRequest = append(m.histDescribeAlarmsForMetricRequest, CloudWatchAPIDescribeAlarmsForMetricRequestParamSet{v0})
	m.mutex.Unlock()
	return m.DescribeAlarmsForMetricRequestFunc(v0)
}
func (m *MockCloudWatchAPI) DescribeAlarmsForMetricRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsForMetricRequest)
}
func (m *MockCloudWatchAPI) DescribeAlarmsForMetricRequestFuncCallParams() []CloudWatchAPIDescribeAlarmsForMetricRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsForMetricRequest
}

func (m *MockCloudWatchAPI) DescribeAlarmsForMetricWithContext(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsForMetricInput, v2 ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	m.mutex.Lock()
	m.histDescribeAlarmsForMetricWithContext = append(m.histDescribeAlarmsForMetricWithContext, CloudWatchAPIDescribeAlarmsForMetricWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.DescribeAlarmsForMetricWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) DescribeAlarmsForMetricWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsForMetricWithContext)
}
func (m *MockCloudWatchAPI) DescribeAlarmsForMetricWithContextFuncCallParams() []CloudWatchAPIDescribeAlarmsForMetricWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsForMetricWithContext
}

func (m *MockCloudWatchAPI) DescribeAlarmsPages(v0 *cloudwatch.DescribeAlarmsInput, v1 func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error {
	m.mutex.Lock()
	m.histDescribeAlarmsPages = append(m.histDescribeAlarmsPages, CloudWatchAPIDescribeAlarmsPagesParamSet{v0, v1})
	m.mutex.Unlock()
	return m.DescribeAlarmsPagesFunc(v0, v1)
}
func (m *MockCloudWatchAPI) DescribeAlarmsPagesFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsPages)
}
func (m *MockCloudWatchAPI) DescribeAlarmsPagesFuncCallParams() []CloudWatchAPIDescribeAlarmsPagesParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsPages
}

func (m *MockCloudWatchAPI) DescribeAlarmsPagesWithContext(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsInput, v2 func(*cloudwatch.DescribeAlarmsOutput, bool) bool, v3 ...request.Option) error {
	m.mutex.Lock()
	m.histDescribeAlarmsPagesWithContext = append(m.histDescribeAlarmsPagesWithContext, CloudWatchAPIDescribeAlarmsPagesWithContextParamSet{v0, v1, v2, v3})
	m.mutex.Unlock()
	return m.DescribeAlarmsPagesWithContextFunc(v0, v1, v2, v3...)
}
func (m *MockCloudWatchAPI) DescribeAlarmsPagesWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsPagesWithContext)
}
func (m *MockCloudWatchAPI) DescribeAlarmsPagesWithContextFuncCallParams() []CloudWatchAPIDescribeAlarmsPagesWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsPagesWithContext
}

func (m *MockCloudWatchAPI) DescribeAlarmsRequest(v0 *cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput) {
	m.mutex.Lock()
	m.histDescribeAlarmsRequest = append(m.histDescribeAlarmsRequest, CloudWatchAPIDescribeAlarmsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.DescribeAlarmsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) DescribeAlarmsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsRequest)
}
func (m *MockCloudWatchAPI) DescribeAlarmsRequestFuncCallParams() []CloudWatchAPIDescribeAlarmsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsRequest
}

func (m *MockCloudWatchAPI) DescribeAlarmsWithContext(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsInput, v2 ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error) {
	m.mutex.Lock()
	m.histDescribeAlarmsWithContext = append(m.histDescribeAlarmsWithContext, CloudWatchAPIDescribeAlarmsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.DescribeAlarmsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) DescribeAlarmsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDescribeAlarmsWithContext)
}
func (m *MockCloudWatchAPI) DescribeAlarmsWithContextFuncCallParams() []CloudWatchAPIDescribeAlarmsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDescribeAlarmsWithContext
}

func (m *MockCloudWatchAPI) DisableAlarmActions(v0 *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	m.mutex.Lock()
	m.histDisableAlarmActions = append(m.histDisableAlarmActions, CloudWatchAPIDisableAlarmActionsParamSet{v0})
	m.mutex.Unlock()
	return m.DisableAlarmActionsFunc(v0)
}
func (m *MockCloudWatchAPI) DisableAlarmActionsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDisableAlarmActions)
}
func (m *MockCloudWatchAPI) DisableAlarmActionsFuncCallParams() []CloudWatchAPIDisableAlarmActionsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDisableAlarmActions
}

func (m *MockCloudWatchAPI) DisableAlarmActionsRequest(v0 *cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput) {
	m.mutex.Lock()
	m.histDisableAlarmActionsRequest = append(m.histDisableAlarmActionsRequest, CloudWatchAPIDisableAlarmActionsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.DisableAlarmActionsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) DisableAlarmActionsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDisableAlarmActionsRequest)
}
func (m *MockCloudWatchAPI) DisableAlarmActionsRequestFuncCallParams() []CloudWatchAPIDisableAlarmActionsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDisableAlarmActionsRequest
}

func (m *MockCloudWatchAPI) DisableAlarmActionsWithContext(v0 aws.Context, v1 *cloudwatch.DisableAlarmActionsInput, v2 ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error) {
	m.mutex.Lock()
	m.histDisableAlarmActionsWithContext = append(m.histDisableAlarmActionsWithContext, CloudWatchAPIDisableAlarmActionsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.DisableAlarmActionsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) DisableAlarmActionsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histDisableAlarmActionsWithContext)
}
func (m *MockCloudWatchAPI) DisableAlarmActionsWithContextFuncCallParams() []CloudWatchAPIDisableAlarmActionsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histDisableAlarmActionsWithContext
}

func (m *MockCloudWatchAPI) EnableAlarmActions(v0 *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	m.mutex.Lock()
	m.histEnableAlarmActions = append(m.histEnableAlarmActions, CloudWatchAPIEnableAlarmActionsParamSet{v0})
	m.mutex.Unlock()
	return m.EnableAlarmActionsFunc(v0)
}
func (m *MockCloudWatchAPI) EnableAlarmActionsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histEnableAlarmActions)
}
func (m *MockCloudWatchAPI) EnableAlarmActionsFuncCallParams() []CloudWatchAPIEnableAlarmActionsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histEnableAlarmActions
}

func (m *MockCloudWatchAPI) EnableAlarmActionsRequest(v0 *cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput) {
	m.mutex.Lock()
	m.histEnableAlarmActionsRequest = append(m.histEnableAlarmActionsRequest, CloudWatchAPIEnableAlarmActionsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.EnableAlarmActionsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) EnableAlarmActionsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histEnableAlarmActionsRequest)
}
func (m *MockCloudWatchAPI) EnableAlarmActionsRequestFuncCallParams() []CloudWatchAPIEnableAlarmActionsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histEnableAlarmActionsRequest
}

func (m *MockCloudWatchAPI) EnableAlarmActionsWithContext(v0 aws.Context, v1 *cloudwatch.EnableAlarmActionsInput, v2 ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error) {
	m.mutex.Lock()
	m.histEnableAlarmActionsWithContext = append(m.histEnableAlarmActionsWithContext, CloudWatchAPIEnableAlarmActionsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.EnableAlarmActionsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) EnableAlarmActionsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histEnableAlarmActionsWithContext)
}
func (m *MockCloudWatchAPI) EnableAlarmActionsWithContextFuncCallParams() []CloudWatchAPIEnableAlarmActionsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histEnableAlarmActionsWithContext
}

func (m *MockCloudWatchAPI) GetDashboard(v0 *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	m.mutex.Lock()
	m.histGetDashboard = append(m.histGetDashboard, CloudWatchAPIGetDashboardParamSet{v0})
	m.mutex.Unlock()
	return m.GetDashboardFunc(v0)
}
func (m *MockCloudWatchAPI) GetDashboardFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetDashboard)
}
func (m *MockCloudWatchAPI) GetDashboardFuncCallParams() []CloudWatchAPIGetDashboardParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetDashboard
}

func (m *MockCloudWatchAPI) GetDashboardRequest(v0 *cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput) {
	m.mutex.Lock()
	m.histGetDashboardRequest = append(m.histGetDashboardRequest, CloudWatchAPIGetDashboardRequestParamSet{v0})
	m.mutex.Unlock()
	return m.GetDashboardRequestFunc(v0)
}
func (m *MockCloudWatchAPI) GetDashboardRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetDashboardRequest)
}
func (m *MockCloudWatchAPI) GetDashboardRequestFuncCallParams() []CloudWatchAPIGetDashboardRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetDashboardRequest
}

func (m *MockCloudWatchAPI) GetDashboardWithContext(v0 aws.Context, v1 *cloudwatch.GetDashboardInput, v2 ...request.Option) (*cloudwatch.GetDashboardOutput, error) {
	m.mutex.Lock()
	m.histGetDashboardWithContext = append(m.histGetDashboardWithContext, CloudWatchAPIGetDashboardWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.GetDashboardWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) GetDashboardWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetDashboardWithContext)
}
func (m *MockCloudWatchAPI) GetDashboardWithContextFuncCallParams() []CloudWatchAPIGetDashboardWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetDashboardWithContext
}

func (m *MockCloudWatchAPI) GetMetricData(v0 *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	m.mutex.Lock()
	m.histGetMetricData = append(m.histGetMetricData, CloudWatchAPIGetMetricDataParamSet{v0})
	m.mutex.Unlock()
	return m.GetMetricDataFunc(v0)
}
func (m *MockCloudWatchAPI) GetMetricDataFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetMetricData)
}
func (m *MockCloudWatchAPI) GetMetricDataFuncCallParams() []CloudWatchAPIGetMetricDataParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetMetricData
}

func (m *MockCloudWatchAPI) GetMetricDataRequest(v0 *cloudwatch.GetMetricDataInput) (*request.Request, *cloudwatch.GetMetricDataOutput) {
	m.mutex.Lock()
	m.histGetMetricDataRequest = append(m.histGetMetricDataRequest, CloudWatchAPIGetMetricDataRequestParamSet{v0})
	m.mutex.Unlock()
	return m.GetMetricDataRequestFunc(v0)
}
func (m *MockCloudWatchAPI) GetMetricDataRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetMetricDataRequest)
}
func (m *MockCloudWatchAPI) GetMetricDataRequestFuncCallParams() []CloudWatchAPIGetMetricDataRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetMetricDataRequest
}

func (m *MockCloudWatchAPI) GetMetricDataWithContext(v0 aws.Context, v1 *cloudwatch.GetMetricDataInput, v2 ...request.Option) (*cloudwatch.GetMetricDataOutput, error) {
	m.mutex.Lock()
	m.histGetMetricDataWithContext = append(m.histGetMetricDataWithContext, CloudWatchAPIGetMetricDataWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.GetMetricDataWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) GetMetricDataWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetMetricDataWithContext)
}
func (m *MockCloudWatchAPI) GetMetricDataWithContextFuncCallParams() []CloudWatchAPIGetMetricDataWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetMetricDataWithContext
}

func (m *MockCloudWatchAPI) GetMetricStatistics(v0 *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	m.mutex.Lock()
	m.histGetMetricStatistics = append(m.histGetMetricStatistics, CloudWatchAPIGetMetricStatisticsParamSet{v0})
	m.mutex.Unlock()
	return m.GetMetricStatisticsFunc(v0)
}
func (m *MockCloudWatchAPI) GetMetricStatisticsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetMetricStatistics)
}
func (m *MockCloudWatchAPI) GetMetricStatisticsFuncCallParams() []CloudWatchAPIGetMetricStatisticsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetMetricStatistics
}

func (m *MockCloudWatchAPI) GetMetricStatisticsRequest(v0 *cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput) {
	m.mutex.Lock()
	m.histGetMetricStatisticsRequest = append(m.histGetMetricStatisticsRequest, CloudWatchAPIGetMetricStatisticsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.GetMetricStatisticsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) GetMetricStatisticsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetMetricStatisticsRequest)
}
func (m *MockCloudWatchAPI) GetMetricStatisticsRequestFuncCallParams() []CloudWatchAPIGetMetricStatisticsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetMetricStatisticsRequest
}

func (m *MockCloudWatchAPI) GetMetricStatisticsWithContext(v0 aws.Context, v1 *cloudwatch.GetMetricStatisticsInput, v2 ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error) {
	m.mutex.Lock()
	m.histGetMetricStatisticsWithContext = append(m.histGetMetricStatisticsWithContext, CloudWatchAPIGetMetricStatisticsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.GetMetricStatisticsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) GetMetricStatisticsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histGetMetricStatisticsWithContext)
}
func (m *MockCloudWatchAPI) GetMetricStatisticsWithContextFuncCallParams() []CloudWatchAPIGetMetricStatisticsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histGetMetricStatisticsWithContext
}

func (m *MockCloudWatchAPI) ListDashboards(v0 *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	m.mutex.Lock()
	m.histListDashboards = append(m.histListDashboards, CloudWatchAPIListDashboardsParamSet{v0})
	m.mutex.Unlock()
	return m.ListDashboardsFunc(v0)
}
func (m *MockCloudWatchAPI) ListDashboardsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListDashboards)
}
func (m *MockCloudWatchAPI) ListDashboardsFuncCallParams() []CloudWatchAPIListDashboardsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListDashboards
}

func (m *MockCloudWatchAPI) ListDashboardsRequest(v0 *cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput) {
	m.mutex.Lock()
	m.histListDashboardsRequest = append(m.histListDashboardsRequest, CloudWatchAPIListDashboardsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.ListDashboardsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) ListDashboardsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListDashboardsRequest)
}
func (m *MockCloudWatchAPI) ListDashboardsRequestFuncCallParams() []CloudWatchAPIListDashboardsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListDashboardsRequest
}

func (m *MockCloudWatchAPI) ListDashboardsWithContext(v0 aws.Context, v1 *cloudwatch.ListDashboardsInput, v2 ...request.Option) (*cloudwatch.ListDashboardsOutput, error) {
	m.mutex.Lock()
	m.histListDashboardsWithContext = append(m.histListDashboardsWithContext, CloudWatchAPIListDashboardsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.ListDashboardsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) ListDashboardsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListDashboardsWithContext)
}
func (m *MockCloudWatchAPI) ListDashboardsWithContextFuncCallParams() []CloudWatchAPIListDashboardsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListDashboardsWithContext
}

func (m *MockCloudWatchAPI) ListMetrics(v0 *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	m.mutex.Lock()
	m.histListMetrics = append(m.histListMetrics, CloudWatchAPIListMetricsParamSet{v0})
	m.mutex.Unlock()
	return m.ListMetricsFunc(v0)
}
func (m *MockCloudWatchAPI) ListMetricsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListMetrics)
}
func (m *MockCloudWatchAPI) ListMetricsFuncCallParams() []CloudWatchAPIListMetricsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListMetrics
}

func (m *MockCloudWatchAPI) ListMetricsPages(v0 *cloudwatch.ListMetricsInput, v1 func(*cloudwatch.ListMetricsOutput, bool) bool) error {
	m.mutex.Lock()
	m.histListMetricsPages = append(m.histListMetricsPages, CloudWatchAPIListMetricsPagesParamSet{v0, v1})
	m.mutex.Unlock()
	return m.ListMetricsPagesFunc(v0, v1)
}
func (m *MockCloudWatchAPI) ListMetricsPagesFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListMetricsPages)
}
func (m *MockCloudWatchAPI) ListMetricsPagesFuncCallParams() []CloudWatchAPIListMetricsPagesParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListMetricsPages
}

func (m *MockCloudWatchAPI) ListMetricsPagesWithContext(v0 aws.Context, v1 *cloudwatch.ListMetricsInput, v2 func(*cloudwatch.ListMetricsOutput, bool) bool, v3 ...request.Option) error {
	m.mutex.Lock()
	m.histListMetricsPagesWithContext = append(m.histListMetricsPagesWithContext, CloudWatchAPIListMetricsPagesWithContextParamSet{v0, v1, v2, v3})
	m.mutex.Unlock()
	return m.ListMetricsPagesWithContextFunc(v0, v1, v2, v3...)
}
func (m *MockCloudWatchAPI) ListMetricsPagesWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListMetricsPagesWithContext)
}
func (m *MockCloudWatchAPI) ListMetricsPagesWithContextFuncCallParams() []CloudWatchAPIListMetricsPagesWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListMetricsPagesWithContext
}

func (m *MockCloudWatchAPI) ListMetricsRequest(v0 *cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput) {
	m.mutex.Lock()
	m.histListMetricsRequest = append(m.histListMetricsRequest, CloudWatchAPIListMetricsRequestParamSet{v0})
	m.mutex.Unlock()
	return m.ListMetricsRequestFunc(v0)
}
func (m *MockCloudWatchAPI) ListMetricsRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListMetricsRequest)
}
func (m *MockCloudWatchAPI) ListMetricsRequestFuncCallParams() []CloudWatchAPIListMetricsRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListMetricsRequest
}

func (m *MockCloudWatchAPI) ListMetricsWithContext(v0 aws.Context, v1 *cloudwatch.ListMetricsInput, v2 ...request.Option) (*cloudwatch.ListMetricsOutput, error) {
	m.mutex.Lock()
	m.histListMetricsWithContext = append(m.histListMetricsWithContext, CloudWatchAPIListMetricsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.ListMetricsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) ListMetricsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histListMetricsWithContext)
}
func (m *MockCloudWatchAPI) ListMetricsWithContextFuncCallParams() []CloudWatchAPIListMetricsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histListMetricsWithContext
}

func (m *MockCloudWatchAPI) PutDashboard(v0 *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	m.mutex.Lock()
	m.histPutDashboard = append(m.histPutDashboard, CloudWatchAPIPutDashboardParamSet{v0})
	m.mutex.Unlock()
	return m.PutDashboardFunc(v0)
}
func (m *MockCloudWatchAPI) PutDashboardFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutDashboard)
}
func (m *MockCloudWatchAPI) PutDashboardFuncCallParams() []CloudWatchAPIPutDashboardParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutDashboard
}

func (m *MockCloudWatchAPI) PutDashboardRequest(v0 *cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput) {
	m.mutex.Lock()
	m.histPutDashboardRequest = append(m.histPutDashboardRequest, CloudWatchAPIPutDashboardRequestParamSet{v0})
	m.mutex.Unlock()
	return m.PutDashboardRequestFunc(v0)
}
func (m *MockCloudWatchAPI) PutDashboardRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutDashboardRequest)
}
func (m *MockCloudWatchAPI) PutDashboardRequestFuncCallParams() []CloudWatchAPIPutDashboardRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutDashboardRequest
}

func (m *MockCloudWatchAPI) PutDashboardWithContext(v0 aws.Context, v1 *cloudwatch.PutDashboardInput, v2 ...request.Option) (*cloudwatch.PutDashboardOutput, error) {
	m.mutex.Lock()
	m.histPutDashboardWithContext = append(m.histPutDashboardWithContext, CloudWatchAPIPutDashboardWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.PutDashboardWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) PutDashboardWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutDashboardWithContext)
}
func (m *MockCloudWatchAPI) PutDashboardWithContextFuncCallParams() []CloudWatchAPIPutDashboardWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutDashboardWithContext
}

func (m *MockCloudWatchAPI) PutMetricAlarm(v0 *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	m.mutex.Lock()
	m.histPutMetricAlarm = append(m.histPutMetricAlarm, CloudWatchAPIPutMetricAlarmParamSet{v0})
	m.mutex.Unlock()
	return m.PutMetricAlarmFunc(v0)
}
func (m *MockCloudWatchAPI) PutMetricAlarmFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutMetricAlarm)
}
func (m *MockCloudWatchAPI) PutMetricAlarmFuncCallParams() []CloudWatchAPIPutMetricAlarmParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutMetricAlarm
}

func (m *MockCloudWatchAPI) PutMetricAlarmRequest(v0 *cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput) {
	m.mutex.Lock()
	m.histPutMetricAlarmRequest = append(m.histPutMetricAlarmRequest, CloudWatchAPIPutMetricAlarmRequestParamSet{v0})
	m.mutex.Unlock()
	return m.PutMetricAlarmRequestFunc(v0)
}
func (m *MockCloudWatchAPI) PutMetricAlarmRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutMetricAlarmRequest)
}
func (m *MockCloudWatchAPI) PutMetricAlarmRequestFuncCallParams() []CloudWatchAPIPutMetricAlarmRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutMetricAlarmRequest
}

func (m *MockCloudWatchAPI) PutMetricAlarmWithContext(v0 aws.Context, v1 *cloudwatch.PutMetricAlarmInput, v2 ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error) {
	m.mutex.Lock()
	m.histPutMetricAlarmWithContext = append(m.histPutMetricAlarmWithContext, CloudWatchAPIPutMetricAlarmWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.PutMetricAlarmWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) PutMetricAlarmWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutMetricAlarmWithContext)
}
func (m *MockCloudWatchAPI) PutMetricAlarmWithContextFuncCallParams() []CloudWatchAPIPutMetricAlarmWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutMetricAlarmWithContext
}

func (m *MockCloudWatchAPI) PutMetricData(v0 *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	m.mutex.Lock()
	m.histPutMetricData = append(m.histPutMetricData, CloudWatchAPIPutMetricDataParamSet{v0})
	m.mutex.Unlock()
	return m.PutMetricDataFunc(v0)
}
func (m *MockCloudWatchAPI) PutMetricDataFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutMetricData)
}
func (m *MockCloudWatchAPI) PutMetricDataFuncCallParams() []CloudWatchAPIPutMetricDataParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutMetricData
}

func (m *MockCloudWatchAPI) PutMetricDataRequest(v0 *cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput) {
	m.mutex.Lock()
	m.histPutMetricDataRequest = append(m.histPutMetricDataRequest, CloudWatchAPIPutMetricDataRequestParamSet{v0})
	m.mutex.Unlock()
	return m.PutMetricDataRequestFunc(v0)
}
func (m *MockCloudWatchAPI) PutMetricDataRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutMetricDataRequest)
}
func (m *MockCloudWatchAPI) PutMetricDataRequestFuncCallParams() []CloudWatchAPIPutMetricDataRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutMetricDataRequest
}

func (m *MockCloudWatchAPI) PutMetricDataWithContext(v0 aws.Context, v1 *cloudwatch.PutMetricDataInput, v2 ...request.Option) (*cloudwatch.PutMetricDataOutput, error) {
	m.mutex.Lock()
	m.histPutMetricDataWithContext = append(m.histPutMetricDataWithContext, CloudWatchAPIPutMetricDataWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.PutMetricDataWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) PutMetricDataWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histPutMetricDataWithContext)
}
func (m *MockCloudWatchAPI) PutMetricDataWithContextFuncCallParams() []CloudWatchAPIPutMetricDataWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histPutMetricDataWithContext
}

func (m *MockCloudWatchAPI) SetAlarmState(v0 *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	m.mutex.Lock()
	m.histSetAlarmState = append(m.histSetAlarmState, CloudWatchAPISetAlarmStateParamSet{v0})
	m.mutex.Unlock()
	return m.SetAlarmStateFunc(v0)
}
func (m *MockCloudWatchAPI) SetAlarmStateFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histSetAlarmState)
}
func (m *MockCloudWatchAPI) SetAlarmStateFuncCallParams() []CloudWatchAPISetAlarmStateParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histSetAlarmState
}

func (m *MockCloudWatchAPI) SetAlarmStateRequest(v0 *cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput) {
	m.mutex.Lock()
	m.histSetAlarmStateRequest = append(m.histSetAlarmStateRequest, CloudWatchAPISetAlarmStateRequestParamSet{v0})
	m.mutex.Unlock()
	return m.SetAlarmStateRequestFunc(v0)
}
func (m *MockCloudWatchAPI) SetAlarmStateRequestFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histSetAlarmStateRequest)
}
func (m *MockCloudWatchAPI) SetAlarmStateRequestFuncCallParams() []CloudWatchAPISetAlarmStateRequestParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histSetAlarmStateRequest
}

func (m *MockCloudWatchAPI) SetAlarmStateWithContext(v0 aws.Context, v1 *cloudwatch.SetAlarmStateInput, v2 ...request.Option) (*cloudwatch.SetAlarmStateOutput, error) {
	m.mutex.Lock()
	m.histSetAlarmStateWithContext = append(m.histSetAlarmStateWithContext, CloudWatchAPISetAlarmStateWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.SetAlarmStateWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) SetAlarmStateWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histSetAlarmStateWithContext)
}
func (m *MockCloudWatchAPI) SetAlarmStateWithContextFuncCallParams() []CloudWatchAPISetAlarmStateWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histSetAlarmStateWithContext
}

func (m *MockCloudWatchAPI) WaitUntilAlarmExists(v0 *cloudwatch.DescribeAlarmsInput) error {
	m.mutex.Lock()
	m.histWaitUntilAlarmExists = append(m.histWaitUntilAlarmExists, CloudWatchAPIWaitUntilAlarmExistsParamSet{v0})
	m.mutex.Unlock()
	return m.WaitUntilAlarmExistsFunc(v0)
}
func (m *MockCloudWatchAPI) WaitUntilAlarmExistsFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histWaitUntilAlarmExists)
}
func (m *MockCloudWatchAPI) WaitUntilAlarmExistsFuncCallParams() []CloudWatchAPIWaitUntilAlarmExistsParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histWaitUntilAlarmExists
}

func (m *MockCloudWatchAPI) WaitUntilAlarmExistsWithContext(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsInput, v2 ...request.WaiterOption) error {
	m.mutex.Lock()
	m.histWaitUntilAlarmExistsWithContext = append(m.histWaitUntilAlarmExistsWithContext, CloudWatchAPIWaitUntilAlarmExistsWithContextParamSet{v0, v1, v2})
	m.mutex.Unlock()
	return m.WaitUntilAlarmExistsWithContextFunc(v0, v1, v2...)
}
func (m *MockCloudWatchAPI) WaitUntilAlarmExistsWithContextFuncCallCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.histWaitUntilAlarmExistsWithContext)
}
func (m *MockCloudWatchAPI) WaitUntilAlarmExistsWithContextFuncCallParams() []CloudWatchAPIWaitUntilAlarmExistsWithContextParamSet {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.histWaitUntilAlarmExistsWithContext
}

func (m *MockCloudWatchAPI) defaultDeleteAlarmsFunc(v0 *cloudwatch.DeleteAlarmsInput) (*cloudwatch.DeleteAlarmsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDeleteAlarmsRequestFunc(v0 *cloudwatch.DeleteAlarmsInput) (*request.Request, *cloudwatch.DeleteAlarmsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDeleteAlarmsWithContextFunc(v0 aws.Context, v1 *cloudwatch.DeleteAlarmsInput, v2 ...request.Option) (*cloudwatch.DeleteAlarmsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDeleteDashboardsFunc(v0 *cloudwatch.DeleteDashboardsInput) (*cloudwatch.DeleteDashboardsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDeleteDashboardsRequestFunc(v0 *cloudwatch.DeleteDashboardsInput) (*request.Request, *cloudwatch.DeleteDashboardsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDeleteDashboardsWithContextFunc(v0 aws.Context, v1 *cloudwatch.DeleteDashboardsInput, v2 ...request.Option) (*cloudwatch.DeleteDashboardsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmHistoryFunc(v0 *cloudwatch.DescribeAlarmHistoryInput) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmHistoryPagesFunc(v0 *cloudwatch.DescribeAlarmHistoryInput, v1 func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmHistoryPagesWithContextFunc(v0 aws.Context, v1 *cloudwatch.DescribeAlarmHistoryInput, v2 func(*cloudwatch.DescribeAlarmHistoryOutput, bool) bool, v3 ...request.Option) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmHistoryRequestFunc(v0 *cloudwatch.DescribeAlarmHistoryInput) (*request.Request, *cloudwatch.DescribeAlarmHistoryOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmHistoryWithContextFunc(v0 aws.Context, v1 *cloudwatch.DescribeAlarmHistoryInput, v2 ...request.Option) (*cloudwatch.DescribeAlarmHistoryOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsFunc(v0 *cloudwatch.DescribeAlarmsInput) (*cloudwatch.DescribeAlarmsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsForMetricFunc(v0 *cloudwatch.DescribeAlarmsForMetricInput) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsForMetricRequestFunc(v0 *cloudwatch.DescribeAlarmsForMetricInput) (*request.Request, *cloudwatch.DescribeAlarmsForMetricOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsForMetricWithContextFunc(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsForMetricInput, v2 ...request.Option) (*cloudwatch.DescribeAlarmsForMetricOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsPagesFunc(v0 *cloudwatch.DescribeAlarmsInput, v1 func(*cloudwatch.DescribeAlarmsOutput, bool) bool) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsPagesWithContextFunc(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsInput, v2 func(*cloudwatch.DescribeAlarmsOutput, bool) bool, v3 ...request.Option) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsRequestFunc(v0 *cloudwatch.DescribeAlarmsInput) (*request.Request, *cloudwatch.DescribeAlarmsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDescribeAlarmsWithContextFunc(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsInput, v2 ...request.Option) (*cloudwatch.DescribeAlarmsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDisableAlarmActionsFunc(v0 *cloudwatch.DisableAlarmActionsInput) (*cloudwatch.DisableAlarmActionsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDisableAlarmActionsRequestFunc(v0 *cloudwatch.DisableAlarmActionsInput) (*request.Request, *cloudwatch.DisableAlarmActionsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultDisableAlarmActionsWithContextFunc(v0 aws.Context, v1 *cloudwatch.DisableAlarmActionsInput, v2 ...request.Option) (*cloudwatch.DisableAlarmActionsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultEnableAlarmActionsFunc(v0 *cloudwatch.EnableAlarmActionsInput) (*cloudwatch.EnableAlarmActionsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultEnableAlarmActionsRequestFunc(v0 *cloudwatch.EnableAlarmActionsInput) (*request.Request, *cloudwatch.EnableAlarmActionsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultEnableAlarmActionsWithContextFunc(v0 aws.Context, v1 *cloudwatch.EnableAlarmActionsInput, v2 ...request.Option) (*cloudwatch.EnableAlarmActionsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetDashboardFunc(v0 *cloudwatch.GetDashboardInput) (*cloudwatch.GetDashboardOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetDashboardRequestFunc(v0 *cloudwatch.GetDashboardInput) (*request.Request, *cloudwatch.GetDashboardOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetDashboardWithContextFunc(v0 aws.Context, v1 *cloudwatch.GetDashboardInput, v2 ...request.Option) (*cloudwatch.GetDashboardOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetMetricDataFunc(v0 *cloudwatch.GetMetricDataInput) (*cloudwatch.GetMetricDataOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetMetricDataRequestFunc(v0 *cloudwatch.GetMetricDataInput) (*request.Request, *cloudwatch.GetMetricDataOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetMetricDataWithContextFunc(v0 aws.Context, v1 *cloudwatch.GetMetricDataInput, v2 ...request.Option) (*cloudwatch.GetMetricDataOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetMetricStatisticsFunc(v0 *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetMetricStatisticsRequestFunc(v0 *cloudwatch.GetMetricStatisticsInput) (*request.Request, *cloudwatch.GetMetricStatisticsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultGetMetricStatisticsWithContextFunc(v0 aws.Context, v1 *cloudwatch.GetMetricStatisticsInput, v2 ...request.Option) (*cloudwatch.GetMetricStatisticsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultListDashboardsFunc(v0 *cloudwatch.ListDashboardsInput) (*cloudwatch.ListDashboardsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultListDashboardsRequestFunc(v0 *cloudwatch.ListDashboardsInput) (*request.Request, *cloudwatch.ListDashboardsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultListDashboardsWithContextFunc(v0 aws.Context, v1 *cloudwatch.ListDashboardsInput, v2 ...request.Option) (*cloudwatch.ListDashboardsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultListMetricsFunc(v0 *cloudwatch.ListMetricsInput) (*cloudwatch.ListMetricsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultListMetricsPagesFunc(v0 *cloudwatch.ListMetricsInput, v1 func(*cloudwatch.ListMetricsOutput, bool) bool) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultListMetricsPagesWithContextFunc(v0 aws.Context, v1 *cloudwatch.ListMetricsInput, v2 func(*cloudwatch.ListMetricsOutput, bool) bool, v3 ...request.Option) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultListMetricsRequestFunc(v0 *cloudwatch.ListMetricsInput) (*request.Request, *cloudwatch.ListMetricsOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultListMetricsWithContextFunc(v0 aws.Context, v1 *cloudwatch.ListMetricsInput, v2 ...request.Option) (*cloudwatch.ListMetricsOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutDashboardFunc(v0 *cloudwatch.PutDashboardInput) (*cloudwatch.PutDashboardOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutDashboardRequestFunc(v0 *cloudwatch.PutDashboardInput) (*request.Request, *cloudwatch.PutDashboardOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutDashboardWithContextFunc(v0 aws.Context, v1 *cloudwatch.PutDashboardInput, v2 ...request.Option) (*cloudwatch.PutDashboardOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutMetricAlarmFunc(v0 *cloudwatch.PutMetricAlarmInput) (*cloudwatch.PutMetricAlarmOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutMetricAlarmRequestFunc(v0 *cloudwatch.PutMetricAlarmInput) (*request.Request, *cloudwatch.PutMetricAlarmOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutMetricAlarmWithContextFunc(v0 aws.Context, v1 *cloudwatch.PutMetricAlarmInput, v2 ...request.Option) (*cloudwatch.PutMetricAlarmOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutMetricDataFunc(v0 *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutMetricDataRequestFunc(v0 *cloudwatch.PutMetricDataInput) (*request.Request, *cloudwatch.PutMetricDataOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultPutMetricDataWithContextFunc(v0 aws.Context, v1 *cloudwatch.PutMetricDataInput, v2 ...request.Option) (*cloudwatch.PutMetricDataOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultSetAlarmStateFunc(v0 *cloudwatch.SetAlarmStateInput) (*cloudwatch.SetAlarmStateOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultSetAlarmStateRequestFunc(v0 *cloudwatch.SetAlarmStateInput) (*request.Request, *cloudwatch.SetAlarmStateOutput) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultSetAlarmStateWithContextFunc(v0 aws.Context, v1 *cloudwatch.SetAlarmStateInput, v2 ...request.Option) (*cloudwatch.SetAlarmStateOutput, error) {
	return nil, nil
}
func (m *MockCloudWatchAPI) defaultWaitUntilAlarmExistsFunc(v0 *cloudwatch.DescribeAlarmsInput) error {
	return nil
}
func (m *MockCloudWatchAPI) defaultWaitUntilAlarmExistsWithContextFunc(v0 aws.Context, v1 *cloudwatch.DescribeAlarmsInput, v2 ...request.WaiterOption) error {
	return nil
}
