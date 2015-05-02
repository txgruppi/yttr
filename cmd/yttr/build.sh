#!/bin/sh

# based on https://github.com/CotaPreco/Horus/blob/develop/build.sh

VERSION=$(date +%Y.%m.%d.)$(git rev-parse --short HEAD)

if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
  VERSION="$VERSION-dirty"
fi

PLATFORMS="darwin/amd64 linux/386 linux/amd64 linux/arm windows/386 windows/amd64"

gox -ldflags "-w -s -X github.com/txgruppi/yttr.Version $VERSION" -osarch="$PLATFORMS" -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
