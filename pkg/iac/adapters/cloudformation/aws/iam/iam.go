package iam

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/iam"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

// Adapt adapts an IAM instance
func Adapt(cfFile parser.FileContext) iam.IAM {
	return iam.IAM{
		PasswordPolicy: iam.PasswordPolicy{
			Metadata:             iacTypes.NewUnmanagedMetadata(),
			ReusePreventionCount: iacTypes.IntDefault(0, iacTypes.NewUnmanagedMetadata()),
			RequireLowercase:     iacTypes.BoolDefault(false, iacTypes.NewUnmanagedMetadata()),
			RequireUppercase:     iacTypes.BoolDefault(false, iacTypes.NewUnmanagedMetadata()),
			RequireNumbers:       iacTypes.BoolDefault(false, iacTypes.NewUnmanagedMetadata()),
			RequireSymbols:       iacTypes.BoolDefault(false, iacTypes.NewUnmanagedMetadata()),
			MaxAgeDays:           iacTypes.IntDefault(0, iacTypes.NewUnmanagedMetadata()),
			MinimumLength:        iacTypes.IntDefault(0, iacTypes.NewUnmanagedMetadata()),
		},
		Policies: getPolicies(cfFile),
		Groups:   getGroups(cfFile),
		Users:    getUsers(cfFile),
		Roles:    getRoles(cfFile),
	}
}
