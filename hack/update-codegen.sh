#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
vendor/k8s.io/code-generator/generate-groups.sh all \
    github.com/grtl/mysql-operator/pkg/client \
    github.com/grtl/mysql-operator/pkg/apis cr:v1 \
    --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt
