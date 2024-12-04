package dockerfile

import (
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/dockerfile/parser"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/generic"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/options"
	"github.com/khulnasoft/tunnel/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner {
	return generic.NewScanner("Dockerfile", types.SourceDockerfile, generic.ParseFunc(parser.Parse), opts...)
}
