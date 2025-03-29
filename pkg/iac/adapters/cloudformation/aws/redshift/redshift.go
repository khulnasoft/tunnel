package redshift

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/redshift"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a RedShift instance
func Adapt(cfFile parser.FileContext) redshift.Redshift {
	return redshift.Redshift{
		Clusters:          getClusters(cfFile),
		SecurityGroups:    getSecurityGroups(cfFile),
		ClusterParameters: getParameters(cfFile),
		ReservedNodes:     nil,
	}
}
