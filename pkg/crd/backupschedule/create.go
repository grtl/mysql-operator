package backupschedule

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/pkg/crd"
)

const (
	// CustomResourceName is the MySQLBackup custom resource definition qualified object name.
	CustomResourceName = "mysqlbackupschedules.cr.mysqloperator.grtl.github.com"
	definitionFilename = "artifacts/backupschedule-crd.yaml"
)

// CreateBackupScheduleCRD registers a MySQLBackupSchedule custom resource definition.
func CreateBackupScheduleCRD(namespace string, clientset apiextensions.Interface) error {
	err := crd.RegisterCRD(namespace, clientset, definitionFilename)
	if err != nil {
		return err
	}

	err = crd.WaitForCRDEstablished(clientset, CustomResourceName)
	if err != nil {
		return errors.NewAggregate([]error{err, crd.UnregisterCRD(clientset, CustomResourceName)})
	}
	return nil
}
