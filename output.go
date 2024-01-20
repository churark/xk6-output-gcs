package gcs

import (
	"go.k6.io/k6/metrics"
	"go.k6.io/k6/output"
)

func init() {
	output.RegisterExtension("gcs", New)
}

type Output struct{}

func New(params output.Params) (output.Output, error) {
	return &Output{}, nil
}

func (o *Output) Description() string {
	return "xk6-output-gcs"
}

func (o *Output) Start() error {
	return nil
}

func (o *Output) Stop() error {
	return nil
}

func (o *Output) AddMetricSamples(samples []metrics.SampleContainer) {
	return
}
