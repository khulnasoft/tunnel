# User Guide

## Discovering Plugins
You can find a list of Tunnel plugins distributed via tunnel-plugin-index [here][tunnel-plugin-index].
However, you can find plugins using the command line as well.

First, refresh your local copy of the plugin index:

```bash
$ tunnel plugin update
```

To list all plugins available, run:

```bash
$ tunnel plugin search
NAME                 DESCRIPTION                                                  MAINTAINER           OUTPUT
aqua                 A plugin for integration with Khulnasoft Security SaaS platform    aquasecurity
kubectl              A plugin scanning the images of a kubernetes resource        aquasecurity
referrer             A plugin for OCI referrers                                   aquasecurity           ✓
```

You can specify search keywords as arguments:

```bash
$ tunnel plugin search referrer

NAME                 DESCRIPTION                                                  MAINTAINER           OUTPUT
referrer             A plugin for OCI referrers                                   aquasecurity           ✓
```

It lists plugins with the keyword in the name or description.

## Installing  Plugins
Plugins can be installed with the `tunnel plugin install` command:

```bash
$ tunnel plugin install referrer
```

This command will download the plugin and install it in the plugin cache.

Tunnel adheres to the XDG specification, so the location depends on whether XDG_DATA_HOME is set.
Tunnel will now search XDG_DATA_HOME for the location of the Tunnel plugins cache.
The preference order is as follows:

- XDG_DATA_HOME if set and .tunnel/plugins exists within the XDG_DATA_HOME dir
- ~/.tunnel/plugins

Furthermore, it is possible to download plugins that are not registered in the index by specifying the URL directly or by specifying the file path.

```bash
$ tunnel plugin install github.com/khulnasoft/tunnel-plugin-kubectl
```
```bash
$ tunnel plugin install https://github.com/khulnasoft/tunnel-plugin-kubectl/archive/refs/heads/main.zip
```
```bash
$ tunnel plugin install ./myplugin.tar.gz
```

If the plugin's Git repository is [properly tagged](./developer-guide.md#tagging-plugin-repositories), you can specify the version to install like this:

```bash
$ tunnel plugin install referrer@v0.3.0
```

!!! note
    The leading `v` in the version is required. Also, the version must follow the [Semantic Versioning](https://semver.org/).

Under the hood Tunnel leverages [go-getter][go-getter] to download plugins.
This means the following protocols are supported for downloading plugins:

- OCI Registries
- Local Files
- Git
- HTTP/HTTPS
- Mercurial
- Amazon S3
- Google Cloud Storage

## Listing Installed Plugins
To list all plugins installed, run:

```bash
$ tunnel plugin list
```

## Using Plugins
Once the plugin is installed, Tunnel will load all available plugins in the cache on the start of the next Tunnel execution.
A plugin will be made in the Tunnel CLI based on the plugin name.
To display all plugins, you can list them by `tunnel --help`

```bash
$ tunnel --help
NAME:
   tunnel - A simple and comprehensive vulnerability scanner for containers

USAGE:
   tunnel [global options] command [command options] target

VERSION:
   dev

Scanning Commands
  config      Scan config files for misconfigurations
  filesystem  Scan local filesystem
  image       Scan a container image
  
...

Plugin Commands
  kubectl     scan kubectl resources
  referrer    Put referrers to OCI registry
```

As shown above, `kubectl` subcommand exists in the `Plugin Commands` section.
To call the kubectl plugin and scan existing Kubernetes deployments, you can execute the following command:

```
$ tunnel kubectl deployment <deployment-id> -- --ignore-unfixed --severity CRITICAL
```

Internally the kubectl plugin calls the kubectl binary to fetch information about that deployment and passes the using images to Tunnel.
You can see the detail [here][tunnel-plugin-kubectl].

If you want to omit even the subcommand, you can use `TUNNEL_RUN_AS_PLUGIN` environment variable.

```bash
$ TUNNEL_RUN_AS_PLUGIN=kubectl tunnel job your-job -- --format json
```

## Installing and Running Plugins on the fly
`tunnel plugin run` installs a plugin and runs it on the fly.
If the plugin is already present in the cache, the installation is skipped.

```bash
tunnel plugin run kubectl pod your-pod -- --exit-code 1
```

## Upgrading Plugins
To upgrade all plugins that you have installed to their latest versions, run:

```bash
$ tunnel plugin upgrade
```

To upgrade only certain plugins, you can explicitly specify their names:

```bash
$ tunnel plugin upgrade <PLUGIN1> <PLUGIN2>
```

## Uninstalling Plugins
Specify a plugin name with `tunnel plugin uninstall` command.

```bash
$ tunnel plugin uninstall kubectl
```

Here's the revised English documentation based on your requested changes:

## Output Mode Support
While plugins are typically intended to be used as subcommands of Tunnel, plugins supporting the output mode can be invoked as part of Tunnel's built-in commands.

!!! warning "EXPERIMENTAL"
    This feature might change without preserving backwards compatibility.

Tunnel supports plugins that are compatible with the output mode, which process Tunnel's output, such as by transforming the output format or sending it elsewhere.
You can determine whether a plugin supports the output mode by checking the `OUTPUT` column in the output of `tunnel plugin search` or `tunnel plugin list`.

```bash
$ tunnel plugin search
NAME                 DESCRIPTION                                                  MAINTAINER           OUTPUT
aqua                 A plugin for integration with Khulnasoft Security SaaS platform    aquasecurity
kubectl              A plugin scanning the images of a kubernetes resource        aquasecurity
referrer             A plugin for OCI referrers                                   aquasecurity           ✓
```

In this case, the `referrer` plugin supports the output mode.

For instance, in the case of image scanning, a plugin supporting the output mode can be called as follows:

```bash
$ tunnel image --format json --output plugin=<plugin_name> [--output-plugin-arg <plugin_flags>] <image_name>
```

Since scan results are passed to the plugin via standard input, plugins must be capable of handling standard input.

!!! warning
    To avoid Tunnel hanging, you need to read all data from `Stdin` before the plugin exits successfully or stops with an error.

While the example passes JSON to the plugin, other formats like SBOM can also be passed (e.g., `--format cyclonedx`).

If a plugin requires flags or other arguments, they can be passed using `--output-plugin-arg`.
This is directly forwarded as arguments to the plugin.
For example, `--output plugin=myplugin --output-plugin-arg "--foo --bar=baz"` translates to `myplugin --foo --bar=baz` in execution.

An example of a plugin supporting the output mode is available [here][tunnel-plugin-count].
It can be used as below:

```bash
# Install the plugin first
$ tunnel plugin install count

# Call the plugin supporting the output mode in image scanning
$ tunnel image --format json --output plugin=count --output-plugin-arg "--published-after 2023-10-01" debian:12
```

## Example

- [kubectl][tunnel-plugin-kubectl]
- [count][tunnel-plugin-count]

[tunnel-plugin-index]: https://khulnasoft.github.io/tunnel-plugin-index/
[go-getter]: https://github.com/hashicorp/go-getter
[tunnel-plugin-kubectl]: https://github.com/khulnasoft/tunnel-plugin-kubectl
[tunnel-plugin-count]: https://github.com/khulnasoft/tunnel-plugin-count
