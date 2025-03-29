package network

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Router struct {
	Metadata          iacTypes.Metadata
	SecurityGroup     iacTypes.StringValue
	NetworkInterfaces []NetworkInterface
}
