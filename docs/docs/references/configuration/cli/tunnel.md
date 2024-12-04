## tunnel

Unified security scanner

### Synopsis

Scanner for vulnerabilities in container images, file systems, and Git repositories, as well as for configuration issues and hard-coded secrets

```
tunnel [global flags] command [flags] target
```

### Examples

```
  # Scan a container image
  $ tunnel image python:3.4-alpine

  # Scan a container image from a tar archive
  $ tunnel image --input ruby-3.1.tar

  # Scan local filesystem
  $ tunnel fs .

  # Run in server mode
  $ tunnel server
```

### Options

```
      --cache-dir string          cache directory (default "/path/to/cache")
  -c, --config string             config path (default "tunnel.yaml")
  -d, --debug                     debug mode
  -f, --format string             version format (json)
      --generate-default-config   write the default config to tunnel-default.yaml
  -h, --help                      help for tunnel
      --insecure                  allow insecure server connections
  -q, --quiet                     suppress progress bar and log output
      --timeout duration          timeout (default 5m0s)
  -v, --version                   show version
```

### SEE ALSO

- [tunnel clean](tunnel_clean.md) - Remove cached files
- [tunnel config](tunnel_config.md) - Scan config files for misconfigurations
- [tunnel convert](tunnel_convert.md) - Convert Tunnel JSON report into a different format
- [tunnel filesystem](tunnel_filesystem.md) - Scan local filesystem
- [tunnel image](tunnel_image.md) - Scan a container image
- [tunnel kubernetes](tunnel_kubernetes.md) - [EXPERIMENTAL] Scan kubernetes cluster
- [tunnel module](tunnel_module.md) - Manage modules
- [tunnel plugin](tunnel_plugin.md) - Manage plugins
- [tunnel registry](tunnel_registry.md) - Manage registry authentication
- [tunnel repository](tunnel_repository.md) - Scan a repository
- [tunnel rootfs](tunnel_rootfs.md) - Scan rootfs
- [tunnel sbom](tunnel_sbom.md) - Scan SBOM for vulnerabilities and licenses
- [tunnel server](tunnel_server.md) - Server mode
- [tunnel version](tunnel_version.md) - Print the version
- [tunnel vex](tunnel_vex.md) - [EXPERIMENTAL] VEX utilities
- [tunnel vm](tunnel_vm.md) - [EXPERIMENTAL] Scan a virtual machine image
