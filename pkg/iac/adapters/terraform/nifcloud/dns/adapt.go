package dns

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/nifcloud/dns"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) dns.DNS {
	return dns.DNS{
		Records: adaptRecords(modules),
	}
}
