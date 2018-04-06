package factory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/testing/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/nauyey/factory"
	"k8s.io/apimachinery/pkg/api/resource"
)

var _ = Describe("Backup", func() {
	When("ChangeDefaults Trait is not specified", func() {
		It("should generate backup with default values", func() {
			backup := new(crv1.MySQLBackup)
			err := factory.Build(MySQLBackupFactory).To(backup)
			Expect(err).NotTo(HaveOccurred())
			Expect(backup).NotTo(BeNil())
			Expect(backup.Name).To(ContainSubstring("backup"))
			Expect(backup.ObjectMeta.Namespace).To(Equal("default"))
			Expect(backup.Spec.Time).NotTo(BeNil())
			Expect(backup.Spec.Storage.IsZero()).To(BeTrue())
		})
	})

	When("ChangeDefaults Trait is specified", func() {
		It("should generate backup with default values changed", func() {
			backup := new(crv1.MySQLBackup)
			err := factory.Build(MySQLBackupFactory, factory.WithTraits("ChangeDefaults")).To(backup)
			Expect(err).NotTo(HaveOccurred())
			Expect(backup).NotTo(BeNil())
			Expect(backup.Name).To(ContainSubstring("backup"))
			Expect(backup.Namespace).To(Equal("default"))
			Expect(backup.Spec.Time).NotTo(BeNil())
			Expect(backup.Spec.Storage).To(Equal(resource.MustParse("1Gi")))
		})
	})
})
