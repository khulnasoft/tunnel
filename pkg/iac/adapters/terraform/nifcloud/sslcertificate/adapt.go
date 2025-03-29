package sslcertificate

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/nifcloud/sslcertificate"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) sslcertificate.SSLCertificate {
	return sslcertificate.SSLCertificate{
		ServerCertificates: adaptServerCertificates(modules),
	}
}
