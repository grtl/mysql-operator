# MySQL Operator
Kubernetes Custom Resource for MySQL.

[![Build Status](https://travis-ci.org/grtl/mysql-operator.svg?branch=master)](https://travis-ci.org/grtl/mysql-operator)
[![Coverage Status](https://coveralls.io/repos/github/grtl/mysql-operator/badge.svg?branch=master)](https://coveralls.io/github/grtl/mysql-operator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/grtl/mysql-operator)](https://goreportcard.com/report/github.com/grtl/mysql-operator)
[![GoDoc](https://godoc.org/github.com/grtl/mysql-operator?status.svg)](https://godoc.org/github.com/grtl/mysql-operator)

# Running MySQL Operator
In order for the custom resources to be properly processed and an actual MySQL cluster
deployed, a running instance of the MySQL Operator is required inside your Kubernetes
infrastructure. The operator listens for changes on `MySQLCluster` and `MySQLBackup`
custom resources and creates appropriate objects.

## As a Kubernetes pod
This is the __recommended__ option. MySQL Operator will run as a pod inside your
Kubernetes cluster.
```sh
> kubectl run mysql-operator --image=mysql-operator:latest
```

## As an out-of-cluster binary
Another option (suitable for development rather than a production-ready solution)
is to run the MySQL Operator binary outside of the Kubernetes cluster.

```sh
> go get -u github.com/grtl/mysql-operator
> mysql-operator -kubeconfig ~/.kube/config
```

Run code directly (ex. after making changes)
```sh
> git clone https://github.com/grtl/mysql-operator && cd $_
> go run -kubeconfig ~/.kube/config
```

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
