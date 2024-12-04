## tunnel clean

Remove cached files

```
tunnel clean [flags]
```

### Examples

```
  # Remove all caches
  $ tunnel clean --all

  # Remove scan cache
  $ tunnel clean --scan-cache

  # Remove vulnerability database
  $ tunnel clean --vuln-db

```

### Options

```
  -a, --all             remove all caches
      --checks-bundle   remove checks bundle
  -h, --help            help for clean
      --java-db         remove Java database
      --scan-cache      remove scan cache (container and VM image analysis results)
      --vex-repo        remove VEX repositories
      --vuln-db         remove vulnerability database
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

- [tunnel](tunnel.md) - Unified security scanner
