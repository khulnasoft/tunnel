## tunnel module uninstall

Uninstall a module

```
tunnel module uninstall [flags] REPOSITORY
```

### Options

```
  -h, --help   help for uninstall
```

### Options inherited from parent commands

```
      --cache-dir string          cache directory (default "/path/to/cache")
  -c, --config string             config path (default "tunnel.yaml")
  -d, --debug                     debug mode
      --enable-modules strings    [EXPERIMENTAL] module names to enable
      --generate-default-config   write the default config to tunnel-default.yaml
      --insecure                  allow insecure server connections
      --module-dir string         specify directory to the wasm modules that will be loaded (default "$HOME/.tunnel/modules")
  -q, --quiet                     suppress progress bar and log output
      --timeout duration          timeout (default 5m0s)
  -v, --version                   show version
```

### SEE ALSO

- [tunnel module](tunnel_module.md) - Manage modules
