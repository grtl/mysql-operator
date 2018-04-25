package backupschedule_test

import (
	. "github.com/grtl/mysql-operator/pkg/operator/backupschedule"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"

	apicorev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	versioned "github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Operator", func() {
	logrus.SetOutput(ioutil.Discard)

	var (
		operator         Operator
		backup           *crv1.MySQLBackupSchedule
		cluster          *crv1.MySQLCluster
		kubeClientset    *fake.Clientset
		clientset        *versioned.Clientset
		cronJobInterface batchv1.CronJobInterface
		pvcInterface     corev1.PersistentVolumeClaimInterface
	)

	BeforeEach(func() {
		clientset = versioned.NewSimpleClientset()
		kubeClientset = fake.NewSimpleClientset()

		operator = NewBackupScheduleOperator(clientset, kubeClientset)

		cronJobInterface = kubeClientset.BatchV1beta1().CronJobs(metav1.NamespaceDefault)
		pvcInterface = kubeClientset.CoreV1().PersistentVolumeClaims(metav1.NamespaceDefault)
	})

	When("a Backup Schedule is added", func() {
		BeforeEach(func() {
			cluster = new(crv1.MySQLCluster)
			err := factory.Build(testingFactory.MySQLClusterFactory).To(cluster)
			Expect(err).NotTo(HaveOccurred())

			backup = new(crv1.MySQLBackupSchedule)
			err = factory.Build(testingFactory.MySQLBackupScheduleFactory,
				factory.WithField("Spec.Cluster", cluster.Name),
				factory.WithTraits("ChangeDefaults")).To(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		JustBeforeEach(func() {
			_, err := clientset.CrV1().MySQLClusters(metav1.NamespaceDefault).Create(cluster)
			Expect(err).NotTo(HaveOccurred())
			err = operator.AddBackupSchedule(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates a PVC", func() {
			pvcs, err := pvcInterface.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(pvcs.Items).To(HaveLen(1))

			pvc := pvcs.Items[0]
			Expect(pvc.Name).To(Equal(PVCName(backup.Name)))
			Expect(pvc.Spec.Resources.Requests).To(Equal(apicorev1.ResourceList{
				"storage": backup.Spec.Storage,
			}))
		})

		It("creates a cron job", func() {
			cronJobs, err := cronJobInterface.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(cronJobs.Items).To(HaveLen(1))

			cronJob := cronJobs.Items[0]
			Expect(cronJob.Name).To(Equal(CronJobName(backup.Name)))
			Expect(cronJob.OwnerReferences[0].UID).To(Equal(backup.UID))
		})
	})

	When("a Backup Schedule without storage specified is added", func() {
		BeforeEach(func() {
			cluster = new(crv1.MySQLCluster)
			err := factory.Build(testingFactory.MySQLClusterFactory,
				factory.WithField("Spec.Storage", resource.MustParse("2Gi"))).To(cluster)
			Expect(err).NotTo(HaveOccurred())

			backup = new(crv1.MySQLBackupSchedule)
			err = factory.Build(testingFactory.MySQLBackupScheduleFactory,
				factory.WithField("Spec.Cluster", cluster.Name)).To(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		JustBeforeEach(func() {
			_, err := clientset.CrV1().MySQLClusters(metav1.NamespaceDefault).Create(cluster)
			Expect(err).NotTo(HaveOccurred())
			err = operator.AddBackupSchedule(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates a PVC with cluster storage value", func() {
			pvcs, err := pvcInterface.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(pvcs.Items).To(HaveLen(1))

			pvc := pvcs.Items[0]
			Expect(pvc.Name).To(Equal(PVCName(backup.Name)))
			Expect(pvc.Spec.Resources.Requests).To(Equal(apicorev1.ResourceList{
				"storage": cluster.Spec.Storage,
			}))
		})
	})
})
