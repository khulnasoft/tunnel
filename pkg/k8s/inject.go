//go:build wireinject
// +build wireinject

package k8s

import (
	"github.com/google/wire"

	"github.com/khulnasoft/tunnel/pkg/cache"
)

func initializeScanK8s(localArtifactCache cache.LocalArtifactCache) *ScanKubernetes {
	wire.Build(ScanSuperSet)
	return &ScanKubernetes{}
}
