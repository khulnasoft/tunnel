## tunnel registry logout

Log out of a registry

```
tunnel registry logout SERVER [flags]
```

### Examples

```
  # Log out of reg.example.com
  tunnel registry logout reg.example.com
```

### Options

```
  -h, --help   help for logout
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

* [tunnel registry](tunnel_registry.md)	 - Manage registry authentication

