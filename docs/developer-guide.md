# Setting up development environment
## Prerequisites
Make sure you have Go installed and correctly configured `$GOPATH`.
The output should look similar to:
```bash
> go version
go version go1.9.2 darwin/amd64
> echo "$GOPATH"
/Users/username/Documents/Go
```

## Install _Dep_
On macOS you may install the latest version via Homebrew:
```bash
> brew install dep
```
For other OS follow the [instructions](https://github.com/golang/dep#setup).

## Clone the project
All projects in Go must exist in the `$GOPATH` 
```bash
> mkdir -p "$GOPATH/src/github.com/grtl" && cd $_
> git clone git@github.com:grtl/mysql-operator.git
# Replace with https://github.com/grtl/mysql-operator for https
```

## Install dependencies
Run the following commands to install all dependencies.
```bash
> dep ensure
```
You're ready to rock!

# Regenerating _pkg/client_ code
```bash
> vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/grtl/mysql-operator/pkg/client \
github.com/grtl/mysql-operator/pkg/apis cr:v1
```