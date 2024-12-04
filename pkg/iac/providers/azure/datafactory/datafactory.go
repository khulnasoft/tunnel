package datafactory

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type DataFactory struct {
	DataFactories []Factory
}

type Factory struct {
	Metadata            iacTypes.Metadata
	EnablePublicNetwork iacTypes.BoolValue
}
