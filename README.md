# MySQL Operator
Kubernetes Custom Resource for MySQL.

[![Build Status](https://travis-ci.org/grtl/mysql-operator.svg?branch=master)](https://travis-ci.org/grtl/mysql-operator)
[![Coverage Status](https://coveralls.io/repos/github/grtl/mysql-operator/badge.svg?branch=master)](https://coveralls.io/github/grtl/mysql-operator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/grtl/mysql-operator)](https://goreportcard.com/report/github.com/grtl/mysql-operator)
[![GoDoc](https://godoc.org/github.com/grtl/mysql-operator?status.svg)](https://godoc.org/github.com/grtl/mysql-operator)

# MySQL Operator Docker image
## Download from DockerHub
Download MySQL Operator image from DockerHub to easily deploy it in your
Kubernetes cluster.
```sh
> docker pull grtl/mysql-operator
```
## Build yourself
Or build it yourself
```sh
> GOOS=linux GOARCH=amd64 go build -o mysql-operator
> docker build .
```
