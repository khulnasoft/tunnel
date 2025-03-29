package digitalocean

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/digitalocean/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/digitalocean/spaces"
)

type DigitalOcean struct {
	Compute compute.Compute
	Spaces  spaces.Spaces
}
