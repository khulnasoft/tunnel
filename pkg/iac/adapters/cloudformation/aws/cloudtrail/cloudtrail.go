package cloudtrail

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/cloudtrail"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a CloudTrail instance
func Adapt(cfFile parser.FileContext) cloudtrail.CloudTrail {
	return cloudtrail.CloudTrail{
		Trails: getCloudTrails(cfFile),
	}
}
