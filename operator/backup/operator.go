package backup

import (
	"k8s.io/client-go/kubernetes"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// Operator represents an object to manipulate Backup custom resources.
type Operator interface {
	ScheduleBackup(backup *crv1.MySQLBackup) error
}

type backupOperator struct {
	clientset kubernetes.Interface
}

// NewBackupOperator returns a new Operator.
func NewBackupOperator(clientset kubernetes.Interface) Operator {
	return &backupOperator{
		clientset: clientset,
	}
}

func (b *backupOperator) ScheduleBackup(backup *crv1.MySQLBackup) error {
	//TODO: implement backup scheduling
	return nil
}
