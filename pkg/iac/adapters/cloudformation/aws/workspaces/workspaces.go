package workspaces

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/workspaces"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Workspaces instance
func Adapt(cfFile parser.FileContext) workspaces.WorkSpaces {
	return workspaces.WorkSpaces{
		WorkSpaces: getWorkSpaces(cfFile),
	}
}
