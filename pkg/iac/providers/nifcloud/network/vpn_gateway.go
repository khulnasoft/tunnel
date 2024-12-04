package network

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type VpnGateway struct {
	Metadata      iacTypes.Metadata
	SecurityGroup iacTypes.StringValue
}
