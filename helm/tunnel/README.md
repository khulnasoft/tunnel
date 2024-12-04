# Tunnel Scanner

Tunnel vulnerability scanner standalone installation.

## TL;DR;

```
$ helm install tunnel . --namespace tunnel --create-namespace
```

## Introduction

This chart bootstraps a Tunnel deployment on a [Kubernetes](http://kubernetes.io) cluster using the
[Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.12+
- Helm 3+

## Installing from the KhulnaSoft Chart Repository

```
helm repo add khulnasoft https://khulnasoft.github.io/helm-charts/
helm repo update
helm search repo tunnel
helm install my-tunnel khulnasoft/tunnel
```

## Installing the Chart

To install the chart with the release name `my-release`:

```
$ helm install my-release .
```

The command deploys Tunnel on the Kubernetes cluster in the default configuration. The [Parameters](#parameters)
section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`.

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the Tunnel chart and their default values.

| Parameter                                  | Description                                                                                                                                     | Default                        |
| ------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------ |
| `image.registry`                           | Image registry                                                                                                                                  | `docker.io`                    |
| `image.repository`                         | Image name                                                                                                                                      | `khulnasoft/tunnel`            |
| `image.tag`                                | Image tag                                                                                                                                       | `{TAG_NAME}`                   |
| `image.pullPolicy`                         | Image pull policy                                                                                                                               | `IfNotPresent`                 |
| `image.pullSecret`                         | The name of an imagePullSecret used to pull tunnel image from e.g. Docker Hub or a private registry                                             |                                |
| `replicaCount`                             | Number of Tunnel Pods to run                                                                                                                    | `1`                            |
| `tunnel.debugMode`                         | The flag to enable or disable Tunnel debug mode                                                                                                 | `false`                        |
| `tunnel.gitHubToken`                       | The GitHub access token to download Tunnel DB. More info: https://github.com/khulnasoft/tunnel#github-rate-limiting                             |                                |
| `tunnel.registryUsername`                  | The username used to log in at dockerhub. More info: https://khulnasoft.github.io/tunnel/dev/advanced/private-registries/docker-hub/            |                                |
| `tunnel.registryPassword`                  | The password used to log in at dockerhub. More info: https://khulnasoft.github.io/tunnel/dev/advanced/private-registries/docker-hub/            |                                |
| `tunnel.registryCredentialsExistingSecret` | Name of Secret containing dockerhub credentials. Alternative to the 2 parameters above, has precedence if set.                                  |                                |
| `tunnel.serviceAccount.annotations`        | Additional annotations to add to the Kubernetes service account resource                                                                        |                                |
| `tunnel.skipDBUpdate`                      | The flag to enable or disable Tunnel DB downloads from GitHub                                                                                   | `false`                        |
| `tunnel.dbRepository`                      | OCI repository to retrieve the tunnel vulnerability database from                                                                               | `ghcr.io/khulnasoft/tunnel-db` |
| `tunnel.cache.redis.enabled`               | Enable Redis as caching backend                                                                                                                 | `false`                        |
| `tunnel.cache.redis.url`                   | Specify redis connection url, e.g. redis://redis.redis.svc:6379                                                                                 | ``                             |
| `tunnel.cache.redis.ttl`                   | Specify redis TTL, e.g. 3600s or 24h                                                                                                            | ``                             |
| `tunnel.cache.redis.tls`                   | Enable Redis TLS with public certificates                                                                                                       | ``                             |
| `tunnel.serverToken`                       | The token to authenticate Tunnel client with Tunnel server                                                                                      | ``                             |
| `tunnel.existingSecret`                    | existingSecret if an existing secret has been created outside the chart. Overrides gitHubToken, registryUsername, registryPassword, serverToken | ``                             |
| `tunnel.podAnnotations`                    | Annotations for pods created by statefulset                                                                                                     | `{}`                           |
| `tunnel.extraEnvVars`                      | extraEnvVars to be set on the container                                                                                                         | `{}`                           |
| `service.name`                             | If specified, the name used for the Tunnel service                                                                                              |                                |
| `service.type`                             | Kubernetes service type                                                                                                                         | `ClusterIP`                    |
| `service.port`                             | Kubernetes service port                                                                                                                         | `4954`                         |
| `service.sessionAffinity`                  | Kubernetes service session affinity                                                                                                             | `ClientIP`                     |
| `httpProxy`                                | The URL of the HTTP proxy server                                                                                                                |                                |
| `httpsProxy`                               | The URL of the HTTPS proxy server                                                                                                               |                                |
| `noProxy`                                  | The URLs that the proxy settings do not apply to                                                                                                |                                |
| `nodeSelector`                             | Node labels for pod assignment                                                                                                                  |                                |
| `affinity`                                 | Affinity settings for pod assignment                                                                                                            |                                |
| `tolerations`                              | Tolerations for pod assignment                                                                                                                  |                                |
| `podAnnotations`                           | Annotations for pods created by statefulset                                                                                                     | `{}`                           |

The above parameters map to the env variables defined in [tunnel](https://github.com/khulnasoft/tunnel#configuration).

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

```
$ helm install my-release . \
       --namespace my-namespace \
       --set "service.port=9090" \
       --set "tunnel.vulnType=os\,library"
```

## Storage

This chart uses a PersistentVolumeClaim to reduce the number of database downloads between POD restarts or updates. The storageclass should have the reclaim policy `Retain`.

## Caching

You can specify a Redis server as cache backend. This Redis server has to be already present. You can use the [bitnami chart](https://bitnami.com/stack/redis/helm).
More Information about the caching backends can be found [here](https://github.com/khulnasoft/tunnel#specify-cache-backend).
