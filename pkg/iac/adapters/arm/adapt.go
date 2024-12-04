package arm

import (
	"context"

	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/appservice"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/authorization"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/container"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/database"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/datafactory"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/datalake"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/keyvault"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/monitor"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/network"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/securitycenter"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/storage"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/arm/synapse"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure"
	scanner "github.com/khulnasoft/tunnel/pkg/iac/scanners/azure"
	"github.com/khulnasoft/tunnel/pkg/iac/state"
)

// Adapt adapts an azure arm instance
func Adapt(ctx context.Context, deployment scanner.Deployment) *state.State {
	return &state.State{
		Azure: adaptAzure(deployment),
	}
}

func adaptAzure(deployment scanner.Deployment) azure.Azure {

	return azure.Azure{
		AppService:     appservice.Adapt(deployment),
		Authorization:  authorization.Adapt(deployment),
		Compute:        compute.Adapt(deployment),
		Container:      container.Adapt(deployment),
		Database:       database.Adapt(deployment),
		DataFactory:    datafactory.Adapt(deployment),
		DataLake:       datalake.Adapt(deployment),
		KeyVault:       keyvault.Adapt(deployment),
		Monitor:        monitor.Adapt(deployment),
		Network:        network.Adapt(deployment),
		SecurityCenter: securitycenter.Adapt(deployment),
		Storage:        storage.Adapt(deployment),
		Synapse:        synapse.Adapt(deployment),
	}

}
