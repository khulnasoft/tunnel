package efs

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type EFS struct {
	FileSystems []FileSystem
}

type FileSystem struct {
	Metadata  iacTypes.Metadata
	Encrypted iacTypes.BoolValue
}
