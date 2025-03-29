package digitalocean

import (
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/digitalocean/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/digitalocean/spaces"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/digitalocean"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) digitalocean.DigitalOcean {
	return digitalocean.DigitalOcean{
		Compute: compute.Adapt(modules),
		Spaces:  spaces.Adapt(modules),
	}
}
