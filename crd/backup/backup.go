package backup

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/util/errors"

	. "github.com/grtl/mysql-operator/crd"
)

const (
	backupCustomResourceName = "mysqlbackups.cr.mysqloperator.grtl.github.com"
	backupDefinitionFilename = "artifacts/backup-crd.yaml"
)

// CreateBackupCRD registers a MySQLBackup custom resource in kubernetes api.
func CreateBackupCRD(clientset apiextensions.Interface) error {
	err := RegisterCRD(clientset, backupDefinitionFilename)
	if err != nil {
		return err
	}

	err = WaitForCRDEstablished(clientset, backupCustomResourceName)
	if err != nil {
		return errors.NewAggregate([]error{err, UnregisterCRD(clientset, backupCustomResourceName)})
	}
	return nil
}
