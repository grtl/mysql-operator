#!/usr/bin/env bash
SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

echo -n >&2 "Installing pre-commit hooks... "
cat > ${SCRIPT_ROOT}/.git/hooks/pre-commit << 'EOF'
#!/bin/bash

echo >&2 "Running pre-commit hooks"

function fmt_error() {
	echo >&2 "Go files must be formatted with gofmt. Please run:"
	for fn in $unformatted; do
	    echo >&2 "  gofmt -w $PWD/$fn"
	done
	exit 1
}

function vet_error() {
	echo >&2 "Go files must pass the go vet checks. Please run:"
	echo >&2 "  go vet $(go list ./... | grep -v '/pkg/client/')"
	echo >&2 "and fix the errors"
	exit 2
}

function codegen_error() {
	echo >&2 "Generated code is out of date. Please run hack/update-codegen.sh"
	exit 3
}

function artifacts_error() {
	echo >&2 "Generated code is out of date. Please run hack/update-artifacts.sh"
	exit 4
}

gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
artifacts=$(git diff --cached --name-only --diff-filter=ACM | grep 'artifacts/.*\.yaml$')

if [ -z "$gofiles" ]; then
	echo >&2 "Skiped fmt, vet, codegen checks, no *.go files were modified"
else
	echo >&2 " - Checking go fmt"
	unformatted=$(gofmt -l $gofiles)
	[ -z "$unformatted" ] || fmt_error

	echo >&2 " - Checking go vet"
	unvet=$(go vet $(go list ./... | grep -v '/pkg/client/'))
	[ -z "$unvet" ] || vet_error

	echo >&2 " - Checking codegen"
	./hack/verify-codegen.sh > /dev/null || codegen_error
fi

if [ -z "$artifacts" ]; then
	echo >&2 "Skiped artifacts checks, no *.go files were modified"
else
	echo >&2 " - Checking artifacts"
	./hack/verify-artifacts.sh > /dev/null || artifacts_error
fi

EOF

chmod +x ${SCRIPT_ROOT}/.git/hooks/pre-commit

echo >&2 "Done."
