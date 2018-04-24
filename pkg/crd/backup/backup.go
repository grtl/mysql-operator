package backup

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/util/errors"

	. "github.com/grtl/mysql-operator/pkg/crd"
)

const (
	// CustomResourceName is the MySQLBackup custom resource definition qualified object name.
	CustomResourceName = "mysqlbackups.cr.mysqloperator.grtl.github.com"
	definitionFilename = "artifacts/backup-crd.yaml"
)

// CreateBackupCRD registers a MySQLBackup custom resource in kubernetes api.
func CreateBackupCRD(clientset apiextensions.Interface) error {
	err := RegisterCRD(clientset, definitionFilename)
	if err != nil {
		return err
	}

	err = WaitForCRDEstablished(clientset, CustomResourceName)
	if err != nil {
		return errors.NewAggregate([]error{err, UnregisterCRD(clientset, CustomResourceName)})
	}
	return nil
}
