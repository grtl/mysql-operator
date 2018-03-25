package backup

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"

	"k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/grtl/mysql-operator/logging"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/util"
)

const (
	cronJobTemplate = "artifacts/backup-cronjob.yaml"
	pvcTemplate     = "artifacts/backup-pvc.yaml"
)

// Operator represents an object to manipulate Backup custom resources.
type Operator interface {
	ScheduleBackup(backup *crv1.MySQLBackup) error
}

type backupOperator struct {
	clientset    kubernetes.Interface
	verClientset versioned.Interface
}

// NewBackupOperator returns a new Operator.
func NewBackupOperator(verClientset versioned.Interface, clientset kubernetes.Interface) Operator {
	return &backupOperator{
		clientset:    clientset,
		verClientset: verClientset,
	}
}

func (b *backupOperator) ScheduleBackup(backup *crv1.MySQLBackup) error {
	clusterName := backup.Spec.Cluster
	_, err := b.verClientset.CrV1().MySQLClusters(metav1.NamespaceDefault).Get(clusterName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	logging.LogBackup(backup).Debug("Creating PVC.")
	err = b.createPVC(backup)
	if err != nil {
		return err
	}

	logging.LogBackup(backup).Debug("Creating cron job.")
	err = b.createCronJob(backup)
	if err != nil {
		// Cleanup - remove already created PVC
		logging.LogBackup(backup).WithField("error", err).Warn("Reverting PVC creation.")
		removeErr := b.removePVC(backup)
		return errors.NewAggregate([]error{err, removeErr})
	}

	return nil
}

func (b *backupOperator) createPVC(backup *crv1.MySQLBackup) error {
	pvcInterface := b.clientset.CoreV1().PersistentVolumeClaims(backup.Namespace)
	pvc, err := pvcForBackup(backup)
	if err != nil {
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
	err := util.ObjectFromTemplate(backup, pvc, pvcTemplate, FuncMap)
	return pvc, err
}

func cronJobForBackup(backup *crv1.MySQLBackup) (*v1beta1.CronJob, error) {
	cronJob := new(v1beta1.CronJob)
	err := util.ObjectFromTemplate(backup, cronJob, cronJobTemplate, FuncMap)
	return cronJob, err
}

func (b *backupOperator) removePVC(backup *crv1.MySQLBackup) error {
	pvcInterface := b.clientset.CoreV1().Services(backup.Namespace)
	return pvcInterface.Delete(backup.Name, new(metav1.DeleteOptions))
}
