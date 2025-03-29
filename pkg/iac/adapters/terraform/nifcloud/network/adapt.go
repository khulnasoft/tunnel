package network

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/nifcloud/network"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) network.Network {

	return network.Network{
		ElasticLoadBalancers: adaptElasticLoadBalancers(modules),
		LoadBalancers:        adaptLoadBalancers(modules),
		Routers:              adaptRouters(modules),
		VpnGateways:          adaptVpnGateways(modules),
	}
}
