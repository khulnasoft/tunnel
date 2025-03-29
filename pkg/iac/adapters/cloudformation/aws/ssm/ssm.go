package ssm

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/ssm"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an SSM instance
func Adapt(cfFile parser.FileContext) ssm.SSM {
	return ssm.SSM{
		Secrets: getSecrets(cfFile),
	}
}
