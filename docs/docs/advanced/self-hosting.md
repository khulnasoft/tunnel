# Self-Hosting Tunnel's Databases

This document explains how to host Tunnel's [external dependencies](./air-gap.md) in your own infrastructure to prevent external network access. If you haven't already, please familiarize yourself with the [Databases document](../configuration/db.md) that explains about the different databases used by Tunnel and the different configuration options that control them. This guide assumes you are already familiar with the concepts explained there.

## OCI databases

The following [Tunnel Databases](../configuration/db.md) are packaged as OCI images:

- `tunnel-db`
- `tunnel-java-db`
- `tunnel-checks`

To host these databases in your own infrastructure:

### Make a local copy

Use any container registry manipulation tool (e.g , [crane](https://github.com/google/go-containerregistry/blob/main/cmd/crane/doc/crane.md, [ORAS](https://oras.land), [regclient](https://github.com/regclient/regclient/tree/main)) to copy the images to your destination registry.

!!! note
You will need to keep the databases updated in order to maintain relevant scanning results over time.

### Configure Tunnel

Use the appropriate [database location flags](../configuration/db.md#database-locations) to change the db-repository location:

- `--db-repository`
- `--java-db-repository`
- `--checks-bundle-repository`

### Authentication

If the registry requires authentication, you can configure it as described in the [private registry authentication document](../advanced/private-registries/index.md).

### OCI Media Types

When serving, proxying, or manipulating Tunnel's databases, note that the media type of the OCI layer is not a standard container image type:

| DB               | Media Type                                                   | Reference                                                                      |
| ---------------- | ------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `tunnel-db`      | `application/vnd.khulnasoft.tunnel.db.layer.v1.tar+gzip`     | <https://go.khulnasoft.com/tunnel-db/pkgs/container/tunnel-db>                 |
| `tunnel-java-db` | `application/vnd.khulnasoft.tunnel.javadb.layer.v1.tar+gzip` | https://github.com/khulnasoft-lab/tunnel-java-db/pkgs/container/tunnel-java-db |
| `tunnel-checks`  | `application/vnd.oci.image.manifest.v1+json`                 | https://github.com/aquasecurity/trivy-checks/pkgs/container/tunnel-checks      |

## Manual cache population

Tunnel uses a local cache directory to store the database files, as described in the [cache](../configuration/cache.md) document.
You can download the databases files and surgically populate the Tunnel cache directory with them.

### Downloading the DB files

On a machine with internet access, pull the database container archive from the public registry into your local workspace:

Note that these examples operate in the current working directory.

=== "Using ORAS"
This example uses [ORAS](https://oras.land), but you can use any other container registry manipulation tool.

    ```shell
    oras pull ghcr.io/khulnasoft/tunnel-db:2
    ```

    You should now have a file called `db.tar.gz`. Next, extract it to reveal the db files:

    ```shell
    tar -xzf db.tar.gz
    ```

=== "Using Tunnel"
This example uses Tunnel to pull the database container archive. The `--cache-dir` flag makes Tunnel download the database files into our current working directory. The `--download-db-only` flag tells Tunnel to only download the database files, not to scan any images.

    ```shell
    tunnel image --cache-dir . --download-db-only
    ```

You should now have 2 new files, `metadata.json` and `tunnel.db`. These are the Tunnel DB files, copy them over to the air-gapped environment.

### Populating the Tunnel Cache

In order to populate the cache, you need to identify the location of the cache directory. If it is under the default location, you can run the following command to find it:

```shell
tunnel -h | grep cache
```

For the example, we will assume the `TUNNEL_CACHE_DIR` variable holds the cache location:

```shell
TUNNEL_CACHE_DIR=/home/user/.cache/tunnel
```

Put the Tunnel DB files in the Tunnel cache directory under a `db` subdirectory:

```shell
# ensure cache db directory exists
mkdir -p ${TUNNEL_CACHE_DIR}/db
# copy the db files
cp /path/to/tunnel.db /path/to/metadata.json ${TUNNEL_CACHE_DIR}/db/
```

### Java DB adaptations

For Java DB the process is the same, except for the following:

1. Image location is `ghcr.io/khulnasoft/tunnel-java-db:1`
2. Archive file name is `javadb.tar.gz`
3. DB file name is `tunnel-java.db`

## VEX Hub

### Make a local copy

To make a copy of VEX Hub in a location that is accessible to Tunnel.

1. Download the [VEX Hub](https://github.com/khulnasoft/vexhub) archive from: <https://github.com/khulnasoft/vexhub/archive/refs/heads/main.zip>.
1. Download the [VEX Hub Repository Manifest](https://github.com/khulnasoft/vex-repo-spec#2-repository-manifest) file from: <https://github.com/khulnasoft/vexhub/blob/main/vex-repository.json>.
1. Create or identify an internal HTTP server that can serve the VEX Hub repository in your environment (e.g `https://server.local`).
1. Make the downloaded archive file available for serving from your server (e.g `https://server.local/main.zip`).
1. Modify the downloaded manifest file's [Location URL](https://github.com/khulnasoft/vex-repo-spec?tab=readme-ov-file#locations-subfields) field to the URL of the archive file on your server (e.g `url: https://server.local/main.zip`).
1. Make the manifest file available for serving from your server under the `/.well-known` path (e.g `https://server.local/.well-known/vex-repository.json`).

### Configure Tunnel

To configure Tunnel to use the local VEX Repository:

1. Locate your [Tunnel VEX configuration file](../supply-chain/vex/repo/#configuration-file) by running `tunnel vex repo init`. Make the following changes to the file.
1. Disable the default VEX Hub repo (`enabled: false`)
1. Add your internal VEX Hub repository as a [custom repository](../supply-chain/vex/repo/#custom-repositories) with the URL pointing to your local server (e.g `url: https://server.local`).

### Authentication

If your server requires authentication, you can configure it as described in the [VEX Repository Authentication document](../supply-chain/vex/repo/#authentication).
