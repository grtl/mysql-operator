package backupinstance

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	batchv1 "k8s.io/api/batch/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/logging"
	"github.com/grtl/mysql-operator/pkg/util"
)

const (
	jobCreateTemplate = "artifacts/backupinstance-job-create.yaml"
	jobDeleteTemplate = "artifacts/backupinstance-job-delete.yaml"
)

// Operator represents an object to manipulate Backup custom resources.
type Operator interface {
	CreateBackup(backup *crv1.MySQLBackupInstance) error
	DeleteBackup(backup *crv1.MySQLBackupInstance) error
}

// NewBackupInstanceOperator returns a new Operator.
func NewBackupInstanceOperator(clientset versioned.Interface, kubeClientset kubernetes.Interface) Operator {
	return &backupInstanceOperator{
		clientset:     clientset,
		kubeClientset: kubeClientset,
	}
}

type backupInstanceOperator struct {
	clientset     versioned.Interface
	kubeClientset kubernetes.Interface
}

func (b *backupInstanceOperator) CreateBackup(backup *crv1.MySQLBackupInstance) error {
	if backup.Status.Phase != crv1.MySQLBackupScheduled {
		logging.LogBackupInstance(backup).Warn("Backup has already started.")
		return nil
	}

	// Make sure the cluster schedule exists (for now we only create backups within a schedule)
	schedulesInterface := b.clientset.CrV1().MySQLBackupSchedules(backup.Namespace)
	_, err := schedulesInterface.Get(backup.Spec.Schedule, metav1.GetOptions{})
	if err != nil {
		return err
	}

	return b.createJobCreate(backup)
}

func (b *backupInstanceOperator) DeleteBackup(backup *crv1.MySQLBackupInstance) error {
	return b.createJobDelete(backup)
}

func (b *backupInstanceOperator) createJobCreate(backup *crv1.MySQLBackupInstance) error {
	jobInterface := b.kubeClientset.BatchV1().Jobs(backup.Namespace)
	job, err := jobForBackup(backup, jobCreateTemplate)
	if err != nil {
		return err
	}

	_, err = jobInterface.Create(job)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogBackupInstance(backup).Warn("Backup create job already exists")
	}

	return nil
}

func (b *backupInstanceOperator) createJobDelete(backup *crv1.MySQLBackupInstance) error {
	jobInterface := b.kubeClientset.BatchV1().Jobs(backup.Namespace)
	job, err := jobForBackup(backup, jobDeleteTemplate)
	if err != nil {
		return err
	}

	_, err = jobInterface.Create(job)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogBackupInstance(backup).Warn("Backup delete job already exists")
	}

	return nil
}

func jobForBackup(backup *crv1.MySQLBackupInstance, template string) (*batchv1.Job, error) {
	job := new(batchv1.Job)
	err := util.ObjectFromTemplate(backup, job, template, FuncMap)
	return job, err
}
