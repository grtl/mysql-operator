package cluster

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/nauyey/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Operator", func() {
	var cluster *crv1.MySQLCluster

	BeforeEach(func() {
		cluster = new(crv1.MySQLCluster)
		err := factory.Build(testingFactory.MySQLClusterFactory).To(cluster)
		Expect(err).NotTo(HaveOccurred())
		cluster.Name = "my-cluster"
	})

	Describe("statefulSetForCluster should generate a stateful set from the template", func() {
		var statefulSet *appsv1.StatefulSet

		BeforeEach(func() {
			var err error
			statefulSet, err = statefulSetForCluster(cluster, nil)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(statefulSet.Name).To(Equal("my-cluster"))
		})

		It("should have the cluster as the owner", func() {
			Expect(statefulSet.OwnerReferences).To(HaveLen(1))
			Expect(statefulSet.OwnerReferences[0].Kind).To(Equal("MySQLCluster"))
			Expect(statefulSet.OwnerReferences[0].Name).To(Equal("my-cluster"))
		})
	})

	Describe("serviceForCluster should generate a service from the template", func() {
		var service *corev1.Service

		BeforeEach(func() {
			var err error
			service, err = serviceForCluster(cluster, serviceTemplate)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(service.Name).To(Equal("my-cluster"))
		})

		It("should have the cluster as the owner", func() {
			Expect(service.OwnerReferences).To(HaveLen(1))
			Expect(service.OwnerReferences[0].Kind).To(Equal("MySQLCluster"))
			Expect(service.OwnerReferences[0].Name).To(Equal("my-cluster"))
		})
	})

	Describe("serviceForCluster should generate a read service from the template", func() {
		var readService *corev1.Service

		BeforeEach(func() {
			var err error
			readService, err = serviceForCluster(cluster, serviceReadTemplate)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should have a name", func() {
			Expect(readService.Name).To(Equal("my-cluster-read"))
		})

		It("should have the cluster as the owner", func() {
			Expect(readService.OwnerReferences).To(HaveLen(1))
			Expect(readService.OwnerReferences[0].Kind).To(Equal("MySQLCluster"))
			Expect(readService.OwnerReferences[0].Name).To(Equal("my-cluster"))
		})
	})
})
