package backupinstance_test

import (
	. "github.com/grtl/mysql-operator/pkg/operator/backupinstance"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	batchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	versioned "github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Operator", func() {
	logrus.SetOutput(ioutil.Discard)

	var (
		operator Operator

		schedule *crv1.MySQLBackupSchedule
		backup   *crv1.MySQLBackupInstance
		cluster  *crv1.MySQLCluster

		kubeClientset *fake.Clientset
		clientset     *versioned.Clientset
		jobInterface  batchv1.JobInterface
	)

	BeforeEach(func() {
		clientset = versioned.NewSimpleClientset()
		kubeClientset = fake.NewSimpleClientset()

		operator = NewBackupInstanceOperator(clientset, kubeClientset)

		jobInterface = kubeClientset.BatchV1().Jobs(metav1.NamespaceDefault)

		cluster = new(crv1.MySQLCluster)
		err := factory.Build(testingFactory.MySQLClusterFactory,
			factory.WithField("ObjectMeta.Namespace", metav1.NamespaceDefault)).To(cluster)
		Expect(err).NotTo(HaveOccurred())

		schedule = new(crv1.MySQLBackupSchedule)
		err = factory.Build(testingFactory.MySQLBackupScheduleFactory,
			factory.WithField("ObjectMeta.Namespace", metav1.NamespaceDefault),
			factory.WithField("Spec.Cluster", cluster.Name),
			factory.WithTraits("ChangeDefaults")).To(schedule)
		Expect(err).NotTo(HaveOccurred())
	})

	When("a Backup Instance is added", func() {
		Describe("within an existing schedule", func() {
			BeforeEach(func() {
				backup = new(crv1.MySQLBackupInstance)
				err := factory.Build(testingFactory.MySQLBackupInstanceFactory,
					factory.WithField("ObjectMeta.Namespace", metav1.NamespaceDefault),
					factory.WithField("Spec.Schedule", schedule.Name)).To(backup)
				Expect(err).NotTo(HaveOccurred())
			})

			JustBeforeEach(func() {
				_, err := clientset.CrV1().MySQLClusters(cluster.Namespace).Create(cluster)
				Expect(err).NotTo(HaveOccurred())
				_, err = clientset.CrV1().MySQLBackupSchedules(schedule.Namespace).Create(schedule)
				Expect(err).NotTo(HaveOccurred())
				err = operator.CreateBackup(backup)
				Expect(err).NotTo(HaveOccurred())
			})

			It("creates a Create Job", func() {
				jobs, err := jobInterface.List(metav1.ListOptions{})
				Expect(err).NotTo(HaveOccurred())
				Expect(jobs.Items).To(HaveLen(1))

				job := jobs.Items[0]
				Expect(job.Name).To(Equal(JobCreateName(backup.Name)))
			})
		})

		Describe("without an existing schedule", func() {
			BeforeEach(func() {
				backup = new(crv1.MySQLBackupInstance)
				err := factory.Build(testingFactory.MySQLBackupInstanceFactory,
					factory.WithField("ObjectMeta.Namespace", metav1.NamespaceDefault),
					factory.WithField("Spec.Schedule", schedule.Name)).To(backup)
				Expect(err).NotTo(HaveOccurred())
			})

			It("fails with an error", func() {
				err := operator.CreateBackup(backup)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	When("a Backup Instance is deleted", func() {
		BeforeEach(func() {
			backup = new(crv1.MySQLBackupInstance)
			err := factory.Build(testingFactory.MySQLBackupInstanceFactory,
				factory.WithField("ObjectMeta.Namespace", metav1.NamespaceDefault),
				factory.WithField("Spec.Schedule", schedule.Name)).To(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		JustBeforeEach(func() {
			_, err := clientset.CrV1().MySQLClusters(metav1.NamespaceDefault).Create(cluster)
			Expect(err).NotTo(HaveOccurred())
			_, err = clientset.CrV1().MySQLBackupSchedules(metav1.NamespaceDefault).Create(schedule)
			Expect(err).NotTo(HaveOccurred())
			_, err = clientset.CrV1().MySQLBackupInstances(metav1.NamespaceDefault).Create(backup)
			Expect(err).NotTo(HaveOccurred())

			err = operator.DeleteBackup(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates a Delete job", func() {
			jobs, err := jobInterface.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(jobs.Items).To(HaveLen(1))

			job := jobs.Items[0]
			Expect(job.Name).To(Equal(JobDeleteName(backup.Name)))
		})
	})
})
