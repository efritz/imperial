package imperial

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/efritz/glock"
)

type (
	CloudwatchReporter struct {
		logger       Logger
		clock        glock.Clock
		configs      []ConfigFunc
		api          cloudwatchiface.CloudWatchAPI
		batchSize    int
		bufferSize   int
		tickDuration time.Duration
		namespaces   map[string]chan *cloudwatch.MetricDatum
		mutex        *sync.RWMutex
		once         *sync.Once
		wg           *sync.WaitGroup
	}
)

func NewCloudwatchReporter(namespace string, configs ...CloudwatchConfigFunc) *CloudwatchReporter {
	config := newCloudwatchConfig(namespace)
	for _, f := range configs {
		f(config)
	}

	return &CloudwatchReporter{
		logger:       config.logger,
		configs:      config.configs,
		api:          makeCloudwatchAPI(config),
		batchSize:    config.batchSize,
		bufferSize:   config.bufferSize,
		tickDuration: config.tickDuration,
		namespaces:   map[string]chan *cloudwatch.MetricDatum{},
		mutex:        &sync.RWMutex{},
		once:         &sync.Once{},
		wg:           &sync.WaitGroup{},
	}
}

func (c *CloudwatchReporter) Report(name string, value int, configs ...ConfigFunc) {
	var (
		options   = applyConfigs(c.configs, configs)
		namespace = options.cloudwatchNamespace
		datum     = &cloudwatch.MetricDatum{
			MetricName: stringptr(name),
			Timestamp:  timeptr(c.clock.Now()),
			Value:      float64ptr(float64(value)),
			Unit:       stringptr(string(options.cloudwatchUnit)),
			Dimensions: serializeCloudwatchDimensions(options.attributes),
		}
	)

	c.ensurePublisher(namespace)

	for {
		select {
		case c.namespaces[namespace] <- datum:
			return
		default:
		}

		select {
		case <-c.namespaces[namespace]:
			c.logger.Printf(
				"Cloudwatch buffer for namespace %s full, dropping oldest datum",
				namespace,
			)
		default:
		}
	}
}

func (c *CloudwatchReporter) Shutdown() {
	c.once.Do(func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		for _, ch := range c.namespaces {
			close(ch)
		}
	})

	c.wg.Wait()
}

func (c *CloudwatchReporter) ensurePublisher(namespace string) {
	c.mutex.RLock()
	if _, ok := c.namespaces[namespace]; ok {
		c.mutex.RUnlock()
		return
	}

	c.mutex.RUnlock()
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.namespaces[namespace]; !ok {
		ch := make(chan *cloudwatch.MetricDatum, c.bufferSize)
		c.namespaces[namespace] = ch

		c.wg.Add(1)
		go c.publish(namespace, ch)
	}
}

func (c *CloudwatchReporter) publish(namespace string, ch <-chan *cloudwatch.MetricDatum) {
	var (
		ticker = c.clock.NewTicker(c.tickDuration)
		data   = []*cloudwatch.MetricDatum{}
	)

	putMetricData := func() {
		if len(data) == 0 {
			return
		}

		input := &cloudwatch.PutMetricDataInput{
			Namespace:  stringptr(namespace),
			MetricData: data,
		}

		if _, err := c.api.PutMetricData(input); err != nil {
			c.logger.Printf(
				"Failed to publish data for to Cloudwatch namespace %s (%s)",
				namespace,
				err.Error(),
			)
		}

		data = data[:0]
	}

loop:
	for {
		select {
		case datum, ok := <-ch:
			if !ok {
				break loop
			}

			data = append(data, datum)

			if len(data) < c.batchSize {
				continue
			}

		case <-ticker.Chan():
		}

		putMetricData()
	}

	putMetricData()
}

func serializeCloudwatchDimensions(attributes map[string]string) []*cloudwatch.Dimension {
	dimensions := make([]*cloudwatch.Dimension, 0, len(attributes))
	for key, value := range attributes {
		dimensions = append(dimensions, &cloudwatch.Dimension{
			Name:  stringptr(key),
			Value: stringptr(value),
		})
	}

	return dimensions
}
