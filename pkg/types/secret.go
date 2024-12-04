package types

import (
	ftypes "github.com/khulnasoft/tunnel/pkg/fanal/types"
)

type DetectedSecret ftypes.SecretFinding

func (DetectedSecret) findingType() FindingType { return FindingTypeSecret }
