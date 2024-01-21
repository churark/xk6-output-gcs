package config

import (
	"context"
	"testing"

	envconfig "github.com/sethvargo/go-envconfig"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	ctx := context.Background()

	patterns := []struct {
		name  string
		setup func(t *testing.T)
		want  *Config
		err   error
	}{
		{
			name: "default",
			setup: func(t *testing.T) {
				t.Helper()
			},
			want: nil,
			err:  envconfig.ErrMissingRequired,
		},
		{
			name: "set envs",
			setup: func(t *testing.T) {
				t.Helper()

				t.Setenv("GCS_PROJECT_ID", "project-id")
				t.Setenv("GCS_BUCKET", "bucket")
				t.Setenv("GCS_CREDENTIAL_JSON", "{'json': 'json'}")
				t.Setenv("GCS_CREDENTIAL_PATH", "credential-path")
			},
			want: &Config{
				ProjectID:      "project-id",
				Bucket:         "bucket",
				CredentialJSON: "{'json': 'json'}",
				CredentialPath: "credential-path",
			},
		},
	}

	for _, tt := range patterns {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			got, err := NewConfig(ctx)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}
