package github

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type BranchProtection struct {
	Metadata             iacTypes.Metadata
	RequireSignedCommits iacTypes.BoolValue
}

func (b BranchProtection) RequiresSignedCommits() bool {
	return b.RequireSignedCommits.IsTrue()
}
