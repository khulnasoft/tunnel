package compute

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type ProjectMetadata struct {
	Metadata      iacTypes.Metadata
	EnableOSLogin iacTypes.BoolValue
}
