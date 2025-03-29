package eks

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/eks"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an EKS instance
func Adapt(cfFile parser.FileContext) eks.EKS {
	return eks.EKS{
		Clusters: getClusters(cfFile),
	}
}
