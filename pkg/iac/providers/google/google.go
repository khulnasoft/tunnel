package google

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/bigquery"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/dns"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/gke"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/iam"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/kms"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/sql"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/storage"
)

type Google struct {
	BigQuery bigquery.BigQuery
	Compute  compute.Compute
	DNS      dns.DNS
	GKE      gke.GKE
	KMS      kms.KMS
	IAM      iam.IAM
	SQL      sql.SQL
	Storage  storage.Storage
}
