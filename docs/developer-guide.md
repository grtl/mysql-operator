# Setting up development environment
## Prerequisites
Make sure you have Go installed and correctly configured `$GOPATH`.
The output should look similar to:
```bash
$ go version
go version go1.9.2 darwin/amd64
$ echo "$GOPATH"
/Users/username/Documents/Go
```

## Clone the project
All projects in Go must exist in the `$GOPATH` 
```bash
mkdir -p "$GOPATH/src/github.com/grtl" && cd $_
git clone git@github.com:grtl/mysql-operator.git
# Replace with https://github.com/grtl/mysql-operator for https
```

## Install _godep_
Install the lastest version via go get:
```bash
go get github.com/tools/godep
```

Enable vendor:
```bash
export GO15VENDOREXPERIMENT=1
```

## Install dependencies
Run the following commands in your project directory to install all dependencies:
```bash
godep restore ./...
```

For further development - to install all new dependencies:
```bash
go get ./...
```
And update the vendor system via godep:
```bash
godep save ./...
```

You're ready to rock!

## Useful hooks
To avoid common mistakes it's recommended to add simple checks to your
pre-commit hooks. Paste the following lines to `.git/hooks/pre-commit` (create
the file if it doesn't exist) and make it executable. The script will run

* `go fmt` to check if modified Go files are correctly formatted
* `go vet` to check for suspicious constructs in the modified Go files
* `hack/verify-codegen.sh` to check if the auto-generated files are up to date

```bash
#!/usr/bin/env bash

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

gofiles=$(git diff --name-only --diff-filter=ACM | grep '.go$')
[ -z "$gofiles" ] && echo >&2 "Skiped checks, no *.go files were modified" && exit 0

echo >&2 " - Checking go fmt"
unformatted=$(gofmt -l $gofiles)
[ -z "$unformatted" ] || fmt_error

echo >&2 " - Checking go vet"
unvet=$(go vet $gofiles)
[ -z "$unvet" ] || vet_error

echo >&2 " - Checking generated code"
./hack/verify-codegen.sh > /dev/null || codegen_error
```

# Regenerating _pkg/client_ code
First install kubernetes code-generator script, by running:

```bash
err="$(go get k8s.io/code-generator 2>&1)"; \
[[ "$err" =~ "package k8s.io/code-generator: no Go files in .*" ]] \
&& echo >&2 "Success" || echo >&2 "$err"
```

The files can be regenerated using the `hack/update-codegen.sh` (make sure
to run it from project root directory).
```bash
./hack/update-codegen.sh
```
