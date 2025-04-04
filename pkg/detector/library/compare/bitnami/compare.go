package bitnami

import (
	version "github.com/bitnami/go-version/pkg/version"
	"golang.org/x/xerrors"

	dbTypes "github.com/khulnasoft-lab/tunnel-db/pkg/types"
	"github.com/khulnasoft/tunnel/pkg/detector/library/compare"
)

// Comparer represents a comparer for Bitnami
type Comparer struct{}

// IsVulnerable checks if the package version is vulnerable to the advisory.
func (n Comparer) IsVulnerable(ver string, advisory dbTypes.Advisory) bool {
	return compare.IsVulnerable(ver, advisory, n.matchVersion)
}

// matchVersion checks if the package version satisfies the given constraint.
func (n Comparer) matchVersion(currentVersion, constraint string) (bool, error) {
	v, err := version.Parse(currentVersion)
	if err != nil {
		return false, xerrors.Errorf("bitnami version error (%s): %s", currentVersion, err)
	}

	c, err := version.NewConstraints(constraint)
	if err != nil {
		return false, xerrors.Errorf("bitnami constraint error (%s): %s", constraint, err)
	}

	return c.Check(v), nil
}
