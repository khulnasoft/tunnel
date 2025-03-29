# Installing Tunnel

In this section you will find an aggregation of the different ways to install Tunnel. Installation options are labeled as either "Official" or "Community". Official installations are developed by the Tunnel team and supported by it. Community installations could be developed by anyone from the Tunnel community, and collected here for your convenience. For support or questions about community installations, please contact the original developers.

!!! note
    If you are looking to integrate Tunnel into another system, such as CI/CD, IDE, Kubernetes, etc, please see [Ecosystem section](../ecosystem/index.md) to explore integrations of Tunnel with other tools.

## Container image (Official)

Use one of the official Tunnel images:

| Registry | Repository | Link |
| --- | --- | --- |
| Docker Hub | `docker.io/khulnasoft/tunnel` | https://hub.docker.com/r/khulnasoft/tunnel |
| GitHub Container Registry (GHCR) | `ghcr.io/khulnasoft/tunnel` | https://github.com/orgs/khulnasoft/packages/container/package/tunnel |
| AWS Elastic Container Registry (ECR) | `public.ecr.aws/khulnasoft/tunnel` | https://gallery.ecr.aws/khulnasoft/tunnel |

!!! Tip
    It is advisable to mount a persistent [cache dir](../docs/configuration/cache.md) on the host into the Tunnel container.

!!! Tip
    For scanning container images with Tunnel, mount the container engine socket from the host into the Tunnel container.

Example:

``` bash
docker run -v /var/run/docker.sock:/var/run/docker.sock -v $HOME/Library/Caches:/root/.cache/ khulnasoft/tunnel:{{ git.tag[1:] }} image python:3.4-alpine
```

## GitHub Release (Official)

1. Download the file for your operating system/architecture from [GitHub Release assets](https://github.com/khulnasoft/tunnel/releases/tag/{{ git.tag }}).  
2. Unpack the downloaded archive (`tar -xzf ./tunnel.tar.gz`).
3. Make sure the binary has execution bit turned on (`chmod +x ./tunnel`).

## Install Script (Official)

For convenience, you can use the install script to download and install Tunnel from GitHub Release.

```bash
curl -sfL https://raw.githubusercontent.com/khulnasoft/tunnel/main/contrib/install.sh | sudo sh -s -- -b /usr/local/bin {{ git.tag }}
```

## RHEL/CentOS (Official)

=== "Repository"
    Add repository setting to `/etc/yum.repos.d`.

    ``` bash
    cat << EOF | sudo tee -a /etc/yum.repos.d/tunnel.repo
    [tunnel]
    name=Tunnel repository
    baseurl=https://khulnasoft.github.io/tunnel-repo/rpm/releases/\$basearch/
    gpgcheck=1
    enabled=1
    gpgkey=https://khulnasoft.github.io/tunnel-repo/rpm/public.key
    EOF
    sudo yum -y update
    sudo yum -y install tunnel
    ```

=== "RPM"

    ``` bash
    rpm -ivh https://github.com/khulnasoft/tunnel/releases/download/{{ git.tag }}/tunnel_{{ git.tag[1:] }}_Linux-64bit.rpm
    ```

## Debian/Ubuntu (Official)

=== "Repository"
    Add repository setting to `/etc/apt/sources.list.d`.

    ``` bash
    sudo apt-get install wget gnupg
    wget -qO - https://khulnasoft.github.io/tunnel-repo/deb/public.key | gpg --dearmor | sudo tee /usr/share/keyrings/tunnel.gpg > /dev/null
    echo "deb [signed-by=/usr/share/keyrings/tunnel.gpg] https://khulnasoft.github.io/tunnel-repo/deb generic main" | sudo tee -a /etc/apt/sources.list.d/tunnel.list
    sudo apt-get update
    sudo apt-get install tunnel
    ```

=== "DEB"

    ``` bash
    wget https://github.com/khulnasoft/tunnel/releases/download/{{ git.tag }}/tunnel_{{ git.tag[1:] }}_Linux-64bit.deb
    sudo dpkg -i tunnel_{{ git.tag[1:] }}_Linux-64bit.deb
    ```

## Homebrew (Official)

Homebrew for macOS and Linux.

```bash
brew install tunnel
```

## Windows (Official)

1. Download tunnel_x.xx.x_windows-64bit.zip file from [releases page](https://github.com/khulnasoft/tunnel/releases/).
2. Unzip file and copy to any folder.

## Arch Linux (Community)

Arch Linux Package Repository.

```bash
sudo pacman -S tunnel
```

References: 
- <https://archlinux.org/packages/extra/x86_64/tunnel/>
- <https://gitlab.archlinux.org/archlinux/packaging/packages/tunnel/-/blob/main/PKGBUILD>


## MacPorts (Community)

[MacPorts](https://www.macports.org) for macOS.

```bash
sudo port install tunnel
```

References:
- <https://ports.macports.org/port/tunnel/details/>

## Nix/NixOS (Community)

Nix package manager for Linux and macOS.

=== "Command line"
    `nix-env --install -A nixpkgs.tunnel`

=== "Configuration"
    ```nix
    # your other config ...
    environment.systemPackages = with pkgs; [
      # your other packages ...
      tunnel
    ];
    ```

=== "Home Manager"
    ```nix
    # your other config ...
    home.packages = with pkgs; [
      # your other packages ...
      tunnel
    ];
    ```

References: 

-  https://github.com/NixOS/nixpkgs/blob/master/pkgs/tools/admin/tunnel/default.nix

## FreeBSD (Official)

Pkg package manager for FreeBSD.

```bash
pkg install tunnel
```

## asdf/mise (Community)

[asdf](https://github.com/asdf-vm/asdf) and [mise](https://github.com/jdx/mise) are quite similar tools you can use to install tunnel.
See their respective documentation for more information of how to install them and use them:

- [asdf](https://asdf-vm.com/guide/getting-started.html)
- [mise](https://mise.jdx.dev/getting-started.html)

The plugin used by both tools is developed [here](https://github.com/zufardhiyaulhaq/asdf-tunnel)


=== "asdf"
    A basic global installation is shown below, for specific version or/and local version to a directory see "asdf" documentation.

    ```shell
    # Install plugin
    asdf plugin add tunnel https://github.com/zufardhiyaulhaq/asdf-tunnel.git

    # Install latest version
    asdf install tunnel latest

    # Set a version globally (on your ~/.tool-versions file)
    asdf global tunnel latest

    # Now tunnel commands are available
    tunnel --version
    ```

=== "mise"
    A basic global installation is shown below, for specific version or/and local version to a directory see "mise" documentation.

    ``` shell
    # Install plugin and install latest version
    mise install tunnel@latest

    # Set a version globally (on your ~/.tool-versions file)
    mise use -g tunnel@latest

    # Now tunnel commands are available
    tunnel --version
    ```
