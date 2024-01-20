package gcs

import (
	k6output "go.k6.io/k6/output"

	"github.com/churark/xk6-output-gcs/internal/output"
)

func init() {
	k6output.RegisterExtension("gcs", output.New)
}
