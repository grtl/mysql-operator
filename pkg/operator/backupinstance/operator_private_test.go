package backupinstance

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	batchv1 "k8s.io/api/batch/v1"

	"github.com/nauyey/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Operator Backup Instance Private", func() {
	var backup *crv1.MySQLBackupInstance

	BeforeEach(func() {
		backup = new(crv1.MySQLBackupInstance)
		err := factory.Build(testingFactory.MySQLBackupInstanceFactory).To(backup)
		Expect(err).NotTo(HaveOccurred())
		backup.Name = "my-backup-2018-04-27-01-21-03"
	})

	Describe("jobForBackup with the createJob template should generate a Create job", func() {
		var job *batchv1.Job

		BeforeEach(func() {
			var err error
			job, err = jobForBackup(backup, jobCreateTemplate)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(job.Name).To(Equal("my-backup-2018-04-27-01-21-03-create"))
		})
	})

	Describe("jobForBackup with the deleteJob template should generate a Delete job", func() {
		var job *batchv1.Job

		BeforeEach(func() {
			var err error
			job, err = jobForBackup(backup, jobDeleteTemplate)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(job.Name).To(Equal("my-backup-2018-04-27-01-21-03-delete"))
		})
	})
})
