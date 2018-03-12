#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
go-bindata -nometadata -o ${SCRIPT_ROOT}/artifacts/artifacts.go -pkg artifacts ${SCRIPT_ROOT}/artifacts/*.yaml
gofmt -s -w ${SCRIPT_ROOT}/artifacts/artifacts.go
