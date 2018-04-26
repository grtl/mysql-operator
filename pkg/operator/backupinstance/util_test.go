package backupinstance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/operator/backupinstance"
)

var _ = Describe("Util", func() {
	const scheduleName = "my-backup-2018-27-04-15-20-00"
	const anotherScheduleName = "another-backup-2018-27-04-10-31-03"

	Describe("CronJobName", func() {
		It("should generate a name for the cron job", func() {
			Expect(JobCreateName(scheduleName)).To(Equal("my-backup-2018-27-04-15-20-00-create"))
			Expect(JobCreateName(anotherScheduleName)).To(Equal("another-backup-2018-27-04-10-31-03-create"))
		})
	})

	Describe("PVCName", func() {
		It("should generate a name for the pvc", func() {
			Expect(JobDeleteName(scheduleName)).To(Equal("my-backup-2018-27-04-15-20-00-delete"))
			Expect(JobDeleteName(anotherScheduleName)).To(Equal("another-backup-2018-27-04-10-31-03-delete"))
		})
	})
})
