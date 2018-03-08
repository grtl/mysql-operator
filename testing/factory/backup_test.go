package factory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/testing/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/nauyey/factory"
)

var _ = Describe("Backup", func() {
	It("should generate backup", func() {
		backup := new(crv1.MySQLBackup)
		err := factory.Build(MySQLBackupFactory).To(backup)
		Expect(err).NotTo(HaveOccurred())
		Expect(backup).NotTo(BeNil())
		Expect(backup.Name).To(ContainSubstring("backup"))
		Expect(backup.Namespace).To(Equal("default"))
	})
})
