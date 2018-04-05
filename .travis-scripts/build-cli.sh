#!/usr/bin/env bash

set -e
set -v
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o msp github.com/grtl/mysql-operator/cli
