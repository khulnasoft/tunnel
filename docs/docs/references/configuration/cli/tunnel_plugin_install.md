## tunnel plugin install

Install a plugin

```
tunnel plugin install NAME | URL | FILE_PATH
```

### Examples

```
  # Install a plugin from the plugin index
  $ tunnel plugin install referrer

  # Specify the version of the plugin to install
  $ tunnel plugin install referrer@v0.3.0

  # Install a plugin from a URL
  $ tunnel plugin install github.com/khulnasoft/tunnel-plugin-referrer
```

### Options

```
  -h, --help   help for install
```

### Options inherited from parent commands

```
      --cache-dir string          cache directory (default "/path/to/cache")
  -c, --config string             config path (default "tunnel.yaml")
  -d, --debug                     debug mode
      --generate-default-config   write the default config to tunnel-default.yaml
      --insecure                  allow insecure server connections
  -q, --quiet                     suppress progress bar and log output
      --timeout duration          timeout (default 5m0s)
  -v, --version                   show version
```

### SEE ALSO

- [tunnel plugin](tunnel_plugin.md) - Manage plugins
