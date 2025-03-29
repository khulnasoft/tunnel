package kinesis

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/kinesis"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Kinesis instance
func Adapt(cfFile parser.FileContext) kinesis.Kinesis {
	return kinesis.Kinesis{
		Streams: getStreams(cfFile),
	}
}
