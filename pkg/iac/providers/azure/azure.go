package azure

import (
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/appservice"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/authorization"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/container"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/database"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/datafactory"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/datalake"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/keyvault"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/monitor"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/network"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/securitycenter"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/storage"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/azure/synapse"
)

type Azure struct {
	AppService     appservice.AppService
	Authorization  authorization.Authorization
	Compute        compute.Compute
	Container      container.Container
	Database       database.Database
	DataFactory    datafactory.DataFactory
	DataLake       datalake.DataLake
	KeyVault       keyvault.KeyVault
	Monitor        monitor.Monitor
	Network        network.Network
	SecurityCenter securitycenter.SecurityCenter
	Storage        storage.Storage
	Synapse        synapse.Synapse
}
