package cocoapods

import (
	"context"
	"os"

	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/dependency/parser/swift/cocoapods"
	"github.com/khulnasoft/tunnel/pkg/fanal/analyzer"
	"github.com/khulnasoft/tunnel/pkg/fanal/analyzer/language"
	"github.com/khulnasoft/tunnel/pkg/fanal/types"
)

func init() {
	analyzer.RegisterAnalyzer(&cocoaPodsLockAnalyzer{})
}

const (
	version = 1
)

// cocoaPodsLockAnalyzer analyzes Podfile.lock
type cocoaPodsLockAnalyzer struct{}

func (a cocoaPodsLockAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	p := cocoapods.NewParser()
	res, err := language.Analyze(types.Cocoapods, input.FilePath, input.Content, p)
	if err != nil {
		return nil, xerrors.Errorf("%s parse error: %w", input.FilePath, err)
	}
	return res, nil
}

func (a cocoaPodsLockAnalyzer) Required(_ string, fileInfo os.FileInfo) bool {
	return fileInfo.Name() == types.CocoaPodsLock
}

func (a cocoaPodsLockAnalyzer) Type() analyzer.Type {
	return analyzer.TypeCocoaPods
}

func (a cocoaPodsLockAnalyzer) Version() int {
	return version
}
