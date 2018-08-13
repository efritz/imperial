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

type Reporter struct {
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

var _ base.SimpleReporter = &Reporter{}

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

func (r *Reporter) Report(name string, value float64, configs ...base.ConfigFunc) {
	var (
		options   = base.ApplyConfigs(r.configs, configs)
		namespace = options.Namespace
		datum     = &cloudwatch.MetricDatum{
			MetricName: aws.String(name),
			Timestamp:  aws.Time(r.clock.Now()),
			Value:      aws.Float64(value),
			Unit:       aws.String(string(options.Unit)),
			Dimensions: serializeDimensions(options.Attributes),
		}
	)

	r.ensurePublisher(namespace)

	for {
		select {
		case r.namespaces[namespace] <- datum:
			return
		default:
		}

		select {
		case <-r.namespaces[namespace]:
			r.logger.Printf(
				"Cloudwatch buffer for namespace %s full, dropping oldest datum",
				namespace,
			)

		default:
		}
	}
}

func (r *Reporter) Shutdown() {
	r.once.Do(func() {
		r.mutex.Lock()
		defer r.mutex.Unlock()

		for _, ch := range r.namespaces {
			close(ch)
		}
	})

	r.wg.Wait()
}

func (r *Reporter) ensurePublisher(namespace string) {
	r.mutex.RLock()
	if _, ok := r.namespaces[namespace]; ok {
		r.mutex.RUnlock()
		return
	}

	r.mutex.RUnlock()
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.namespaces[namespace]; !ok {
		ch := make(chan *cloudwatch.MetricDatum, r.bufferSize)
		r.namespaces[namespace] = ch

		r.wg.Add(1)
		go r.publish(namespace, ch)
	}
}

func (r *Reporter) publish(namespace string, ch <-chan *cloudwatch.MetricDatum) {
	defer r.wg.Done()

	var (
		data   = []*cloudwatch.MetricDatum{}
		ticker = r.clock.NewTicker(r.tickDuration)
	)

	putMetricData := func() {
		if len(data) == 0 {
			return
		}

		input := &cloudwatch.PutMetricDataInput{
			Namespace:  aws.String(namespace),
			MetricData: data,
		}

		if _, err := r.api.PutMetricData(input); err != nil {
			r.logger.Printf(
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

			if len(data) < r.batchSize {
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
