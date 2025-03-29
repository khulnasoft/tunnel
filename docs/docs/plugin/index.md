# Plugins
Tunnel provides a plugin feature to allow others to extend the Tunnel CLI without the need to change the Tunnel code base.
This plugin system was inspired by the plugin system used in [kubectl][kubectl], [Helm][helm], and [Conftest][conftest].

## Overview
Tunnel plugins are add-on tools that integrate seamlessly with Tunnel.
They provide a way to extend the core feature set of Tunnel, but without requiring every new feature to be written in Go and added to the core tool.

- They can be added and removed from a Tunnel installation without impacting the core Tunnel tool.
- They can be written in any programming language.
- They integrate with Tunnel, and will show up in Tunnel help and subcommands.

!!! warning
    Tunnel plugins available in public are not audited for security.
    You should install and run third-party plugins at your own risk, since they are arbitrary programs running on your machine.

## Quickstart
Tunnel helps you discover and install plugins on your machine.

You can install and use a wide variety of Tunnel plugins to enhance your experience.

Let’s get started:

1. Download the plugin list:

    ```bash
    $ tunnel plugin update
    ```

2. Discover Tunnel plugins available on the plugin index:

    ```bash
    $ tunnel plugin search
    NAME                 DESCRIPTION                                                  MAINTAINER           OUTPUT
    khulnasoft                 A plugin for integration with Khulnasoft Security SaaS platform    khulnasoft
    kubectl              A plugin scanning the images of a kubernetes resource        khulnasoft
    referrer             A plugin for OCI referrers                                   khulnasoft           ✓
    [...]
    ```

3. Choose a plugin from the list and install it:

    ```bash
    $ tunnel plugin install referrer
    ```

4. Use the installed plugin:

    ```bash
    $ tunnel referrer --help
    ```

5. Keep your plugins up-to-date:

    ```bash
    $ tunnel plugin upgrade
    ```

6. Uninstall a plugin you no longer use:

    ```bash
    tunnel plugin uninstall referrer
    ``` 

This is practically all you need to know to start using Tunnel plugins.


[kubectl]: https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/
[helm]: https://helm.sh/docs/topics/plugins/
[conftest]: https://www.conftest.dev/plugins/
