package cloudstack

import (
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/cloudstack/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/cloudstack"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) cloudstack.CloudStack {
	return cloudstack.CloudStack{
		Compute: compute.Adapt(modules),
	}
}
