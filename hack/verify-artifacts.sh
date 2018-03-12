#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
go-bindata -nometadata -o _tmp_artifacts.go -pkg artifacts artifacts/*.yaml
gofmt -s -w _tmp_artifacts.go

ret=0
diff -Naup _tmp_artifacts.go ${SCRIPT_ROOT}/artifacts/artifacts.go || ret=$?

rm _tmp_artifacts.go

if [[ $ret -eq 0 ]]
then
    echo "Artifacts up to date."
else
    echo "Artifacts is out of date. Please run hack/update-artifacts.sh"
    exit 1
fi
