# Rust

Tunnel supports [Cargo](https://doc.rust-lang.org/stable/cargo/), which is the Rust package manager.
The following scanners are supported for Cargo.

| Package manager | SBOM  | Vulnerability | License |
| --------------- | :---: | :-----------: | :-----: |
| Cargo           |   ✓   |       ✓       |    -    |

In addition, it supports binaries built with [cargo-auditable](https://github.com/rust-secure-code/cargo-auditable).

| Artifact | SBOM  | Vulnerability | License |
| -------- | :---: | :-----------: | :-----: |
| Binaries |   ✓   |       ✓       |    -    |

## Features
The following table provides an outline of the features Tunnel offers.

| Package manager | File       | Transitive dependencies | Dev dependencies | [Dependency graph][dependency-graph] | Position |
|-----------------|------------|:-----------------------:|:-----------------|:------------------------------------:|:--------:|
| Cargo           | Cargo.lock |            ✓            | Excluded[^1]     |                  ✓                   |    ✓     |


| Artifact | Transitive dependencies | Dev dependencies | Dependency graph | Position |
| -------- | :---------------------: | :--------------- | :--------------: | :------: |
| Binaries |            ✓            | Excluded         |        -         |    -     |


### Cargo
Tunnel searches for `Cargo.lock` to detect dependencies.

Tunnel also supports dependency trees; however, to display an accurate tree, it needs to know whether each package is a direct dependency of the project.
Since this information is not included in `Cargo.lock`, Tunnel parses `Cargo.toml`, which should be located next to `Cargo.lock`.
If you want to see the dependency tree, please ensure that `Cargo.toml` is present.

Scan `Cargo.lock` and `Cargo.toml` together also removes developer dependencies.

### Binaries
Tunnel scans binaries built with [cargo-auditable](https://github.com/rust-secure-code/cargo-auditable).
If such a binary exists, Tunnel will identify it as being built with cargo-audit and scan it.

[^1]: When you scan Cargo.lock and Cargo.toml together.

[dependency-graph]: ../../configuration/reporting.md#show-origins-of-vulnerable-dependencies