package container

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/container"
	"github.com/khulnasoft/tunnel/pkg/iac/scanners/azure"
)

func Adapt(deployment azure.Deployment) container.Container {
	return container.Container{
		KubernetesClusters: adaptKubernetesClusters(deployment),
	}
}

func adaptKubernetesClusters(deployment azure.Deployment) []container.KubernetesCluster {

	return nil
}
