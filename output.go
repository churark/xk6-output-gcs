package gcs

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/output"
)

type Entry struct {
	Metric string            `json:"metric"`
	Type   string            `json:"type"`
	Value  float64           `json:"value"`
	Tags   map[string]string `json:"tags"`
	Time   int64             `json:"time"`
}

func init() {
	output.RegisterExtension("gcs", New)
}

type Output struct {
	output.SampleBuffer

	scli            *storage.Client
	cfg             *Config
	logger          logrus.FieldLogger
	periodicFlusher *output.PeriodicFlusher
	mu              sync.Mutex
	res             []Entry
}

func New(params output.Params) (output.Output, error) {
	ctx := context.Background()
	scli, err := storage.NewClient(ctx)
	if err != nil {
		params.Logger.Errorf("failed to create gcs client: %v", err)
		return nil, err
	}

	cfg, err := NewConfig(ctx)
	if err != nil {
		params.Logger.Errorf("failed to create config: %v", err)
		return nil, err
	}

	return &Output{
		scli:   scli,
		cfg:    cfg,
		logger: params.Logger,
		mu:     sync.Mutex{},
		res:    []Entry{},
	}, nil
}

func (o *Output) Description() string {
	return "xk6-output-gcs"
}

func (o *Output) Start() error {
	o.logger.Debug("Starting...")

	periodicFlusher, err := output.NewPeriodicFlusher(time.Second*10, o.flush)
	if err != nil {
		return err
	}
	o.periodicFlusher = periodicFlusher

	o.logger.Debug("Started!")
	return nil
}

func (o *Output) Stop() error {
	o.logger.Debug("Stopping...")
	defer o.logger.Debug("Stopped!")

	o.periodicFlusher.Stop()

	ret, err := json.Marshal(o.res)
	if err != nil {
		o.logger.Errorf("failed to marshal results: %v", err)
		return err
	}

	o.logger.Debug("Writing results...")
	ctx := context.Background()
	sc := o.scli.Bucket(o.cfg.Bucket).Object(fmt.Sprintf("%d.json", time.Now().Unix())).NewWriter(ctx)
	if _, err := sc.Write(ret); err != nil {
		o.logger.Errorf("failed to write results: %v", err)
		return err
	}

	if err := o.scli.Close(); err != nil {
		o.logger.Errorf("failed to close storage client: %", err)
	}
	o.logger.Debug("Wrote results!")

	return nil
}

func (o *Output) flush() {
	samplesContainers := o.GetBufferedSamples() //nolint:typecheck
	for i := range samplesContainers {
		samples := samplesContainers[i].GetSamples()

		for _, sample := range samples {
			mappedEntry := Entry{
				Metric: sample.Metric.Name,
				Type:   sample.Metric.Type.String(),
				Value:  sample.Value,
				Tags:   sample.GetTags().Map(),
				Time:   sample.Time.Unix(),
			}
			o.add(mappedEntry)
		}
	}
}

func (o *Output) add(e Entry) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.res = append(o.res, e)
}
