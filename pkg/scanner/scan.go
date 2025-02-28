package scanner

import (
	"context"

	"github.com/google/wire"
	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/cache"
	"github.com/khulnasoft/tunnel/pkg/clock"
	"github.com/khulnasoft/tunnel/pkg/fanal/artifact"
	aimage "github.com/khulnasoft/tunnel/pkg/fanal/artifact/image"
	flocal "github.com/khulnasoft/tunnel/pkg/fanal/artifact/local"
	"github.com/khulnasoft/tunnel/pkg/fanal/artifact/repo"
	"github.com/khulnasoft/tunnel/pkg/fanal/artifact/sbom"
	"github.com/khulnasoft/tunnel/pkg/fanal/artifact/vm"
	"github.com/khulnasoft/tunnel/pkg/fanal/image"
	ftypes "github.com/khulnasoft/tunnel/pkg/fanal/types"
	"github.com/khulnasoft/tunnel/pkg/log"
	"github.com/khulnasoft/tunnel/pkg/report"
	"github.com/khulnasoft/tunnel/pkg/rpc/client"
	"github.com/khulnasoft/tunnel/pkg/scanner/local"
	"github.com/khulnasoft/tunnel/pkg/types"
)

///////////////
// Standalone
///////////////

// StandaloneSuperSet is used in the standalone mode
var StandaloneSuperSet = wire.NewSet(
	// Cache
	cache.New,
	wire.Bind(new(cache.ArtifactCache), new(cache.Cache)),
	wire.Bind(new(cache.LocalArtifactCache), new(cache.Cache)),

	local.SuperSet,
	wire.Bind(new(Driver), new(local.Scanner)),
	NewScanner,
)

// StandaloneDockerSet binds docker dependencies
var StandaloneDockerSet = wire.NewSet(
	image.NewContainerImage,
	aimage.NewArtifact,
	StandaloneSuperSet,
)

// StandaloneArchiveSet binds archive scan dependencies
var StandaloneArchiveSet = wire.NewSet(
	image.NewArchiveImage,
	aimage.NewArtifact,
	StandaloneSuperSet,
)

// StandaloneFilesystemSet binds filesystem dependencies
var StandaloneFilesystemSet = wire.NewSet(
	flocal.ArtifactSet,
	StandaloneSuperSet,
)

// StandaloneRepositorySet binds repository dependencies
var StandaloneRepositorySet = wire.NewSet(
	repo.ArtifactSet,
	StandaloneSuperSet,
)

// StandaloneSBOMSet binds sbom dependencies
var StandaloneSBOMSet = wire.NewSet(
	sbom.NewArtifact,
	StandaloneSuperSet,
)

// StandaloneVMSet binds vm dependencies
var StandaloneVMSet = wire.NewSet(
	vm.ArtifactSet,
	StandaloneSuperSet,
)

/////////////////
// Client/Server
/////////////////

// RemoteSuperSet is used in the client mode
var RemoteSuperSet = wire.NewSet(
	// Cache
	cache.NewRemoteCache,
	wire.Bind(new(cache.ArtifactCache), new(*cache.RemoteCache)), // No need for LocalArtifactCache

	client.NewScanner,
	wire.Value([]client.Option(nil)),
	wire.Bind(new(Driver), new(client.Scanner)),
	NewScanner,
)

// RemoteFilesystemSet binds filesystem dependencies for client/server mode
var RemoteFilesystemSet = wire.NewSet(
	flocal.ArtifactSet,
	RemoteSuperSet,
)

// RemoteRepositorySet binds repository dependencies for client/server mode
var RemoteRepositorySet = wire.NewSet(
	repo.ArtifactSet,
	RemoteSuperSet,
)

// RemoteSBOMSet binds sbom dependencies for client/server mode
var RemoteSBOMSet = wire.NewSet(
	sbom.NewArtifact,
	RemoteSuperSet,
)

// RemoteVMSet binds vm dependencies for client/server mode
var RemoteVMSet = wire.NewSet(
	vm.ArtifactSet,
	RemoteSuperSet,
)

// RemoteDockerSet binds remote docker dependencies
var RemoteDockerSet = wire.NewSet(
	aimage.NewArtifact,
	image.NewContainerImage,
	RemoteSuperSet,
)

// RemoteArchiveSet binds remote archive dependencies
var RemoteArchiveSet = wire.NewSet(
	aimage.NewArtifact,
	image.NewArchiveImage,
	RemoteSuperSet,
)

// Scanner implements the Artifact and Driver operations
type Scanner struct {
	driver   Driver
	artifact artifact.Artifact
}

// Driver defines operations of scanner
type Driver interface {
	Scan(ctx context.Context, target, artifactKey string, blobKeys []string, options types.ScanOptions) (
		results types.Results, osFound ftypes.OS, err error)
}

// NewScanner is the factory method of Scanner
func NewScanner(driver Driver, ar artifact.Artifact) Scanner {
	return Scanner{
		driver:   driver,
		artifact: ar,
	}
}

// ScanArtifact scans the artifacts and returns results
func (s Scanner) ScanArtifact(ctx context.Context, options types.ScanOptions) (types.Report, error) {
	artifactInfo, err := s.artifact.Inspect(ctx)
	if err != nil {
		return types.Report{}, xerrors.Errorf("failed analysis: %w", err)
	}
	defer func() {
		if err := s.artifact.Clean(artifactInfo); err != nil {
			log.Warn("Failed to clean the artifact",
				log.String("artifact", artifactInfo.Name), log.Err(err))
		}
	}()

	results, osFound, err := s.driver.Scan(ctx, artifactInfo.Name, artifactInfo.ID, artifactInfo.BlobIDs, options)
	if err != nil {
		return types.Report{}, xerrors.Errorf("scan failed: %w", err)
	}

	ptros := &osFound
	if osFound.Detected() && osFound.Eosl {
		log.Warn("This OS version is no longer supported by the distribution",
			log.String("family", string(osFound.Family)), log.String("version", osFound.Name))
		log.Warn("The vulnerability detection may be insufficient because security updates are not provided")
	} else if !osFound.Detected() {
		ptros = nil
	}

	return types.Report{
		SchemaVersion: report.SchemaVersion,
		CreatedAt:     clock.Now(ctx),
		ArtifactName:  artifactInfo.Name,
		ArtifactType:  artifactInfo.Type,
		Metadata: types.Metadata{
			OS: ptros,

			// Container image
			ImageID:     artifactInfo.ImageMetadata.ID,
			DiffIDs:     artifactInfo.ImageMetadata.DiffIDs,
			RepoTags:    artifactInfo.ImageMetadata.RepoTags,
			RepoDigests: artifactInfo.ImageMetadata.RepoDigests,
			ImageConfig: artifactInfo.ImageMetadata.ConfigFile,
		},
		Results: results,
		BOM:     artifactInfo.BOM,
	}, nil
}
