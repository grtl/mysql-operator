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
kubectl run mysql-operator --image=grtl/mysql-operator:latest
```

## As an out-of-cluster binary
Another option (suitable for development rather than a production-ready solution)
is to run the MySQL Operator binary outside of the Kubernetes cluster.

```sh
go get -u github.com/grtl/mysql-operator
mysql-operator -kubeconfig ~/.kube/config
```

Run code directly (ex. after making changes)
```sh
git clone https://github.com/grtl/mysql-operator && cd $_
go run -kubeconfig ~/.kube/config
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

# Usage
Make sure the __operator is up and running__ before creating any custom resources.
All resources may be created with standard Kubernetes yaml files.

## Creating a cluster
### Using `kubectl`
#### Creating a Secret
First, a Kubernetes _Secret_ containing the database password needs to be created.
```sh
kubectl create secret generic my-secret --from-literal=password="P4sSw0rD"
```

#### Creating a MySQLCluster
```sh
kubectl apply -f cluster-config.yaml
```

Example `cluster-config.yaml` (minimal)
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"
```

Example `cluster-config.yaml` (fully customized)
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"        # Name of the secret containing the password
  port: 3306                 # Port on which the service will expose the MySQL
  replicas: 2                # Number of replicas
  storage: "1Gi"             # Persistance Volume Claim size
  mysqlImage: "mysql:latest" # MySQL image
```

## Restoring a Cluster from Backup
At the moment there is no way to find the name of the backup instance from the
Kubernetes structure [see issue](https://github.com/grtl/mysql-operator/issues/106).
They need to be manually checked by connecting to the backup pod and listing the
directories.

### Using `kubectl`
```sh
kubectl apply -f cluster-restore-config.yaml
```

Example `cluster-restore-config.yaml` (minimal)
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"
  fromBackup:
    backupName: "my-backup"      # Backup name
    instance: "2017-12-14-01-22" # Backup instance created in the backup job
```

Example `cluster-restore-config.yaml` (fully customized)
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"
  fromBackup:
    backupName: "my-backup"
    instance: "2017-12-14-01-22"
  port: 3306
  replicas: 2
  storage: "1Gi"
  mysqlImage: "mysql:latest"
```

### Geting the backup instance name
__This is a temporary solution for finding the backup instance name__
```yaml
kind: Pod
apiVersion: v1
metadata:
  name: my-pod
spec:
  volumes:
    - name: backup
      persistentVolumeClaim:
       claimName: my-backup
  containers:
    - name: backup-container
      image: alpine:latest
      volumeMounts:
        - mountPath: "/mysql/my-backup"
          name: backup
```
```sh
kubectl exec -it my-pod -- /bin/sh -c "ls /mysql/my-backup"
```

## Deleting a Cluster
Simply delete the cluster custom resource
```sh
kubectl delete mysqlcluster "my-cluster"
```

## Creating a Backup
### Using `kubectl`
```sh
kubectl apply -f backup-config.yaml
```
Example `backup-config.yaml`
```
apiVersion: cr.mysqlbackup.grtl.github.com/v1
kind: MySQLBackup
metadata:
  name: "my-backup"
spec:
  cluster: "my-cluster"
  time: "*/1 * * * *"
```
