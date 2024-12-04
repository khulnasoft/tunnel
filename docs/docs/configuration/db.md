# Tunnel Databases

When you install Tunnel, the installed artifact contains the scanner engine but is lacking relevant security information needed to make security detections and recommendations.
These so called "databases" are automatically fetched and maintained by Tunnel as needed, so normally you shouldn't notice or worry about them.  
This document elaborates on the database management mechanism and its configuration options.

Tunnel relies on the following databases:

| DB                 | Artifact name    | Contents                                      | Purpose                                                                                    |
| ------------------ | ---------------- | --------------------------------------------- | ------------------------------------------------------------------------------------------ |
| Vulnerabilities DB | `tunnel-db`      | CVE information collected from various feeds  | used only for [vulnerability scanning](../scanner/vulnerability.md)                        |
| Java DB            | `tunnel-java-db` | Index of Java artifacts and their hash digest | used to identify Java artifacts only in [JAR scanning](../coverage/language/java.md)       |
| Checks Bundle      | `tunnel-checks`  | Logic of misconfiguration checks              | used only in [misconfiguration/IaC scanning](../scanner/misconfiguration/check/builtin.md) |

!!! note
This is not an exhaustive list of Tunnel's external connectivity requirements.
There are additional external resources which may be required by specific Tunnel features.
To learn about external connectivity requirements, see the [Advanced Network Scenarios](../advanced/air-gap.md).

## Locations

Tunnel's databases are published to the following locations:

| Registry   | Image Address                               | Link                                                 |
| ---------- | ------------------------------------------- | ---------------------------------------------------- |
| GHCR       | `ghcr.io/khulnasoft/tunnel-db`              | <https://ghcr.io/khulnasoft/tunnel-db>               |
|            | `ghcr.io/khulnasoft/tunnel-java-db`         | <https://ghcr.io/khulnasoft/tunnel-java-db>          |
|            | `ghcr.io/khulnasoft/tunnel-checks`          | <https://ghcr.io/khulnasoft/tunnel-checks>           |
| Docker Hub | `khulnasoft/tunnel-db`                      | <https://hub.docker.com/r/khulnasoft/tunnel-db>      |
|            | `khulnasoft/tunnel-java-db`                 | <https://hub.docker.com/r/khulnasoft/tunnel-java-db> |
|            | `khulnasoft/tunnel-checks`                  | <https://hub.docker.com/r/khulnasoft/tunnel-checks>  |
| AWS ECR    | `public.ecr.aws/aquasecurity/tunnel-db`     | <https://gallery.ecr.aws/aquasecurity/tunnel-db>     |
|            | `public.ecr.aws/aquasecurity/trivy-java-db` | <https://gallery.ecr.aws/aquasecurity/trivy-java-db> |
|            | `public.ecr.aws/aquasecurity/trivy-checks`  | <https://gallery.ecr.aws/aquasecurity/trivy-checks>  |

In addition, images are also available via pull-through cache registries like [Google Container Registry Mirror](https://cloud.google.com/artifact-registry/docs/pull-cached-dockerhub-images).

## Default Locations

Tunnel will attempt to pull images from the following registries in the order specified.

1. `ghcr.io/khulnasoft`
2. `ghcr.io/khulnasoft`

You can specify additional alternative repositories as explained in the [configuring database locations section](#database-locations).

## DB Management Configuration

### Database Locations

You can configure Tunnel to download databases from alternative locations by using the flags:

- `--db-repository`
- `--java-db-repository`
- `--checks-bundle-repository`

The value should be an image address in a container registry.

For example:

```
tunnel image --db-repository registry.gitlab.com/gitlab-org/security-products/dependencies/tunnel-db alpine
```

The flags accepts multiple values, which can be used to specify multiple alternative repository locations. In case of a transient errors (e.g. status 429 or 5xx), Tunnel will fall back to alternative registries in the order specified.

For example:

```
tunnel image --db-repository my.registry.local/tunnel-db --db-repository registry.gitlab.com/gitlab-org/security-products/dependencies/tunnel-db alpine
```

The Checks Bundle registry location option does not support fallback through multiple options. This is because in case of a failure pulling the Checks Bundle, Tunnel will use the embedded checks as a fallback.

!!! note
Setting the repository location flags override the default values which include the official db locations. In case you want to preserve the default locations, you should include them in the list the you set as repository locations.

!!!note
When pulling `tunnel-db` or `tunnel-java-db`, if image tag is not specified, Tunnel defaults to the db schema number instead of the `latest` tag.

### Skip updates

You can configure Tunnel to not attempt to download any or all database(s), using the flags:

- `--skip-db-update`
- `--skip-java-db-update`
- `--skip-check-update`

For example:

```
tunnel image --skip-db-update --skip-java-db-update --skip-check-update alpine
```

### Only update

You can ask `Tunnel` to only update the database without performing a scan. This action will ensure Tunnel is up to date, and populate Tunnel's database cache for subsequent scans.

- `--download-db-only`
- `--download-java-db-only`

For example:

```
tunnel image --download-db-only
```

Note that currently there is no option to download only the Checks Bundle.

### Remove Databases

`tunnel clean` command removes caches and databases.
You can select which cache component to remove:

| option            | description                                                 |
| ----------------- | ----------------------------------------------------------- |
| `-a`/`--all`      | remove all caches                                           |
| `--checks-bundle` | remove checks bundle                                        |
| `--java-db`       | remove Java database                                        |
| `--scan-cache`    | remove scan cache (container and VM image analysis results) |
| `--vuln-db`       | remove vulnerability database                               |

Example:

```
$ tunnel clean --vuln-db --java-db
2024-06-24T11:42:31+06:00       INFO    Removing vulnerability database...
2024-06-24T11:42:31+06:00       INFO    Removing Java database...
```
