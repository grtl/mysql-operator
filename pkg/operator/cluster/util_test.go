package cluster_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/operator/cluster"
)

var _ = Describe("Util", func() {
	const clusterName = "my-cluster"
	const anotherClusterName = "another-cluster"

	Describe("StatefulSetName", func() {
		It("should generate a name for the stateful set", func() {
			Expect(StatefulSetName(clusterName)).To(Equal("my-cluster"))
			Expect(StatefulSetName(anotherClusterName)).To(Equal("another-cluster"))
		})
	})

	Describe("ServiceName", func() {
		It("should generate a name for the service", func() {
			Expect(ServiceName(clusterName)).To(Equal("my-cluster"))
			Expect(ServiceName(anotherClusterName)).To(Equal("another-cluster"))
		})
	})

	Describe("ReadServiceName", func() {
		It("should generate a name for the read service", func() {
			Expect(ReadServiceName(clusterName)).To(Equal("my-cluster-read"))
			Expect(ReadServiceName(anotherClusterName)).To(Equal("another-cluster-read"))
		})
	})
})
