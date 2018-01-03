#!/usr/bin/env sh

set -e
GOOS=linux GOARCH=amd64 go build -o mysql-operator
docker build . -t mysql-operator:testing