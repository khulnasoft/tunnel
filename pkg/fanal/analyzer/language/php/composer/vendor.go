package composer

import (
	"context"
	"os"
	"path/filepath"

	"github.com/khulnasoft/tunnel/pkg/dependency/parser/php/composer"
	"github.com/khulnasoft/tunnel/pkg/fanal/analyzer"
	"github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language"
	"github.com/khulnasoft/tunnel/pkg/fanal/types"
)

func init() {
	analyzer.RegisterAnalyzer(&composerVendorAnalyzer{})
}

const (
	composerInstalledAnalyzerVersion = 1
)

// composerVendorAnalyzer analyzes 'installed.json'
type composerVendorAnalyzer struct{}

func (a composerVendorAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	return language.Analyze(types.ComposerVendor, input.FilePath, input.Content, composer.NewParser())
}

func (a composerVendorAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	return filepath.Base(filePath) == types.ComposerInstalledJson
}

func (a composerVendorAnalyzer) Type() analyzer.Type {
	return analyzer.TypeComposerVendor
}

func (a composerVendorAnalyzer) Version() int {
	return composerInstalledAnalyzerVersion
}
