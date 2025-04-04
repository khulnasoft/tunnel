package google

import (
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/bigquery"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/dns"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/gke"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/iam"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/kms"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/sql"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google/storage"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) google.Google {
	return google.Google{
		BigQuery: bigquery.Adapt(modules),
		Compute:  compute.Adapt(modules),
		DNS:      dns.Adapt(modules),
		GKE:      gke.Adapt(modules),
		KMS:      kms.Adapt(modules),
		IAM:      iam.Adapt(modules),
		SQL:      sql.Adapt(modules),
		Storage:  storage.Adapt(modules),
	}
}
