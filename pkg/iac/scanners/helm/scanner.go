package helm

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/liamg/memoryfs"
	"helm.sh/helm/v3/pkg/chartutil"

	"github.com/khulnasoft/tunnel/pkg/iac/detection"
	"github.com/khulnasoft/tunnel/pkg/iac/ignore"
	"github.com/khulnasoft/tunnel/pkg/iac/rego"
	"github.com/khulnasoft/tunnel/pkg/iac/scan"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/helm/parser"
	kparser "github.com/khulnasoft/tunnel/pkg/iac/scanners/kubernetes/parser"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/options"
	"github.com/khulnasoft/tunnel/pkg/iac/types"
	"github.com/khulnasoft/tunnel/pkg/log"
)

var _ scanners.FSScanner = (*Scanner)(nil)
var _ options.ConfigurableScanner = (*Scanner)(nil)

type Scanner struct {
	mu            sync.Mutex
	logger        *log.Logger
	options       []options.ScannerOption
	parserOptions []parser.Option
	regoScanner   *rego.Scanner
}

// New creates a new Scanner
func New(opts ...options.ScannerOption) *Scanner {
	s := &Scanner{
		options: opts,
		logger:  log.WithPrefix("helm scanner"),
	}

	for _, option := range opts {
		option(s)
	}
	return s
}

func (s *Scanner) addParserOptions(opts ...parser.Option) {
	s.parserOptions = append(s.parserOptions, opts...)
}

func (s *Scanner) Name() string {
	return "Helm"
}

func (s *Scanner) ScanFS(ctx context.Context, fsys fs.FS, dir string) (scan.Results, error) {

	if err := s.initRegoScanner(fsys); err != nil {
		return nil, fmt.Errorf("failed to init rego scanner: %w", err)
	}

	var results []scan.Result
	if err := fs.WalkDir(fsys, dir, func(filePath string, d fs.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if detection.IsArchive(filePath) {
			scanResults, err := s.getScanResults(filePath, ctx, fsys)
			if err != nil {
				return err
			}
			results = append(results, scanResults...)
		} else if path.Base(filePath) == chartutil.ChartfileName {
			if scanResults, err := s.getScanResults(filepath.Dir(filePath), ctx, fsys); err != nil {
				return err
			} else {
				results = append(results, scanResults...)
			}
			return fs.SkipDir
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return results, nil

}

func (s *Scanner) getScanResults(path string, ctx context.Context, target fs.FS) (results []scan.Result, err error) {
	helmParser, err := parser.New(path, s.parserOptions...)
	if err != nil {
		return nil, err
	}

	if err := helmParser.ParseFS(ctx, target, path); err != nil {
		return nil, err
	}

	chartFiles, err := helmParser.RenderedChartFiles()
	if err != nil { // not valid helm, maybe some other yaml etc., abort
		s.logger.Error(
			"Failed to render Chart files",
			log.FilePath(path), log.Err(err),
		)
		return nil, nil
	}

	for _, file := range chartFiles {
		file := file
		s.logger.Debug("Processing rendered chart file", log.FilePath(file.TemplateFilePath))

		ignoreRules := ignore.Parse(file.ManifestContent, file.TemplateFilePath, "")
		manifests, err := kparser.Parse(ctx, strings.NewReader(file.ManifestContent), file.TemplateFilePath)
		if err != nil {
			return nil, fmt.Errorf("unmarshal yaml: %w", err)
		}
		for _, manifest := range manifests {
			fileResults, err := s.regoScanner.ScanInput(ctx, types.SourceKubernetes, rego.Input{
				Path:     file.TemplateFilePath,
				Contents: manifest,
				FS:       target,
			})
			if err != nil {
				return nil, fmt.Errorf("scanning error: %w", err)
			}

			if len(fileResults) > 0 {
				renderedFS := memoryfs.New()
				if err := renderedFS.MkdirAll(filepath.Dir(file.TemplateFilePath), fs.ModePerm); err != nil {
					return nil, err
				}
				if err := renderedFS.WriteLazyFile(file.TemplateFilePath, func() (io.Reader, error) {
					return strings.NewReader(file.ManifestContent), nil
				}, fs.ModePerm); err != nil {
					return nil, err
				}
				fileResults.SetSourceAndFilesystem(helmParser.ChartSource, renderedFS, detection.IsArchive(helmParser.ChartSource))
				fileResults.Ignore(ignoreRules, nil)
			}

			results = append(results, fileResults...)
		}

	}
	return results, nil
}

func (s *Scanner) initRegoScanner(srcFS fs.FS) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.regoScanner != nil {
		return nil
	}
	regoScanner := rego.NewScanner(s.options...)
	if err := regoScanner.LoadPolicies(srcFS); err != nil {
		return err
	}
	s.regoScanner = regoScanner
	return nil
}
