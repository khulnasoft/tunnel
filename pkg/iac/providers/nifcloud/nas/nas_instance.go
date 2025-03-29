package nas

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type NASInstance struct {
	Metadata  iacTypes.Metadata
	NetworkID iacTypes.StringValue
}
