# Rootfs

Rootfs scanning is for special use cases such as

- Host machine
- [Root filesystem](../advanced/container/embed-in-dockerfile.md)
- [Unpacked filesystem](../advanced/container/unpacked-filesystem.md)

```bash
$ tunnel rootfs /path/to/rootfs
```

!!! note
Rootfs scanning works differently from the Filesystem scanning.
You should use `tunnel fs` to scan your local projects in CI/CD.
See [here](../scanner/vulnerability.md) for the differences.
