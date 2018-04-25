package factory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/testing/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/nauyey/factory"
)

var _ = Describe("Factory Backup Instance", func() {

	It("should generate Backup Schedule with default values changed", func() {
		backup := new(crv1.MySQLBackupInstance)
		err := factory.Build(MySQLBackupInstanceFactory).To(backup)
		Expect(err).NotTo(HaveOccurred())
		Expect(backup).NotTo(BeNil())
		Expect(backup.Spec.Schedule).To(ContainSubstring("backup"))
		Expect(backup.Name).To(ContainSubstring(backup.Spec.Schedule))
		Expect(backup.Namespace).To(Equal("default"))
	})
})
