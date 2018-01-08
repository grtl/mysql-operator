#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

# Generate groups doesn't work with absolute paths so we need to hack it
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null \
    || echo ../../../k8s.io/code-generator)}

${CODEGEN_PKG}/generate-groups.sh all \
    github.com/grtl/mysql-operator/pkg/client \
    github.com/grtl/mysql-operator/pkg/apis cr:v1 \
    --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt