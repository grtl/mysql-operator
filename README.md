# MySQL Operator
Kubernetes Custom Resource for MySQL.

[![Build Status](https://travis-ci.org/grtl/mysql-operator.svg?branch=master)](https://travis-ci.org/grtl/mysql-operator)
[![Coverage Status](https://coveralls.io/repos/github/grtl/mysql-operator/badge.svg?branch=master)](https://coveralls.io/github/grtl/mysql-operator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/grtl/mysql-operator)](https://goreportcard.com/report/github.com/grtl/mysql-operator)
[![GoDoc](https://godoc.org/github.com/grtl/mysql-operator?status.svg)](https://godoc.org/github.com/grtl/mysql-operator)

# Running MySQL Operator
In order for the custom resources to be properly processed and an actual MySQL cluster
deployed, a running instance of the MySQL Operator is required inside your Kubernetes
infrastructure. The operator listens for changes on `MySQLCluster` and `MySQLBackupSchedule`
custom resources and creates the appropriate objects.

## As a Kubernetes pod
This is the __recommended__ option. MySQL Operator will run as a pod inside your
Kubernetes cluster.
```bash
kubectl run mysql-operator --image=grtl/mysql-operator:latest
```

## As an out-of-cluster binary
Another option (suitable for development rather than a production-ready solution)
is to run the MySQL Operator binary outside of the Kubernetes cluster.

```bash
go get -u github.com/grtl/mysql-operator
mysql-operator -kubeconfig ~/.kube/config
```

Run code directly (ex. after making changes)
```bash
git clone https://github.com/grtl/mysql-operator && cd $_
go run -kubeconfig ~/.kube/config
```

# MySQL Operator Docker image
## Download from DockerHub
Download MySQL Operator image from DockerHub to easily deploy it in your
Kubernetes cluster.
```bash
docker pull grtl/mysql-operator
```
## Build yourself
Or build it yourself
```bash
GOOS=linux GOARCH=amd64 go build -o mysql-operator && docker build .
```

# Usage
Make sure the __operator is up and running__ before creating any custom resources.
All resources may be created with standard Kubernetes yaml files.

## Clusters
### Creating a cluster
First, a Kubernetes _Secret_ containing the database password needs to be created.
```bash
kubectl create secret generic my-secret --from-literal=password="P4sSw0rD"
```

Then you can create a cluster from a yaml file.
```bash
kubectl create -f cluster-config.yaml
```

#### Example `cluster-config.yaml` _(minimal)_
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"
```

#### Example `cluster-config.yaml` _(fully customized)_
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"   # Name of the secret containing the password
  port: 3306            # Port on which the service will expose the MySQL
  replicas: 2           # Number of replicas
  storage: "1Gi"        # Persistance Volume Claim size for each replica
  image: "mysql:latest" # MySQL image
```

### Restoring a cluster from the backup
While creating a cluster its data may be restored from an existing
[backup instance](#backup-instances). The only difference between the
configuration files for [creating a cluster](#creating-a-cluster) and
the one to restore a cluster from backup is an additional field
`fromBackup` field pointing to the [backup instance](#backup-instances).
```bash
kubectl create -f cluster-restore-config.yaml
```

#### Example `cluster-restore-config.yaml` (minimal)
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"
  fromBackup: "my-backup-2017-12-14-01-22"
```

#### Example `cluster-restore-config.yaml` (fully customized)
```yaml
apiVersion: cr.mysqloperator.grtl.github.com/v1
kind: MySQLCluster
metadata:
  name: "my-cluster"
spec:
  secret: "my-secret"
  fromBackup: "my-backup-2017-12-14-01-22"
  port: 3306
  replicas: 2
  storage: "1Gi"
  image: "mysql:latest"
```

### Deleting a cluster
Simply delete the cluster custom resource
```bash
kubectl delete mc "my-cluster"
```

## Backup Schedules
### Creating a backup schedule
You can create a backup schedule, which will automatically create backups
according to the schedule (cron job style).

```bash
kubectl create -f backup-config.yaml
```

Example `backup-config.yaml`
```yaml
apiVersion: cr.mysqlbackup.grtl.github.com/v1
kind: MySQLBackup
metadata:
  name: "my-backup"
spec:
  cluster: "my-cluster"
  time: "*/5 * * * *"    # Create a backup every 5 minutes
```

### Deleting a backup schedule
Simply delete the backup schedule custom resource
```bash
kubectl delete mbs "my-backup"
```

## Backup Instances
A backup schedule will create backup instances according to the schedule.

### Getting backup instances
Get all backup instances 
```bash
kubectl get mbi
```
Get all backup instances created within a given backup schedule
```bash
kubectl get mbi -l schedule="my-backup"
```

Get all backup instances created for a given cluster
```bash
kubectl get mbi -l cluster="my-cluster"
```

Standard `kubectl` output for backup instances lacks important fields, for
a valuable output we recommend using output flag with the following configuration.
```bash
kubectl get mbi -o custom-columns="NAME:metadata.name,STATUS:status.phase,\
    SCHEDULE:spec.schedule,CLUSTER:spec.cluster,CREATED:metadata.creationTimestamp"
```
Example output:
```
NAME                 STATUS      SCHEDULE       CLUSTER      CREATED
my-backup-instance   Completed   my-backup      my-cluster   2018-04-27T14:42:33Z
```

### Deleting a backup instance
Simply delete the backup instance custom resource
```bash
kubectl delete mbi "my-backup-instance"
```
Be aware that removing a backup instance will delete its contents from
the Persistent Volume.
