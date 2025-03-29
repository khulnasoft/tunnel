package neptune

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/neptune"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Neptune instance
func Adapt(cfFile parser.FileContext) neptune.Neptune {
	return neptune.Neptune{
		Clusters: getClusters(cfFile),
	}
}
