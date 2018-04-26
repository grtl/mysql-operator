package backupschedule

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"

	"github.com/nauyey/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Operator Backup Schedule Private", func() {
	var schedule *crv1.MySQLBackupSchedule

	BeforeEach(func() {
		schedule = new(crv1.MySQLBackupSchedule)
		err := factory.Build(testingFactory.MySQLBackupScheduleFactory).To(schedule)
		Expect(err).NotTo(HaveOccurred())
		schedule.Name = "my-schedule"
	})

	Describe("pvcForSchedule should generate a pvc from the template", func() {
		var pvc *corev1.PersistentVolumeClaim

		BeforeEach(func() {
			var err error
			pvc, err = pvcForSchedule(schedule)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(pvc.Name).To(Equal("my-schedule"))
		})
	})

	Describe("cronJobForSchedule should generate a cronjob from the template", func() {
		var cronJob *v1beta1.CronJob

		BeforeEach(func() {
			var err error
			cronJob, err = cronJobForSchedule(schedule)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(cronJob.Name).To(Equal("my-schedule-job"))
		})

		It("should have the Backup Schedule as the owner", func() {
			Expect(cronJob.OwnerReferences).To(HaveLen(1))
			Expect(cronJob.OwnerReferences[0].Kind).To(Equal("MySQLBackupSchedule"))
			Expect(cronJob.OwnerReferences[0].Name).To(Equal("my-schedule"))
		})
	})
})
