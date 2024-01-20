package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	ProjectID string `env:"GCS_PROJECT_ID,required"`
	Bucket    string `env:"GCS_BUCKET,required"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{}
	if err := envconfig.Process(ctx, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
