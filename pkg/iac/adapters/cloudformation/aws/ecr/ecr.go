package ecr

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/ecr"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an ECR instance
func Adapt(cfFile parser.FileContext) ecr.ECR {
	return ecr.ECR{
		Repositories: getRepositories(cfFile),
	}
}
