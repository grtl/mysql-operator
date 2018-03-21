#!/usr/bin/env sh

set -e
set -v
ginkgo -cover -skipPackage e2e ./...
