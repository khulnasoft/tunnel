package gemspec

import (
	"context"
	"os"
	"path/filepath"
	"regexp"

	"github.com/khulnasoft/tunnel/pkg/dependency/parser/ruby/gemspec"
	"github.com/khulnasoft/tunnel/pkg/fanal/analyzer"
	"github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language"
	"github.com/khulnasoft/tunnel/pkg/fanal/types"
)

func init() {
	analyzer.RegisterAnalyzer(&gemspecLibraryAnalyzer{})
}

const version = 1

var fileRegex = regexp.MustCompile(`.*/specifications/.+\.gemspec`)

type gemspecLibraryAnalyzer struct{}

func (a gemspecLibraryAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	return language.AnalyzePackage(types.GemSpec, input.FilePath, input.Content,
		gemspec.NewParser(), input.Options.FileChecksum)
}

func (a gemspecLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	return fileRegex.MatchString(filepath.ToSlash(filePath))
}

func (a gemspecLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeGemSpec
}

func (a gemspecLibraryAnalyzer) Version() int {
	return version
}
