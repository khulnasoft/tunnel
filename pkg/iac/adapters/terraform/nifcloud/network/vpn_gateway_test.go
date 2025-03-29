package network

import (
	"testing"

	"github.com/khulnasoft/tunnel/internal/testutil"
	"github.com/khulnasoft/tunnel/pkg/iac/adapters/terraform/tftestutil"
	"github.com/khulnasoft/tunnel/pkg/iac/providers/nifcloud/network"
	iacTypes "github.com/khulnasoft/tunnel/pkg/iac/types"
)

func Test_adaptVpnGateways(t *testing.T) {
	tests := []struct {
		name      string
		terraform string
		expected  []network.VpnGateway
	}{
		{
			name: "configured",
			terraform: `
			resource "nifcloud_vpn_gateway" "example" {
				security_group  = "example-security-group"
			}
`,
			expected: []network.VpnGateway{{
				Metadata:      iacTypes.NewTestMetadata(),
				SecurityGroup: iacTypes.String("example-security-group", iacTypes.NewTestMetadata()),
			}},
		},
		{
			name: "defaults",
			terraform: `
			resource "nifcloud_vpn_gateway" "example" {
			}
`,

			expected: []network.VpnGateway{{
				Metadata:      iacTypes.NewTestMetadata(),
				SecurityGroup: iacTypes.String("", iacTypes.NewTestMetadata()),
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modules := tftestutil.CreateModulesFromSource(t, test.terraform, ".tf")
			adapted := adaptVpnGateways(modules)
			testutil.AssertDefsecEqual(t, test.expected, adapted)
		})
	}
}
