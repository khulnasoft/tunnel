# CentOS
Tunnel supports the following scanners for OS packages.

|    Scanner    | Supported |
| :-----------: | :-------: |
|     SBOM      |     ✓     |
| Vulnerability |     ✓     |
|    License    |     ✓     |

Please see [here](index.md#supported-os) for supported versions.

The table below outlines the features offered by Tunnel.

|                Feature                | Supported |
| :-----------------------------------: | :-------: |
|        Unfixed vulnerabilities        |     ✓     |
| [Dependency graph][dependency-graph] |     ✓     |

## SBOM
Same as [RHEL](rhel.md#sbom).

## Vulnerability
CentOS does not provide straightforward machine-readable security advisories.
As a result, Tunnel utilizes the security advisories from [Red Hat Enterprise Linux (RHEL)](rhel.md#vulnerability) for detecting vulnerabilities in CentOS.
This approach might lead to situations where, even though Tunnel displays a fixed version, CentOS might not have the patch available yet.
Since patches released for RHEL often become available in CentOS after some time, it's usually just a matter of waiting.

!!! note
    The case for CentOS Stream, which is not supported by Tunnel, is entirely different from CentOS.

As Tunnel relies on Red Hat's advisories, please refer to [Red Hat](rhel.md) for details regarding vulnerability severity and status.


## License
Same as [RHEL](rhel.md#license).


[dependency-graph]: ../../configuration/reporting.md#show-origins-of-vulnerable-dependencies