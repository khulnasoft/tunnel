package analyzer

import (
	"context"
	"slices"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/fanal/types"
	"github.com/khulnasoft/tunnel/pkg/log"
	"github.com/khulnasoft/tunnel/pkg/misconf"
)

var configAnalyzerConstructors = make(map[Type]configAnalyzerConstructor)

type configAnalyzerConstructor func(ConfigAnalyzerOptions) (ConfigAnalyzer, error)

// RegisterConfigAnalyzer adds a constructor of config analyzer
func RegisterConfigAnalyzer(t Type, init configAnalyzerConstructor) {
	configAnalyzerConstructors[t] = init
}

// DeregisterConfigAnalyzer is mainly for testing
func DeregisterConfigAnalyzer(t Type) {
	delete(configAnalyzerConstructors, t)
}

// ConfigAnalyzer defines an interface for analyzer of container image config
type ConfigAnalyzer interface {
	Type() Type
	Version() int
	Analyze(ctx context.Context, input ConfigAnalysisInput) (*ConfigAnalysisResult, error)
	Required(osFound types.OS) bool
}

// ConfigAnalyzerOptions is used to initialize config analyzers
type ConfigAnalyzerOptions struct {
	FilePatterns         []string
	DisabledAnalyzers    []Type
	MisconfScannerOption misconf.ScannerOption
	SecretScannerOption  SecretScannerOption
}

type ConfigAnalysisInput struct {
	OS     types.OS
	Config *v1.ConfigFile
}

type ConfigAnalysisResult struct {
	Misconfiguration *types.Misconfiguration
	Secret           *types.Secret
	HistoryPackages  types.Packages
}

func (r *ConfigAnalysisResult) Merge(newResult *ConfigAnalysisResult) {
	if newResult == nil {
		return
	}
	if newResult.Misconfiguration != nil {
		r.Misconfiguration = newResult.Misconfiguration
	}
	if newResult.Secret != nil {
		r.Secret = newResult.Secret
	}
	if newResult.HistoryPackages != nil {
		r.HistoryPackages = newResult.HistoryPackages
	}
}

type ConfigAnalyzerGroup struct {
	configAnalyzers []ConfigAnalyzer
}

func NewConfigAnalyzerGroup(opts ConfigAnalyzerOptions) (ConfigAnalyzerGroup, error) {
	var g ConfigAnalyzerGroup
	for t, newConfigAnalyzer := range configAnalyzerConstructors {
		// Skip the handler if it is disabled
		if slices.Contains(opts.DisabledAnalyzers, t) {
			continue
		}
		a, err := newConfigAnalyzer(opts)
		if err != nil {
			return ConfigAnalyzerGroup{}, xerrors.Errorf("config analyzer %s initialize error: %w", t, err)
		}

		g.configAnalyzers = append(g.configAnalyzers, a)
	}

	return g, nil
}

// AnalyzerVersions returns analyzer version identifier used for cache keys.
func (ag *ConfigAnalyzerGroup) AnalyzerVersions() Versions {
	versions := make(map[string]int)
	for _, ca := range ag.configAnalyzers {
		versions[string(ca.Type())] = ca.Version()
	}
	return Versions{
		Analyzers: versions,
	}
}

func (ag *ConfigAnalyzerGroup) AnalyzeImageConfig(ctx context.Context, targetOS types.OS, config *v1.ConfigFile) *ConfigAnalysisResult {
	input := ConfigAnalysisInput{
		OS:     targetOS,
		Config: config,
	}
	result := new(ConfigAnalysisResult)
	for _, a := range ag.configAnalyzers {
		if !a.Required(targetOS) {
			continue
		}

		r, err := a.Analyze(ctx, input)
		if err != nil {
			log.Debug("Image config analysis error", log.Err(err))
			continue
		}

		result.Merge(r)
	}
	return result
}
