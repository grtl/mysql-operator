package cluster

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/util/errors"

	. "github.com/grtl/mysql-operator/crd"
)

const (
	clusterCustomResourceName = "mysqlclusters.cr.mysqloperator.grtl.github.com"
	clusterDefinitionFilename = "artifacts/mysql-crd.yaml"
)

// CreateClusterCRD registers a MySQLCluster custom resource in kubernetes api.
func CreateClusterCRD(clientset apiextensions.Interface) error {
	err := RegisterCRD(clientset, clusterDefinitionFilename)
	if err != nil {
		return err
	}

	err = WaitForCRDEstablished(clientset, clusterCustomResourceName)
	if err != nil {
		return errors.NewAggregate([]error{err, UnregisterCRD(clientset, clusterCustomResourceName)})
	}
	return nil
}
