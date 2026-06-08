#!/bin/bash
set -e

ARCH=$(uname -m)
if [ "$ARCH" = "x86_64" ]; then
    ALPINE_ARCH="x86_64"
elif [ "$ARCH" = "aarch64" ]; then
    ALPINE_ARCH="aarch64"   # works on M1 macs running linux
else
    echo "unsupported arch: $ARCH"
    exit 1
fi

ALPINE_VERSION="3.19.0"
URL="https://dl-cdn.alpinelinux.org/alpine/v3.19/releases/$ALPINE_ARCH/alpine-minirootfs-$ALPINE_VERSION-$ALPINE_ARCH.tar.gz"

echo "setting up rootfs..."
mkdir -p rootfs
wget -q --show-progress "$URL" -O /tmp/alpine.tar.gz
tar -xzf /tmp/alpine.tar.gz -C rootfs
rm /tmp/alpine.tar.gz

mkdir -p rootfs/proc rootfs/sys rootfs/dev

echo "done. run: make build"