<div align="center">
<img src="docs/imgs/logo.png" width="200">

[![GitHub Release][release-img]][release]
[![Test][test-img]][test]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]
[![GitHub Downloads][github-downloads-img]][release]
![Docker Pulls][docker-pulls]

[📖 Documentation][docs]
</div>

Tunnel ([pronunciation][pronunciation]) is a comprehensive and versatile security scanner.
Tunnel has *scanners* that look for security issues, and *targets* where it can find those issues.

Targets (what Tunnel can scan):

- Container Image
- Filesystem
- Git Repository (remote)
- Virtual Machine Image
- Kubernetes

Scanners (what Tunnel can find there):

- OS packages and software dependencies in use (SBOM)
- Known vulnerabilities (CVEs)
- IaC issues and misconfigurations
- Sensitive information and secrets
- Software licenses

Tunnel supports most popular programming languages, operating systems, and platforms. For a complete list, see the [Scanning Coverage] page.

To learn more, go to the [Tunnel homepage][homepage] for feature highlights, or to the [Documentation site][docs] for detailed information.

## Quick Start

### Get Tunnel

Tunnel is available in most common distribution channels. The full list of installation options is available in the [Installation] page. Here are a few popular examples:

- `brew install tunnel`
- `docker run khulnasoft/tunnel`
- Download binary from <https://github.com/khulnasoft/tunnel/releases/latest/>
- See [Installation] for more

Tunnel is integrated with many popular platforms and applications. The complete list of integrations is available in the [Ecosystem] page. Here are a few popular examples:

- [GitHub Actions](https://github.com/khulnasoft/tunnel-action)
- [Kubernetes operator](https://github.com/khulnasoft/tunnel-operator)
- [VS Code plugin](https://github.com/khulnasoft/tunnel-vscode-extension)
- See [Ecosystem] for more

### Canary builds
There are canary builds ([Docker Hub](https://hub.docker.com/r/khulnasoft/tunnel/tags?page=1&name=canary), [GitHub](https://github.com/khulnasoft/tunnel/pkgs/container/tunnel/75776514?tag=canary), [ECR](https://gallery.ecr.aws/khulnasoft/tunnel#canary) images and [binaries](https://github.com/khulnasoft/tunnel/actions/workflows/canary.yaml)) as generated every push to main branch.

Please be aware: canary builds might have critical bugs, it's not recommended for use in production.

### General usage

```bash
tunnel <target> [--scanners <scanner1,scanner2>] <subject>
```

Examples:

```bash
tunnel image python:3.4-alpine
```

<details>
<summary>Result</summary>

https://user-images.githubusercontent.com/1161307/171013513-95f18734-233d-45d3-aaf5-d6aec687db0e.mov

</details>

```bash
tunnel fs --scanners vuln,secret,misconfig myproject/
```

<details>
<summary>Result</summary>

https://user-images.githubusercontent.com/1161307/171013917-b1f37810-f434-465c-b01a-22de036bd9b3.mov

</details>

```bash
tunnel k8s --report summary cluster
```

<details>
<summary>Result</summary>

![k8s summary](docs/imgs/tunnel-k8s.png)

</details>

## FAQ

### How to pronounce the name "Tunnel"?

`tri` is pronounced like **tri**gger, `vy` is pronounced like en**vy**.

## Want more? Check out Khulnasoft

If you liked Tunnel, you will love Khulnasoft which builds on top of Tunnel to provide even more enhanced capabilities for a complete security management offering.  
You can find a high level comparison table specific to Tunnel users [here](https://tunnel.dev/latest/commercial/compare/).
In addition check out the <https://khulnasoft.com> website for more information about our products and services.
If you'd like to contact Khulnasoft or request a demo, please use this form: <https://www.khulnasoft.com/demo>

## Community

Tunnel is an [Khulnasoft Security][khulnasoft] open source project.  
Learn about our open source work and portfolio [here][oss].  
Contact us about any matter by opening a GitHub Discussion [here][discussions]

Please ensure to abide by our [Code of Conduct][code-of-conduct] during all interactions.

[test]: https://github.com/khulnasoft/tunnel/actions/workflows/test.yaml
[test-img]: https://github.com/khulnasoft/tunnel/actions/workflows/test.yaml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/khulnasoft/tunnel
[go-report-img]: https://goreportcard.com/badge/github.com/khulnasoft/tunnel
[release]: https://github.com/khulnasoft/tunnel/releases
[release-img]: https://img.shields.io/github/release/khulnasoft/tunnel.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/khulnasoft/tunnel/total?logo=github
[docker-pulls]: https://img.shields.io/docker/pulls/khulnasoft/tunnel?logo=docker&label=docker%20pulls%20%2F%20tunnel
[license]: https://github.com/khulnasoft/tunnel/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[homepage]: https://tunnel.dev
[docs]: https://tunnel.dev/latest/docs/
[pronunciation]: #how-to-pronounce-the-name-tunnel
[code-of-conduct]: https://github.com/khulnasoft/community/blob/main/CODE_OF_CONDUCT.md

[Installation]:https://tunnel.dev/latest/getting-started/installation/
[Ecosystem]: https://tunnel.dev/latest/ecosystem/
[Scanning Coverage]: https://tunnel.dev/latest/docs/coverage/

[alpine]: https://ariadne.space/2021/06/08/the-vulnerability-remediation-lifecycle-of-alpine-containers/
[rego]: https://www.openpolicyagent.org/docs/latest/#rego
[sigstore]: https://www.sigstore.dev/

[khulnasoft]: https://khulnasoft.com
[oss]: https://www.khulnasoft.com/products/open-source-projects/
[discussions]: https://github.com/khulnasoft/tunnel/discussions
