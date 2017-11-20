# MySQL Operator
Kubernetes Custom Resource Definition for MySQL Cluster.

# Generating `pkg/client` code
```sh
$ vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/grtl/mysql-operator/pkg/client \
github.com/grtl/mysql-operator/pkg/apis cr:v1
```
