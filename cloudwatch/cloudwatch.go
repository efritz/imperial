package cloudwatch

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/efritz/glock"

	"github.com/efritz/imperial/base"
)

type (
	Reporter struct {
		logger       base.Logger
		clock        glock.Clock
		configs      []base.ConfigFunc
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

func NewReporter(namespace string, configs ...ConfigFunc) *Reporter {
	config := newConfig(namespace)
	for _, f := range configs {
		f(config)
	}

	return &Reporter{
		logger:       config.logger,
		clock:        config.clock,
		configs:      config.configs,
		api:          makeAPI(config),
		batchSize:    config.batchSize,
		bufferSize:   config.bufferSize,
		tickDuration: config.tickDuration,
		namespaces:   map[string]chan *cloudwatch.MetricDatum{},
		mutex:        &sync.RWMutex{},
		once:         &sync.Once{},
		wg:           &sync.WaitGroup{},
	}
}

func (c *Reporter) Report(name string, value int, configs ...base.ConfigFunc) {
	var (
		options   = base.ApplyConfigs(c.configs, configs)
		namespace = options.Namespace
		datum     = &cloudwatch.MetricDatum{
			MetricName: aws.String(name),
			Timestamp:  aws.Time(c.clock.Now()),
			Value:      aws.Float64(float64(value)),
			Unit:       aws.String(string(options.Unit)),
			Dimensions: serializeDimensions(options.Attributes),
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

func (c *Reporter) Shutdown() {
	c.once.Do(func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		for _, ch := range c.namespaces {
			close(ch)
		}
	})

	c.wg.Wait()
}

func (c *Reporter) ensurePublisher(namespace string) {
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

func (c *Reporter) publish(namespace string, ch <-chan *cloudwatch.MetricDatum) {
	defer c.wg.Done()

	var (
		ticker = c.clock.NewTicker(c.tickDuration)
		data   = []*cloudwatch.MetricDatum{}
	)

	putMetricData := func() {
		if len(data) == 0 {
			return
		}

		input := &cloudwatch.PutMetricDataInput{
			Namespace:  aws.String(namespace),
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

func serializeDimensions(attributes map[string]string) []*cloudwatch.Dimension {
	dimensions := make([]*cloudwatch.Dimension, 0, len(attributes))
	for key, value := range attributes {
		dimensions = append(dimensions, serializeDimension(key, value))
	}

	return dimensions
}

func serializeDimension(key, value string) *cloudwatch.Dimension {
	return &cloudwatch.Dimension{
		Name:  aws.String(key),
		Value: aws.String(value),
	}
}
