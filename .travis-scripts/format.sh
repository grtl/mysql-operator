#!/usr/bin/env sh

set -e
set -v
test -z "$(go fmt $(go list ./...))"
go vet $(go list ./... | grep -v '/pkg/client/')
