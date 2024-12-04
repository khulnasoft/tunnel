package nifcloud

import (
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud/computing"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud/dns"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud/nas"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud/network"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud/rdb"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud/sslcertificate"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/nifcloud"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) nifcloud.Nifcloud {
	return nifcloud.Nifcloud{
		Computing:      computing.Adapt(modules),
		DNS:            dns.Adapt(modules),
		NAS:            nas.Adapt(modules),
		Network:        network.Adapt(modules),
		RDB:            rdb.Adapt(modules),
		SSLCertificate: sslcertificate.Adapt(modules),
	}
}
