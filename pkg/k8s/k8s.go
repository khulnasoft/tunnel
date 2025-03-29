package k8s

import (
	"context"

	"github.com/google/wire"

	ftypes "github.com/khulnasoft/tunnel/pkg/fanal/types"
	"github.com/khulnasoft/tunnel/pkg/scan"
	"github.com/khulnasoft/tunnel/pkg/scan/local"
	"github.com/khulnasoft/tunnel/pkg/types"
)

// ScanSuperSet binds the dependencies for k8s
var ScanSuperSet = wire.NewSet(
	local.SuperSet,
	wire.Bind(new(scan.Backend), new(local.Service)),
	NewScanKubernetes,
)

// ScanKubernetes implements the scanner
type ScanKubernetes struct {
	localScanner local.Service
}

// NewScanKubernetes is the factory method for scanner
func NewScanKubernetes(s local.Service) *ScanKubernetes {
	return &ScanKubernetes{localScanner: s}
}

// NewKubernetesScanner is the factory method for scanner
func NewKubernetesScanner() *ScanKubernetes {
	return initializeScanK8s(nil)
}

// Scan scans k8s core components and return it findings
func (sk ScanKubernetes) Scan(ctx context.Context, target types.ScanTarget, options types.ScanOptions) (types.Results, ftypes.OS, error) {
	return sk.localScanner.ScanTarget(ctx, target, options)
}
