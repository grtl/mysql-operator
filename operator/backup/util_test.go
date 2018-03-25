package backup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/operator/backup"
)

var _ = Describe("Util", func() {
	const backupName = "my-backup"
	const anotherBackupName = "another-backup"

	Describe("CronJobName", func() {
		It("should generate a name for the cron job", func() {
			Expect(CronJobName(backupName)).To(Equal("my-backup-job"))
			Expect(CronJobName(anotherBackupName)).To(Equal("another-backup-job"))
		})
	})

	Describe("PVCName", func() {
		It("should generate a name for the pvc", func() {
			Expect(PVCName(backupName)).To(Equal("my-backup"))
			Expect(PVCName(anotherBackupName)).To(Equal("another-backup"))
		})
	})
})
