package accessanalyzer

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws/accessanalyzer"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an AccessAnalyzer instance
func Adapt(cfFile parser.FileContext) accessanalyzer.AccessAnalyzer {
	return accessanalyzer.AccessAnalyzer{
		Analyzers: getAccessAnalyzer(cfFile),
	}
}
