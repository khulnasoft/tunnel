# Developer Guide

## Developing Tunnel plugins

This section will guide you through the process of developing Tunnel plugins.
To help you get started quickly, we have published a [plugin template repository][plugin-template].
You can use this template as a starting point for your plugin development.

### Introduction

If you are looking to start developing plugins for Tunnel, read [the user guide](./user-guide.md) first.

The development process involves the following steps:

- Create a repository for your plugin, named `tunnel-plugin-<name>`.
- Create an executable binary that can be invoked as `tunnel <name>`.
- Place the executable binary in a repository.
- Create a `plugin.yaml` file that describes the plugin.
- (Submit your plugin to the [Tunnel plugin index][tunnel-plugin-index].)

After you develop a plugin with a good name following the best practices and publish it, you can submit your plugin to the [Tunnel plugin index][tunnel-plugin-index].

### Naming

This section describes guidelines for naming your plugins.

#### Use `tunnel-plugin-` prefix

The name of the plugin repository should be prefixed with `tunnel-plugin-`.

#### Use lowercase and hyphens

Plugin names must be all lowercase and separate words with hyphens.
Don’t use camelCase, PascalCase, or snake_case; use kebab-case.

- NO: `tunnel OpenSvc`
- YES: `tunnel open-svc`

#### Be specific

Plugin names should not be verbs or nouns that are generic, already overloaded, or likely to be used for broader purposes by another plugin.

- NO: tunnel sast (Too broad)
- YES: tunnel govulncheck

#### Be unique

Find a unique name for your plugin that differentiates it from other plugins that perform a similar function.

- NO: `tunnel images` (Unclear how it is different from the builtin “image" command)
- YES: `tunnel registry-images` (Unique name).

#### Prefix Vendor Identifiers

Use vendor-specific strings as prefix, separated with a dash.
This makes it easier to search/group plugins that are about a specific vendor.

- NO: `tunnel security-hub-aws (Makes it harder to search or locate in a plugin list)
- YES: `tunnel aws-security-hub (Will show up together with other aws-\* plugins)

### Choosing a language

Since Tunnel plugins are standalone executables, you can write them in any programming language.

If you are planning to write a plugin with Go, check out [the Report struct](https://github.com/khulnasoft/tunnel/blob/787b466e069e2d04e73b3eddbda621e5eec8543b/pkg/types/report.go#L13-L24),
which is the output of Tunnel scan.

### Writing your plugin

Each plugin has a top-level directory, and then a `plugin.yaml` file.

```bash
your-plugin/
  |
  |- plugin.yaml
  |- your-plugin.sh
```

In the example above, the plugin is contained inside a directory named `your-plugin`.
It has two files: `plugin.yaml` (required) and an executable script, `your-plugin.sh` (optional).

#### Writing a plugin manifest

The plugin manifest is a simple YAML file named `plugin.yaml`.
Here is an example YAML of [tunnel-plugin-kubectl][tunnel-plugin-kubectl] plugin that adds support for Kubernetes scanning.

```yaml
name: "kubectl"
version: "0.1.0"
repository: github.com/khulnasoft/tunnel-plugin-kubectl
maintainer: khulnasoft
output: false
summary: Scan kubectl resources
description: |-
  A Tunnel plugin that scans the images of a kubernetes resource.
  Usage: tunnel kubectl TYPE[.VERSION][.GROUP] NAME
platforms:
  - selector: # optional
      os: darwin
      arch: amd64
    uri: ./tunnel-kubectl # where the execution file is (local file, http, git, etc.)
    bin: ./tunnel-kubectl # path to the execution file
  - selector: # optional
      os: linux
      arch: amd64
    uri: https://github.com/khulnasoft/tunnel-plugin-kubectl/releases/download/v0.1.0/tunnel-kubectl.tar.gz
    bin: ./tunnel-kubectl
```

We encourage you to copy and adapt plugin manifests of existing plugins.

- [count][tunnel-plugin-count]
- [referrer][tunnel-plugin-referrer]

The `plugin.yaml` field should contain the following information:

- name: The name of the plugin. This also determines how the plugin will be made available in the Tunnel CLI. For example, if the plugin is named kubectl, you can call the plugin with `tunnel kubectl`. (required)
- version: The version of the plugin. [Semantic Versioning][semver] should be used. (required)
- repository: The repository name where the plugin is hosted. (required)
- maintainer: The name of the maintainer of the plugin. (required)
- output: Whether the plugin supports [the output mode](./user-guide.md#output-mode-support). (optional)
- usage: Deprecated: use summary instead. (optional)
- summary: A short usage description. (required)
- description: A long description of the plugin. This is where you could provide a helpful documentation of your plugin. (required)
- platforms: (required)
  - selector: The OS/Architecture specific variations of a execution file. (optional)
    - os: OS information based on GOOS (linux, darwin, etc.) (optional)
    - arch: The architecture information based on GOARCH (amd64, arm64, etc.) (optional)
  - uri: Where the executable file is. Relative path from the root directory of the plugin or remote URL such as HTTP and S3. (required)
  - bin: Which file to call when the plugin is executed. Relative path from the root directory of the plugin. (required)

The following rules will apply in deciding which platform to select:

- If both `os` and `arch` under `selector` match the current platform, search will stop and the platform will be used.
- If `selector` is not present, the platform will be used.
- If `os` matches and there is no more specific `arch` match, the platform will be used.
- If no `platform` match is found, Tunnel will exit with an error.

After determining platform, Tunnel will download the execution file from `uri` and store it in the plugin cache.
When the plugin is called via Tunnel CLI, `bin` command will be executed.

#### Tagging plugin repositories

If you are hosting your plugin in a Git repository, it is strongly recommended to tag your releases with a version number.
By tagging your releases, Tunnel can install specific versions of your plugin.

```bash
$ tunnel plugin install referrer@v0.3.0
```

When tagging versions, you must follow [the Semantic Versioning][semver] and prefix the tag with `v`, like `v1.2.3`.

#### Plugin arguments/flags

The plugin is responsible for handling flags and arguments.
Any arguments are passed to the plugin from the `tunnel` command.

#### Testing plugin installation locally

A plugin should be archived `*.tar.gz`.
After you have archived your plugin into a `.tar.gz` file, you can verify that your plugin installs correctly with Tunnel.

```bash
$ tar -czvf myplugin.tar.gz plugin.yaml script.py
plugin.yaml
script.py

$ tunnel plugin install myplugin.tar.gz
2023-03-03T19:04:42.026+0600	INFO	Installing the plugin from myplugin.tar.gz...
2023-03-03T19:04:42.026+0600	INFO	Loading the plugin metadata...

$ tunnel myplugin
Hello from Tunnel demo plugin!
```

## Publishing plugins

The [plugin.yaml](#writing-a-plugin-manifest) file is the core of your plugin, so as long as it is published somewhere, your plugin can be installed.
If you choose to publish your plugin on GitHub, you can make it installable by placing the plugin.yaml file in the root directory of your repository.
Users can then install your plugin with the command, `tunnel plugin install github.com/org/repo`.

While the `uri` specified in the plugin.yaml file doesn't necessarily need to point to the same repository, it's a good practice to host the executable file within the same repository when using GitHub.
You can utilize GitHub Releases to distribute the executable file.
For an example of how to structure your plugin repository, refer to [the plugin template repository][plugin-template].

## Distributing plugins via the Tunnel plugin index

Tunnel can install plugins directly by specifying a repository, like `tunnel plugin install github.com/khulnasoft/tunnel-plugin-referrer`,
so you don't necessarily need to register your plugin in the Tunnel plugin index.
However, we would recommend distributing your plugin via the Tunnel plugin index
since it makes it easier for other users to find (`tunnel plugin search`) and install your plugin (e.g. `tunnel plugin install kubectl`).

### Pre-submit checklist

- Review [the plugin naming guide](#naming).
- Ensure the `plugin.yaml` file has all the required fields.
- Tag a git release with a semantic version (e.g. v1.0.0).
- [Test your plugin installation locally](#testing-plugin-installation-locally).

### Submitting plugins

Submitting your plugin to the plugin index is a straightforward process.
All you need to do is create a YAML file for your plugin and place it in the [plugins/](https://github.com/khulnasoft/tunnel-plugin-index/tree/main/plugins) directory of [the index repository][tunnel-plugin-index].

Once you've done that, create a pull request (PR) and have it reviewed by the maintainers.
Once your PR is merged, the index will be updated, and your plugin will be available for installation.
[The plugin index page][plugin-list] will also be automatically updated to list your newly added plugin.

The content of the YAML file is very simple.
You only need to specify the name of your plugin and the repository where it is distributed.

```yaml
name: referrer
repository: github.com/khulnasoft/tunnel-plugin-referrer
```

After your PR is merged, the CI system will automatically retrieve the `plugin.yaml` file from your repository and update [the index.yaml file][index].
If any required fields are missing from your `plugin.yaml`, the CI will fail, so make sure your `plugin.yaml` has all the required fields before creating a PR.
Once [the index.yaml][index] has been updated, running `tunnel plugin update` will download the updated index to your local machine.

[plugin-template]: https://github.com/khulnasoft/tunnel-plugin-template
[plugin-list]: https://khulnasoft.github.io/tunnel-plugin-index/
[index]: https://khulnasoft.github.io/tunnel-plugin-index/v1/index.yaml
[semver]: https://semver.org/
[tunnel-plugin-index]: https://github.com/khulnasoft/tunnel-plugin-index
[tunnel-plugin-kubectl]: https://github.com/khulnasoft/tunnel-plugin-kubectl
[tunnel-plugin-count]: https://github.com/khulnasoft/tunnel-plugin-count/blob/main/plugin.yaml
[tunnel-plugin-referrer]: https://github.com/khulnasoft/tunnel-plugin-referrer/blob/main/plugin.yaml
