package cluster

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/pkg/crd"
)

const (
	// CustomResourceName is the MySQLCluster custom resource definition qualified object name.
	CustomResourceName = "mysqlclusters.cr.mysqloperator.grtl.github.com"
	definitionFilename = "artifacts/cluster-crd.yaml"
)

// CreateClusterCRD registers a MySQLCluster custom resource in kubernetes api.
func CreateClusterCRD(clientset apiextensions.Interface) error {
	err := crd.RegisterCRD(clientset, definitionFilename)
	if err != nil {
		return err
	}

	err = crd.WaitForCRDEstablished(clientset, CustomResourceName)
	if err != nil {
		return errors.NewAggregate([]error{err, crd.UnregisterCRD(clientset, CustomResourceName)})
	}
	return nil
}
