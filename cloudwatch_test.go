package imperial

//go:generate go-mockgen github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface -i CloudWatchAPI -o mock_cloudwatch_api_test.go -f

import (
	"time"

	"github.com/aphistic/sweet"
	cloudwatch "github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/efritz/glock"
	. "github.com/onsi/gomega"
)

type CloudwatchSuite struct{}

func (s *CloudwatchSuite) TestReport(t sweet.T) {
	reporter, api, clock := makeCloudwatchReporter()

	t1 := time.Now()
	t2 := time.Now().Add(time.Minute)

	clock.SetCurrent(t1)
	reporter.Report("a", 1)
	reporter.Report("b", 2)
	reporter.Report("c", 3)
	clock.SetCurrent(t2)
	reporter.Report("d", 4)
	reporter.Report("e", 5)

	// Wait until publish
	Eventually(api.PutMetricDataFuncCallCount).Should(Equal(1))

	// Inspect payload
	input := api.PutMetricDataFuncCallParams()[0].Arg0
	Expect(*input.Namespace).To(Equal("ns"))
	Expect(input.MetricData).To(HaveLen(5))
	Expect(*input.MetricData[0].MetricName).To(Equal("a"))
	Expect(*input.MetricData[1].MetricName).To(Equal("b"))
	Expect(*input.MetricData[2].MetricName).To(Equal("c"))
	Expect(*input.MetricData[3].MetricName).To(Equal("d"))
	Expect(*input.MetricData[4].MetricName).To(Equal("e"))
	Expect(*input.MetricData[0].Value).To(Equal(float64(1)))
	Expect(*input.MetricData[1].Value).To(Equal(float64(2)))
	Expect(*input.MetricData[2].Value).To(Equal(float64(3)))
	Expect(*input.MetricData[3].Value).To(Equal(float64(4)))
	Expect(*input.MetricData[4].Value).To(Equal(float64(5)))
	Expect(*input.MetricData[0].Timestamp).To(Equal(t1))
	Expect(*input.MetricData[3].Timestamp).To(Equal(t2))
}
func (s *CloudwatchSuite) TestReportWithAttributes(t sweet.T) {
	reporter, api, _ := makeCloudwatchReporter(
		WithCloudwatchReportConfigs(WithAttributes(map[string]string{
			"x": "xv",
			"y": "xy",
		})),
	)

	reporter.Report("a", 1)
	reporter.Report("b", 2, WithAttributes(map[string]string{"z": "z1"}))
	reporter.Report("c", 3)
	reporter.Report("d", 4, WithAttributes(map[string]string{"z": "z2"}))
	reporter.Report("e", 5)

	// Wait until publish
	Eventually(api.PutMetricDataFuncCallCount).Should(Equal(1))

	d1 := serializeCloudwatchDimension("x", "xv")
	d2 := serializeCloudwatchDimension("y", "xy")
	d3 := serializeCloudwatchDimension("z", "z1")
	d4 := serializeCloudwatchDimension("z", "z2")

	// Inspect payload
	input := api.PutMetricDataFuncCallParams()[0].Arg0
	Expect(input.MetricData).To(HaveLen(5))
	Expect(input.MetricData[0].Dimensions).To(ConsistOf(d1, d2))
	Expect(input.MetricData[1].Dimensions).To(ConsistOf(d1, d2, d3))
	Expect(input.MetricData[2].Dimensions).To(ConsistOf(d1, d2))
	Expect(input.MetricData[3].Dimensions).To(ConsistOf(d1, d2, d4))
	Expect(input.MetricData[4].Dimensions).To(ConsistOf(d1, d2))
}

func (s *CloudwatchSuite) TestReportMultipleNamespaces(t sweet.T) {
	reporter, api, _ := makeCloudwatchReporter(
		WithCloudwatchReportConfigs(WithAttributes(map[string]string{
			"x": "xv",
			"y": "xy",
		})),
	)

	for i := 0; i < 5; i++ {
		reporter.Report("a", i, WithCloudwatchNamespace("foo"))
		reporter.Report("b", i, WithCloudwatchNamespace("bar"))
		reporter.Report("c", i, WithCloudwatchNamespace("baz"))

		if i%2 == 0 {
			// Will not trigger a batch to send
			reporter.Report("d", i, WithCloudwatchNamespace("bonk"))
		}
	}

	// Wait until publish
	Eventually(api.PutMetricDataFuncCallCount).Should(Equal(3))

	namespaces := []string{}
	for _, call := range api.PutMetricDataFuncCallParams() {
		namespaces = append(namespaces, *call.Arg0.Namespace)
	}

	Expect(namespaces).To(ConsistOf("foo", "bar", "baz"))
}

func (s *CloudwatchSuite) TestMultipleBatches(t sweet.T) {
	reporter, api, _ := makeCloudwatchReporter()

	for i := 0; i < 3; i++ {
		reporter.Report("a", 10*i+1)
		reporter.Report("b", 10*i+2)
		reporter.Report("c", 10*i+3)
		reporter.Report("d", 10*i+4)
		reporter.Report("e", 10*i+5)

		// Wait until publish
		Eventually(api.PutMetricDataFuncCallCount).Should(Equal(i + 1))

		// Inspect payload
		input := api.PutMetricDataFuncCallParams()[0].Arg0
		Expect(input.MetricData).To(HaveLen(5))
		Expect(*input.MetricData[0].Value).To(Equal(float64(10*i + 1)))
		Expect(*input.MetricData[1].Value).To(Equal(float64(10*i + 2)))
		Expect(*input.MetricData[2].Value).To(Equal(float64(10*i + 3)))
		Expect(*input.MetricData[3].Value).To(Equal(float64(10*i + 4)))
		Expect(*input.MetricData[4].Value).To(Equal(float64(10*i + 5)))
	}

	reporter.Report("a", 41)
	reporter.Report("b", 42)
	Consistently(api.PutMetricDataFuncCallCount).Should(Equal(3))
}

func (s *CloudwatchSuite) TestPartialBatchTick(t sweet.T) {
	reporter, api, clock := makeCloudwatchReporter()

	for i := 0; i < 3; i++ {
		reporter.Report("a", 10*i+1)
		Consistently(api.PutMetricDataFuncCallCount).Should(Equal(i))
		clock.Advance(time.Second * 5)

		// Wait until publish
		Eventually(api.PutMetricDataFuncCallCount).Should(Equal(i + 1))

		// Inspect payload
		input := api.PutMetricDataFuncCallParams()[0].Arg0
		Expect(input.MetricData).To(HaveLen(1))
		Expect(*input.MetricData[0].Value).To(Equal(float64(10*i + 1)))
	}
}

func (s *CloudwatchSuite) TestFullBuffer(t sweet.T) {
	var (
		block            = make(chan struct{})
		reporter, api, _ = makeCloudwatchReporter(
			WithCloudwatchBatchSize(5),
			WithCloudwatchBufferSize(25))
	)

	api.PutMetricDataFunc = func(*cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
		<-block
		return nil, nil
	}

	for i := 0; i < 500; i++ {
		if i == 450 {
			// Wait for writes to propagate
			<-time.After(time.Millisecond * 50)
			close(block)
			<-time.After(time.Millisecond * 50)
		}

		reporter.Report("a", i)
	}

	// Wait until publish
	Eventually(api.PutMetricDataFuncCallCount).Should(BeNumerically("~", 15, 5))
}

func (s *CloudwatchSuite) TestShutdown(t sweet.T) {
	var (
		sync             = make(chan struct{})
		block            = make(chan struct{})
		reporter, api, _ = makeCloudwatchReporter()
	)

	api.PutMetricDataFunc = func(input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
		if *input.Namespace == "c" {
			<-block
		}

		return nil, nil
	}

	reporter.Report("a", 1, WithCloudwatchNamespace("a"))
	reporter.Report("b", 2, WithCloudwatchNamespace("b"))
	reporter.Report("c", 3, WithCloudwatchNamespace("c"))
	reporter.Report("d", 4, WithCloudwatchNamespace("d"))
	reporter.Report("e", 5, WithCloudwatchNamespace("e"))

	go func() {
		defer close(sync)
		reporter.Shutdown()
	}()

	Consistently(sync).ShouldNot(BeClosed())
	close(block)
	Eventually(sync).Should(BeClosed())
}

//
// Constructors

func makeCloudwatchReporter(configs ...CloudwatchConfigFunc) (Reporter, *MockCloudWatchAPI, *glock.MockClock) {
	var (
		api   = NewMockCloudWatchAPI()
		clock = glock.NewMockClock()
	)

	reporter := NewCloudwatchReporter("ns", append(
		[]CloudwatchConfigFunc{
			WithCloudwatchAPI(api),
			WithCloudwatchClock(clock),
			WithCloudwatchBatchSize(5),
		},
		configs...,
	)...)

	return reporter, api, clock
}
