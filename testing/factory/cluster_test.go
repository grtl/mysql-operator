package factory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/testing/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/nauyey/factory"
)

var _ = Describe("Cluster", func() {
	It("should generate cluster", func() {
		cluster := new(crv1.MySQLCluster)
		err := factory.Build(MySQLClusterFactory).To(cluster)
		Expect(err).NotTo(HaveOccurred())
		Expect(cluster).NotTo(BeNil())
		Expect(cluster.Name).To(ContainSubstring("cluster"))
		Expect(cluster.Namespace).To(Equal("default"))
		Expect(cluster.Spec.Password).NotTo(BeEmpty())
		Expect(cluster.Spec.Storage).NotTo(BeNil())
	})
})
