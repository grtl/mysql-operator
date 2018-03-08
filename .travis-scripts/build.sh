#!/usr/bin/env sh

set -e
set -v
GOOS=linux GOARCH=amd64 go build -o mysql-operator
docker build . -t mysql-operator:testing