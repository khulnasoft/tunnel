#!/bin/bash

DEBIAN_RELEASES=$(debian-distro-info --supported)
UBUNTU_RELEASES=$(sort -u <(ubuntu-distro-info --supported-esm) <(ubuntu-distro-info --supported))

cd tunnel-repo/deb

for release in generic ${DEBIAN_RELEASES[@]} ${UBUNTU_RELEASES[@]}; do
  echo "Removing deb package of $release"
  reprepro -A i386 remove $release tunnel
  reprepro -A amd64 remove $release tunnel
  reprepro -A arm64 remove $release tunnel
done

for release in generic ${DEBIAN_RELEASES[@]} ${UBUNTU_RELEASES[@]}; do
  echo "Adding deb package to $release"
  reprepro includedeb $release ../../dist/*Linux-32bit.deb
  reprepro includedeb $release ../../dist/*Linux-64bit.deb
  reprepro includedeb $release ../../dist/*Linux-ARM64.deb
done

git add .
git commit -m "Update deb packages"
git push origin main
