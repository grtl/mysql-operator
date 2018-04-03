package factory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/testing/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/nauyey/factory"
)

var _ = Describe("Cluster", func() {
	When("ChangeDefaults Trait is not specified", func() {
		It("should generate a cluster", func() {
			cluster := new(crv1.MySQLCluster)
			err := factory.Build(MySQLClusterFactory).To(cluster)
			Expect(err).NotTo(HaveOccurred())
			Expect(cluster).NotTo(BeNil())
			Expect(cluster.Name).To(ContainSubstring("cluster"))
			Expect(cluster.Namespace).To(Equal("default"))
			Expect(cluster.Spec.Secret).NotTo(BeEmpty())
			Expect(cluster.Spec.Storage).NotTo(BeNil())
		})

		It("should generate a cluster with default values", func() {
			cluster := new(crv1.MySQLCluster)
			err := factory.Build(MySQLClusterFactory).To(cluster)
			Expect(err).NotTo(HaveOccurred())
			Expect(cluster).NotTo(BeNil())
			cluster.WithDefaults()
			Expect(cluster.Spec.Replicas).To(Equal(uint32(2)))
			Expect(cluster.Spec.Port).To(Equal(int32(3306)))
			Expect(cluster.Spec.Image).To(Equal("mysql:latest"))
		})
	})

	When("ChangeDefaults Trait is specified", func() {
		It("should generate a cluster", func() {
			cluster := new(crv1.MySQLCluster)
			err := factory.Build(MySQLClusterFactory, factory.WithTraits("ChangeDefaults")).To(cluster)
			Expect(err).NotTo(HaveOccurred())
			Expect(cluster).NotTo(BeNil())
			Expect(cluster.Name).To(ContainSubstring("cluster"))
			Expect(cluster.Namespace).To(Equal("default"))
			Expect(cluster.Spec.Secret).NotTo(BeEmpty())
			Expect(cluster.Spec.Storage).NotTo(BeNil())
		})

		It("should generate values for fields with default values", func() {
			cluster := new(crv1.MySQLCluster)
			err := factory.Build(MySQLClusterFactory, factory.WithTraits("ChangeDefaults")).To(cluster)
			Expect(err).NotTo(HaveOccurred())
			Expect(cluster).NotTo(BeNil())
			cluster.WithDefaults()
			Expect(cluster.Spec.Replicas).NotTo(Equal(uint32(2)))
			Expect(cluster.Spec.Port).NotTo(Equal(int32(3306)))
			Expect(cluster.Spec.Image).NotTo(Equal("mysql:latest"))
		})
	})
})
