package ospkg

import (
	"context"
	"time"

	"github.com/samber/lo"
	"golang.org/x/xerrors"

	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/alma"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/alpine"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/amazon"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/azure"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/chainguard"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/debian"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/oracle"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/photon"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/redhat"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/rocky"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/suse"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/ubuntu"
	"github.com/khulnasoft/tunnel/pkg/detector/ospkg/wolfi"
	ftypes "github.com/khulnasoft/tunnel/pkg/fanal/types"
	"github.com/khulnasoft/tunnel/pkg/log"
	"github.com/khulnasoft/tunnel/pkg/types"
)

var (
	// ErrUnsupportedOS defines error for unsupported OS
	ErrUnsupportedOS = xerrors.New("unsupported os")

	drivers = map[ftypes.OSType]Driver{
		ftypes.Alpine:             alpine.NewScanner(),
		ftypes.Alma:               alma.NewScanner(),
		ftypes.Amazon:             amazon.NewScanner(),
		ftypes.Azure:              azure.NewAzureScanner(),
		ftypes.CBLMariner:         azure.NewMarinerScanner(),
		ftypes.Debian:             debian.NewScanner(),
		ftypes.Ubuntu:             ubuntu.NewScanner(),
		ftypes.RedHat:             redhat.NewScanner(),
		ftypes.CentOS:             redhat.NewScanner(),
		ftypes.Rocky:              rocky.NewScanner(),
		ftypes.Oracle:             oracle.NewScanner(),
		ftypes.OpenSUSETumbleweed: suse.NewScanner(suse.OpenSUSETumbleweed),
		ftypes.OpenSUSELeap:       suse.NewScanner(suse.OpenSUSE),
		ftypes.SLES:               suse.NewScanner(suse.SUSEEnterpriseLinux),
		ftypes.SLEMicro:           suse.NewScanner(suse.SUSEEnterpriseLinuxMicro),
		ftypes.Photon:             photon.NewScanner(),
		ftypes.Wolfi:              wolfi.NewScanner(),
		ftypes.Chainguard:         chainguard.NewScanner(),
	}
)

// RegisterDriver is defined for extensibility and not supposed to be used in Tunnel.
func RegisterDriver(name ftypes.OSType, driver Driver) {
	drivers[name] = driver
}

// Driver defines operations for OS package scan
type Driver interface {
	Detect(context.Context, string, *ftypes.Repository, []ftypes.Package) ([]types.DetectedVulnerability, error)
	IsSupportedVersion(context.Context, ftypes.OSType, string) bool
}

// Detect detects the vulnerabilities
func Detect(ctx context.Context, _, osFamily ftypes.OSType, osName string, repo *ftypes.Repository, _ time.Time, pkgs []ftypes.Package) ([]types.DetectedVulnerability, bool, error) {
	ctx = log.WithContextPrefix(ctx, string(osFamily))

	driver, err := newDriver(osFamily)
	if err != nil {
		return nil, false, ErrUnsupportedOS
	}

	eosl := !driver.IsSupportedVersion(ctx, osFamily, osName)

	// Package `gpg-pubkey` doesn't use the correct version.
	// We don't need to find vulnerabilities for this package.
	filteredPkgs := lo.Filter(pkgs, func(pkg ftypes.Package, index int) bool {
		return pkg.Name != "gpg-pubkey"
	})
	vulns, err := driver.Detect(ctx, osName, repo, filteredPkgs)
	if err != nil {
		return nil, false, xerrors.Errorf("failed detection: %w", err)
	}

	return vulns, eosl, nil
}

func newDriver(osFamily ftypes.OSType) (Driver, error) {
	if driver, ok := drivers[osFamily]; ok {
		return driver, nil
	}

	log.Warn("Unsupported os", log.String("family", string(osFamily)))
	return nil, ErrUnsupportedOS
}
