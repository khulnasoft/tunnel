package kms

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type KMS struct {
	Keys []Key
}

const (
	KeyUsageSignAndVerify = "SIGN_VERIFY"
)

type Key struct {
	Metadata        iacTypes.Metadata
	Usage           iacTypes.StringValue
	RotationEnabled iacTypes.BoolValue
}
