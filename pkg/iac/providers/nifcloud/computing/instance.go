package computing

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Instance struct {
	Metadata          iacTypes.Metadata
	SecurityGroup     iacTypes.StringValue
	NetworkInterfaces []NetworkInterface
}

type NetworkInterface struct {
	Metadata  iacTypes.Metadata
	NetworkID iacTypes.StringValue
}
