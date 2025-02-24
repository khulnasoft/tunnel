# Built-in Checks

## Checks Sources

Tunnel has an extensive library of misconfiguration checks that is maintained at <https://github.com/khulnasoft/tunnel-checks>.  
Tunnel checks are mainly written in [Rego][rego], while some checks are written in Go.  
See [here](../../../coverage/iac/index.md) for the list of supported config types.

## Checks Bundle

When performing a misconfiguration scan, Tunnel will automatically download the relevant Checks bundle. The bundle is cached locally and Tunnel will reuse it for subsequent scans on the same machine. Tunnel takes care of updating the cache automatically, so normally users can be oblivious to it.

## Checks Distribution

Tunnel checks are distributed as an [OPA bundle][opa-bundle] hosted in the following GitHub Container Registry: <https://ghcr.io/khulnasoft-lab/tunnel-audit>.  
Tunnel checks for updates to OPA bundle on GHCR every 24 hours and pulls it if there are any updates.

### External connectivity

Tunnel needs to connect to the internet to download the bundle. If you are running Tunnel in an air-gapped environment, or an tightly controlled network, please refer to the [Advanced Network Scenarios document](../../../advanced/air-gap.md).  
The Checks bundle is also embedded in the Tunnel binary (at build time), and will be used as a fallback if Tunnel is unable to download the bundle. This means that you can still scan for misconfigurations in an air-gapped environment using the Checks from the time of the Tunnel release you are using.

[rego]: https://www.openpolicyagent.org/docs/latest/policy-language/
[opa-bundle]: https://www.openpolicyagent.org/docs/latest/management-bundles/
