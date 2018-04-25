package factory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/testing/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/nauyey/factory"
	"k8s.io/apimachinery/pkg/api/resource"
)

var _ = Describe("Factory Backup Schedule", func() {
	When("ChangeDefaults Trait is not specified", func() {
		It("should generate Backup Schedule with default values", func() {
			schedule := new(crv1.MySQLBackupSchedule)
			err := factory.Build(MySQLBackupScheduleFactory).To(schedule)
			Expect(err).NotTo(HaveOccurred())
			Expect(schedule).NotTo(BeNil())
			Expect(schedule.Name).To(ContainSubstring("backup"))
			Expect(schedule.ObjectMeta.Namespace).To(Equal("default"))
			Expect(schedule.Spec.Time).NotTo(BeNil())
			Expect(schedule.Spec.Storage.IsZero()).To(BeTrue())
		})
	})

	When("ChangeDefaults Trait is specified", func() {
		It("should generate Backup Schedule with default values changed", func() {
			schedule := new(crv1.MySQLBackupSchedule)
			err := factory.Build(MySQLBackupScheduleFactory, factory.WithTraits("ChangeDefaults")).To(schedule)
			Expect(err).NotTo(HaveOccurred())
			Expect(schedule).NotTo(BeNil())
			Expect(schedule.Name).To(ContainSubstring("backup"))
			Expect(schedule.Namespace).To(Equal("default"))
			Expect(schedule.Spec.Time).NotTo(BeNil())
			Expect(schedule.Spec.Storage).To(Equal(resource.MustParse("1Gi")))
		})
	})
})
