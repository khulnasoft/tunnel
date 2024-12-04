package state

import (
	"reflect"

	"github.com/khulnasoft/tunnel/pkg/iac/providers/aws"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/cloudstack"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/digitalocean"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/github"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/google"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/kubernetes"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/nifcloud"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/openstack"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/oracle"
	"github.com/khulnasoft/tunnel/pkg/iac/rego/convert"
)

type State struct {
	AWS          aws.AWS
	Azure        azure.Azure
	CloudStack   cloudstack.CloudStack
	DigitalOcean digitalocean.DigitalOcean
	GitHub       github.GitHub
	Google       google.Google
	Kubernetes   kubernetes.Kubernetes
	OpenStack    openstack.OpenStack
	Oracle       oracle.Oracle
	Nifcloud     nifcloud.Nifcloud
}

func (a *State) ToRego() any {
	return convert.StructToRego(reflect.ValueOf(a))
}
