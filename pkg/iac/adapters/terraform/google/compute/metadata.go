package compute

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/khulnasoft/tunnel/pkg/iac/providers/google/compute"
	"github.com/khulnasoft/tunnel/pkg/iac/terraform"
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

func adaptProjectMetadata(modules terraform.Modules) compute.ProjectMetadata {
	metadata := compute.ProjectMetadata{
		Metadata: iacTypes.NewUnmanagedMetadata(),
		EnableOSLogin: iacTypes.BoolUnresolvable(
			iacTypes.NewUnmanagedMetadata(),
		),
	}
	for _, metadataBlock := range modules.GetResourcesByType("google_compute_project_metadata") {
		metadata.Metadata = metadataBlock.GetMetadata()
		if metadataAttr := metadataBlock.GetAttribute("metadata"); metadataAttr.IsNotNil() {
			if val := metadataAttr.MapValue("enable-oslogin"); val.Type() == cty.Bool {
				metadata.EnableOSLogin = iacTypes.BoolExplicit(val.True(), metadataAttr.GetMetadata())
			}
		}
	}
	return metadata
}
