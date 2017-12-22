# MySQL Operator
Kubernetes Custom Resource for MySQL.

[![Build Status](https://travis-ci.org/grtl/mysql-operator.svg?branch=master)](https://travis-ci.org/grtl/mysql-operator)
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
> GOOS=linux GOARCH=amd64 go build -o docker/mysql-operator
> cd docker && docker build .
```
