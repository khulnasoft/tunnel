package workspaces

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type WorkSpaces struct {
	WorkSpaces []WorkSpace
}

type WorkSpace struct {
	Metadata   iacTypes.Metadata
	RootVolume Volume
	UserVolume Volume
}

type Volume struct {
	Metadata   iacTypes.Metadata
	Encryption Encryption
}

type Encryption struct {
	Metadata iacTypes.Metadata
	Enabled  iacTypes.BoolValue
}
