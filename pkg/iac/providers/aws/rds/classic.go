package rds

import (
	"github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Classic struct {
	DBSecurityGroups []DBSecurityGroup
}

type DBSecurityGroup struct {
	Metadata types.Metadata
}
