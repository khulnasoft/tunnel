package rdb

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type DBSecurityGroup struct {
	Metadata    iacTypes.Metadata
	Description iacTypes.StringValue
	CIDRs       []iacTypes.StringValue
}
