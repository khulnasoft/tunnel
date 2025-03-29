# Julia

## Features

Tunnel supports [Pkg.jl](https://pkgdocs.julialang.org/v1/), which is the Julia package manager.
The following scanners are supported.

| Package manager | SBOM | Vulnerability | License |
|-----------------|:----:|:-------------:|:-------:|
| Pkg.jl          |  ✓   |       -       |    -    |

The following table provides an outline of the features Tunnel offers.

| Package manager | File          | Transitive dependencies | Dev dependencies | License | Dependency graph | Position |
| --------------- | ------------- | :---------------------: | :--------------- | :-----: | :--------------: | :------: |
| Pkg.jl          | Manifest.toml |            ✅            | Excluded[^1]     |    -    |        ✅         |    ✅     |

### Pkg.jl

Tunnel searches for `Manifest.toml` to detect dependencies.

Tunnel also supports dependency trees; however, to display an accurate tree, it needs to know whether each package is a direct dependency of the project.
Since this information is not included in `Manifest.toml`, Tunnel parses `Project.toml`, which should be located next to `Project.toml`.
If you want to see the dependency tree, please ensure that `Project.toml` is present.

Scanning `Manifest.toml` and `Project.toml` together also removes developer dependencies.

Dependency extensions are currently ignored.

[^1]: When you scan `Manifest.toml` and `Project.toml` together.
