package sslcertificate

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type ServerCertificate struct {
	Metadata   iacTypes.Metadata
	Expiration iacTypes.TimeValue
}
