package cloudwatch

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/efritz/glock"

	"github.com/efritz/imperial/base"
)

type (
	config struct {
		logger       base.Logger
		clock        glock.Clock
		configs      []base.ConfigFunc
		session      *session.Session
		api          cloudwatchiface.CloudWatchAPI
		batchSize    int
		bufferSize   int
		tickDuration time.Duration
	}

	ConfigFunc func(r *config)
)

func newConfig(namespace string) *config {
	return &config{
		logger:       base.NewNilLogger(),
		clock:        glock.NewRealClock(),
		configs:      []base.ConfigFunc{base.WithNamespace(namespace)},
		batchSize:    5000,
		bufferSize:   1000,
		tickDuration: time.Second * 5,
	}
}

func WithReportConfigs(configs ...base.ConfigFunc) ConfigFunc {
	return func(c *config) { c.configs = append(c.configs, configs...) }
}

func WithBatchSize(batchSize int) ConfigFunc {
	return func(c *config) { c.batchSize = batchSize }
}

func WithBufferSize(bufferSize int) ConfigFunc {
	return func(c *config) { c.bufferSize = bufferSize }
}

func WithTickDuration(tickDuration time.Duration) ConfigFunc {
	return func(c *config) { c.tickDuration = tickDuration }
}

func WithLogger(logger base.Logger) ConfigFunc {
	return func(c *config) { c.logger = logger }
}

func WithClock(clock glock.Clock) ConfigFunc {
	return func(c *config) { c.clock = clock }
}

func WithSession(session *session.Session) ConfigFunc {
	return func(c *config) { c.session = session }
}

func WithAPI(api cloudwatchiface.CloudWatchAPI) ConfigFunc {
	return func(c *config) { c.api = api }
}

//
//

func makeAPI(config *config) cloudwatchiface.CloudWatchAPI {
	if config.api != nil {
		return config.api
	}

	awsSession := config.session
	if awsSession == nil {
		awsSession = session.New()
	}

	return cloudwatch.New(awsSession)
}
