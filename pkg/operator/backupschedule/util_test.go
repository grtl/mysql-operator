package backupschedule_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/operator/backupschedule"
)

var _ = Describe("Util", func() {
	const scheduleName = "my-backup"
	const anotherScheduleName = "another-backup"

	Describe("CronJobName", func() {
		It("should generate a name for the cron job", func() {
			Expect(CronJobName(scheduleName)).To(Equal("my-backup-job"))
			Expect(CronJobName(anotherScheduleName)).To(Equal("another-backup-job"))
		})
	})

	Describe("PVCName", func() {
		It("should generate a name for the pvc", func() {
			Expect(PVCName(scheduleName)).To(Equal("my-backup"))
			Expect(PVCName(anotherScheduleName)).To(Equal("another-backup"))
		})
	})
})
