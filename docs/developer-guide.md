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

## Clone the project
All projects in Go must exist in the `$GOPATH` 
```bash
> mkdir -p "$GOPATH/src/github.com/grtl" && cd $_
> git clone git@github.com:grtl/mysql-operator.git
# Replace with https://github.com/grtl/mysql-operator for https
```

## Install _godep_
Install the lastest version via go get:
```bash
> go get github.com/tools/godep
```

Enable vendor:
```bash
> export GO15VENDOREXPERIMENT=1
```

## Install dependencies
Run the following commands in your project directory to install all dependencies:
```bash
> godep restore ./...
```

For further development - to install all project dependencies:
```bash
> go get ./...
```
And update the vendor system via godep:
```bash
> godep save ./...
```

You're ready to rock!

# Regenerating _pkg/client_ code
```bash
> vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/grtl/mysql-operator/pkg/client \
github.com/grtl/mysql-operator/pkg/apis cr:v1
```
