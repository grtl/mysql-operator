package backup

import (
	"github.com/grtl/mysql-operator/logging"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/util"
	"k8s.io/api/batch/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"fmt"
	"k8s.io/api/core/v1"
)

const (
	cronJobTemplate = "artifacts/backup-cronjob.yaml"
	pvcTemplate = "artifacts/backup-pvc.yaml"
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
	//TODO: error if cluster doesn't exist
	//TODO: implement OneTime backups
	//TODO: PVC storage, PVC overriding
	//TODO: Default name if not provided

	logging.LogBackup(backup).Debug("Creating PVC.")
	err := b.createPVC(backup)
	if err != nil {
		return err
	}

	logging.LogBackup(backup).Debug("Creating cron job.")
	err = b.createCronJob(backup)
	if err != nil {
		//TODO: remove PVC
		return err
	}

	return nil
}

func (b *backupOperator) createPVC(backup *crv1.MySQLBackup) error {
	pvcInterface := b.clientset.CoreV1().PersistentVolumeClaims(backup.Namespace)
	pvc, err := pvcForBackup(backup)
	if err != nil {
		logging.LogBackup(backup).Debug("PVC error: " + fmt.Sprint(err))
		return err
	}

	_, err = pvcInterface.Create(pvc)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogBackup(backup).Warn("PVC already exists")
	}

	return nil
}

func (b *backupOperator) createCronJob(backup *crv1.MySQLBackup) error {
	cronJobInterface := b.clientset.BatchV1beta1().CronJobs(backup.Namespace)
	cronJob, err := cronJobForBackup(backup)
	if err != nil {
		logging.LogBackup(backup).Debug("Cron job error: " + fmt.Sprint(err))
		return err
	}

	_, err = cronJobInterface.Create(cronJob)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogBackup(backup).Warn("Backup already exists")
	}

	return nil
}

func pvcForBackup(backup *crv1.MySQLBackup) (*v1.PersistentVolumeClaim, error) {
	pvc := new(v1.PersistentVolumeClaim)
	err := util.ObjectFromTemplate(backup, pvc, pvcTemplate)
	return pvc, err
}

func cronJobForBackup(backup *crv1.MySQLBackup) (*v1beta1.CronJob, error) {
	cronJob := new(v1beta1.CronJob)
	err := util.ObjectFromTemplate(backup, cronJob, cronJobTemplate)
	return cronJob, err
}
