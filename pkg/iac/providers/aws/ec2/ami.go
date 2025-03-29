package ec2

import iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"

type RequestedAMI struct {
	Metadata iacTypes.Metadata
	Owners   iacTypes.StringValueList
}
