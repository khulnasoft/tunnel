# Signature Verification

All binaries and container images are signed by [Cosign](https://github.com/sigstore/cosign).

## Verifying container image

Use the following command for keyless [verification](https://docs.sigstore.dev/cosign/verify/):

```shell
cosign verify khulnasoft/tunnel:<version> \
--certificate-identity-regexp 'https://github\.com/khulnasoft/tunnel/\.github/workflows/.+' \
--certificate-oidc-issuer "https://token.actions.githubusercontent.com"
```

You should get the following output

```
Verification for index.docker.io/khulnasoft/tunnel:latest --
The following checks were performed on each of these signatures:
   - The cosign claims were validated
   - Existence of the claims in the transparency log was verified offline
   - The code-signing certificate was verified using trusted certificate authority certificates

   ....
```

## Verifying binary

Download the required tarball, associated signature and certificate files from the [GitHub Release](https://github.com/khulnasoft/tunnel/releases).

Use the following command for keyless verification:

```shell
cosign verify-blob <path to binray> \
--certificate <path to cert> \
--signature <path to sig> \
--certificate-identity-regexp 'https://github\.com/khulnasoft/tunnel/\.github/workflows/.+' \
--certificate-oidc-issuer "https://token.actions.githubusercontent.com"
```

You should get the following output

```
Verified OK
```

## Verifying a GPG signature

RPM and Deb packages are also signed by GPG.

### Verifying RPM

The public key is available at <https://khulnasoft.github.io/tunnel-repo/rpm/public.key>.

First, download and import the key:

```shell
curl https://khulnasoft.github.io/tunnel-repo/rpm/public.key \
--output pub.key
rpm --import pub.key
rpm -q --queryformat "%{SUMMARY}\n" $(rpm -q gpg-pubkey)
```

You should get the following output:

```
gpg(tunnel)
```

Then you can verify the signature:

```shell
curl -L https://github.com/khulnasoft/tunnel/releases/download/<version>/<file name>.rpm \
--output tunnel.rpm
rpm -K tunnel.rpm
```

You should get the following output

```
tunnel.rpm: digests signatures OK
```
