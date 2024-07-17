#!/usr/bin/env bash

set -e


version=$(git rev-parse --short HEAD)-$(date +%Y%m%d%H%M%S)

target=$1
export GOOS=$(go env GOOS)
export GOARCH=$(go env GOARCH)
case $target in
"linux")
  export GOOS=linux
  export GOARCH=amd64
  ;;
esac

#echo "ENV GOOS:$GOOS, GOARCH:$GOARCH"

program=$(go list -m)

set -x
mkdir -p bin/ && go build -v -ldflags "-w -s  -libgcc=none  -X cobra-example/main.Version=${version}" -o bin/$program  main.go
set +x