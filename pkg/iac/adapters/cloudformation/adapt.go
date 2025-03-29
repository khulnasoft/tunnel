package cloudformation

import (
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/cloudformation/aws"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
	"github.com/khulnasoft/tunnel/pkg/iac/state"
)

// Adapt adapts the Cloudformation instance
func Adapt(cfFile parser.FileContext) *state.State {
	return &state.State{
		AWS: aws.Adapt(cfFile),
	}
}
