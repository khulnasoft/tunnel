package compute

import (
	"github.com/khulnasoft/tunnel/pkg/iac/types"
)

type Network struct {
	Metadata    types.Metadata
	Firewall    *Firewall
	Subnetworks []SubNetwork
}
