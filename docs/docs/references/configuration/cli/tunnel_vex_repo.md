## tunnel vex repo

Manage VEX repositories

### Examples

```
  # Initialize the configuration file
  $ tunnel vex repo init

  # List VEX repositories
  $ tunnel vex repo list

  # Download the VEX repositories
  $ tunnel vex repo download

```

### Options

```
  -h, --help   help for repo
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

* [tunnel vex](tunnel_vex.md)	 - [EXPERIMENTAL] VEX utilities
* [tunnel vex repo download](tunnel_vex_repo_download.md)	 - Download the VEX repositories
* [tunnel vex repo init](tunnel_vex_repo_init.md)	 - Initialize a configuration file
* [tunnel vex repo list](tunnel_vex_repo_list.md)	 - List VEX repositories

