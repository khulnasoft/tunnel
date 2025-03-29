package ec2

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Subnet struct {
	Metadata            iacTypes.Metadata
	MapPublicIpOnLaunch iacTypes.BoolValue
}
