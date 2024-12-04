package terraform

import (
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/aws"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/azure"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/cloudstack"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/digitalocean"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/github"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/google"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/kubernetes"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/nifcloud"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/openstack"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/oracle"
	"github.com/khulnasoft/tunnel/pkg/iac/state"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) *state.State {
	return &state.State{
		AWS:          aws.Adapt(modules),
		Azure:        azure.Adapt(modules),
		CloudStack:   cloudstack.Adapt(modules),
		DigitalOcean: digitalocean.Adapt(modules),
		GitHub:       github.Adapt(modules),
		Google:       google.Adapt(modules),
		Kubernetes:   kubernetes.Adapt(modules),
		Nifcloud:     nifcloud.Adapt(modules),
		OpenStack:    openstack.Adapt(modules),
		Oracle:       oracle.Adapt(modules),
	}
}
