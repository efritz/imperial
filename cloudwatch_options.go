package imperial

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/efritz/glock"
)

type (
	cloudwatchConfig struct {
		logger       Logger
		clock        glock.Clock
		configs      []ConfigFunc
		session      *session.Session
		api          cloudwatchiface.CloudWatchAPI
		batchSize    int
		bufferSize   int
		tickDuration time.Duration
	}

	CloudwatchConfigFunc func(r *cloudwatchConfig)
)

func newCloudwatchConfig(namespace string) *cloudwatchConfig {
	return &cloudwatchConfig{
		logger:       NewNilLogger(),
		clock:        glock.NewRealClock(),
		configs:      []ConfigFunc{WithCloudwatchNamespace(namespace)},
		batchSize:    5000,
		bufferSize:   1000,
		tickDuration: time.Second * 5,
	}
}

func WithCloudwatchReportConfigs(configs ...ConfigFunc) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.configs = append(c.configs, configs...) }
}

func WithCloudwatchBatchSize(batchSize int) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.batchSize = batchSize }
}

func WithCloudwatchBufferSize(bufferSize int) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.bufferSize = bufferSize }
}

func WithCloudwatchTickDuration(tickDuration time.Duration) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.tickDuration = tickDuration }
}

func WithCloudwatchLogger(logger Logger) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.logger = logger }
}

func WithCloudwatchClock(clock glock.Clock) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.clock = clock }
}

func WithCloudwatchSession(session *session.Session) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.session = session }
}

func WithCloudwatchAPI(api cloudwatchiface.CloudWatchAPI) CloudwatchConfigFunc {
	return func(c *cloudwatchConfig) { c.api = api }
}

//
//

func makeCloudwatchAPI(config *cloudwatchConfig) cloudwatchiface.CloudWatchAPI {
	if config.api != nil {
		return config.api
	}

	awsSession := config.session
	if awsSession == nil {
		awsSession = session.New()
	}

	return cloudwatch.New(awsSession)
}
