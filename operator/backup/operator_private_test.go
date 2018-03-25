package backup

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"

	"github.com/nauyey/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

var _ = Describe("Operator", func() {
	var backup *crv1.MySQLBackup

	BeforeEach(func() {
		backup = new(crv1.MySQLBackup)
		err := factory.Build(testingFactory.MySQLBackupFactory).To(backup)
		Expect(err).NotTo(HaveOccurred())
		backup.Name = "my-backup"
	})

	Describe("pvcForBackup should generate a pvc from the template", func() {
		var pvc *corev1.PersistentVolumeClaim

		BeforeEach(func() {
			var err error
			pvc, err = pvcForBackup(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(pvc.Name).To(Equal("my-backup"))
		})
	})

	Describe("cronJobForBackup should generate a cronjob from the template", func() {
		var cronJob *v1beta1.CronJob

		BeforeEach(func() {
			var err error
			cronJob, err = cronJobForBackup(backup)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(cronJob.Name).To(Equal("my-backup-job"))
		})

		It("should have the backup as the owner", func() {
			Expect(cronJob.OwnerReferences).To(HaveLen(1))
			Expect(cronJob.OwnerReferences[0].Kind).To(Equal("MySQLBackup"))
			Expect(cronJob.OwnerReferences[0].Name).To(Equal("my-backup"))
		})
	})
})
