package sam

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Application struct {
	Metadata     iacTypes.Metadata
	LocationPath iacTypes.StringValue
	Location     Location
}

type Location struct {
	Metadata        iacTypes.Metadata
	ApplicationID   iacTypes.StringValue
	SemanticVersion iacTypes.StringValue
}
