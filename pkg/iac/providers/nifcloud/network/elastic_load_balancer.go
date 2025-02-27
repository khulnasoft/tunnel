package network

import (
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

type ElasticLoadBalancer struct {
	Metadata          iacTypes.Metadata
	NetworkInterfaces []NetworkInterface
	Listeners         []ElasticLoadBalancerListener
}

type ElasticLoadBalancerListener struct {
	Metadata iacTypes.Metadata
	Protocol iacTypes.StringValue
}
