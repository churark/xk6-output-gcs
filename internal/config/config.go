package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	ProjectID      string `env:"GCS_PROJECT_ID,required"`
	Bucket         string `env:"GCS_BUCKET,required"`
	CredentialJSON string `env:"GCS_CREDENTIAL_JSON"`
	CredentialPath string `env:"GCS_CREDENTIAL_PATH"`
}

func NewConfig(ctx context.Context) (*Config, error) {
	conf := &Config{}
	if err := envconfig.Process(ctx, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
