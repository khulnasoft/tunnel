Tunnel can download images from a private registry without the need for installing Docker or any other 3rd party tools.
This makes it easy to run within a CI process.

## Login

You can log in to a private registry using the `tunnel registry login` command.
It uses the Docker configuration file (`~/.docker/config.json`) to store the credentials under the hood, and the configuration file path can be configured by `DOCKER_CONFIG` environment variable.

```shell
$ cat ~/my_password.txt | tunnel registry login --username foo --password-stdin ghcr.io
$ tunnel image ghcr.io/your/private_image
```

## Passing Credentials

You can also provide your credentials when scanning.

```shell
$ TUNNEL_USERNAME=YOUR_USERNAME TUNNEL_PASSWORD=YOUR_PASSWORD tunnel image YOUR_PRIVATE_IMAGE
```

!!! warning
When passing credentials via environment variables or CLI flags, Tunnel will attempt to use these credentials for all registries encountered during scanning, regardless of the target registry.
This can potentially lead to unintended credential exposure.
To mitigate this risk:

    1. Set credentials cautiously and only when necessary.
    2. Prefer using `tunnel registry login` to pre-configure credentials with specific registries, which ensures credentials are only sent to appropriate registries.

Tunnel also supports providing credentials through CLI flags:

```shell
$ TUNNEL_PASSWORD=YOUR_PASSWORD tunnel image --username YOUR_USERNAME YOUR_PRIVATE_IMAGE
```

!!! warning
The CLI flag `--password` is available, but its use is not recommended for security reasons.

You can also store your credentials in `tunnel.yaml`.
For more information, please refer to [the documentation](../../references/configuration/config-file.md).

It can handle multiple sets of credentials as well:

```shell
$ export TUNNEL_USERNAME=USERNAME1,USERNAME2
$ export TUNNEL_PASSWORD=PASSWORD1,PASSWORD2
$ tunnel image YOUR_PRIVATE_IMAGE
```

In the example above, Tunnel attempts to use two pairs of credentials:

- USERNAME1/PASSWORD1
- USERNAME2/PASSWORD2

Please note that the number of usernames and passwords must be the same.

!!! note
`--password-stdin` doesn't support comma-separated passwords.
