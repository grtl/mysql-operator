package backupschedule

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"

	"k8s.io/api/batch/v1beta1"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/logging"
	"github.com/grtl/mysql-operator/pkg/util"
)

const (
	cronJobTemplate = "artifacts/backupschedule-cronjob.yaml"
	pvcTemplate     = "artifacts/backupschedule-pvc.yaml"
)

// Operator represents an object to manipulate Backup custom resources.
type Operator interface {
	AddBackupSchedule(backup *crv1.MySQLBackupSchedule) error
}

type backupScheduleOperator struct {
	clientset     versioned.Interface
	kubeClientset kubernetes.Interface
}

// NewBackupScheduleOperator returns a new Operator.
func NewBackupScheduleOperator(clientset versioned.Interface, kubeClientset kubernetes.Interface) Operator {
	return &backupScheduleOperator{
		clientset:     clientset,
		kubeClientset: kubeClientset,
	}
}

func (b *backupScheduleOperator) AddBackupSchedule(schedule *crv1.MySQLBackupSchedule) error {
	clusterName := schedule.Spec.Cluster
	cluster, err := b.clientset.CrV1().MySQLClusters(metav1.NamespaceDefault).
		Get(clusterName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if schedule.Spec.Storage.IsZero() {
		schedule.Spec.Storage = cluster.Spec.Storage
	}

	logging.LogBackupSchedule(schedule).Debug("Creating PVC.")
	err = b.createPVC(schedule)
	if err != nil {
		return err
	}

	logging.LogBackupSchedule(schedule).Debug("Creating cron job.")
	err = b.createCronJob(schedule)
	if err != nil {
		// Cleanup - remove already created PVC
		logging.LogBackupSchedule(schedule).WithField("fail", err).Warn("Reverting PVC creation.")
		removeErr := b.removePVC(schedule)
		return errors.NewAggregate([]error{err, removeErr})
	}

	return nil
}

func (b *backupScheduleOperator) createPVC(schedule *crv1.MySQLBackupSchedule) error {
	pvcInterface := b.kubeClientset.CoreV1().PersistentVolumeClaims(schedule.Namespace)
	pvc, err := pvcForSchedule(schedule)
	if err != nil {
		return err
	}

	_, err = pvcInterface.Create(pvc)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogBackupSchedule(schedule).Warn("PVC already exists")
	}

	return nil
}

func (b *backupScheduleOperator) createCronJob(schedule *crv1.MySQLBackupSchedule) error {
	cronJobInterface := b.kubeClientset.BatchV1beta1().CronJobs(schedule.Namespace)
	cronJob, err := cronJobForSchedule(schedule)
	if err != nil {
		return err
	}

	_, err = cronJobInterface.Create(cronJob)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogBackupSchedule(schedule).Warn("Backup already exists")
	}

	return nil
}

func (b *backupScheduleOperator) removePVC(schedule *crv1.MySQLBackupSchedule) error {
	pvcInterface := b.kubeClientset.CoreV1().Services(schedule.Namespace)
	return pvcInterface.Delete(schedule.Name, new(metav1.DeleteOptions))
}

func pvcForSchedule(schedule *crv1.MySQLBackupSchedule) (*v1.PersistentVolumeClaim, error) {
	pvc := new(v1.PersistentVolumeClaim)
	err := util.ObjectFromTemplate(schedule, pvc, pvcTemplate, FuncMap)
	return pvc, err
}

func cronJobForSchedule(schedule *crv1.MySQLBackupSchedule) (*v1beta1.CronJob, error) {
	cronJob := new(v1beta1.CronJob)
	err := util.ObjectFromTemplate(schedule, cronJob, cronJobTemplate, FuncMap)
	return cronJob, err
}
