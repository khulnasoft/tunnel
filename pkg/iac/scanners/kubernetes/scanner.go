package kubernetes

import (
	"context"
	"io"

	"github.com/khulnasoft/tunnel/pkg/iac/scanners/generic"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/kubernetes/parser"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/options"
	"github.com/khulnasoft/tunnel/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner {
	return generic.NewScanner("Kubernetes", types.SourceKubernetes, generic.ParseFunc(parse), opts...)
}

func parse(ctx context.Context, r io.Reader, path string) (any, error) {
	return parser.Parse(ctx, r, path)
}
