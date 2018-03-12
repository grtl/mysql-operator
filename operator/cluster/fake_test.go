package cluster_test

import (
	. "github.com/grtl/mysql-operator/operator/cluster"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"github.com/nauyey/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

var _ = Describe("Fake cluster operator", func() {
	var (
		operator *FakeClusterOperator
		cluster  *crv1.MySQLCluster
	)

	BeforeEach(func() {
		operator = NewFakeOperator()
		cluster = new(crv1.MySQLCluster)
		err := factory.Build(testingFactory.MySQLClusterFactory).To(cluster)
		Expect(err).NotTo(HaveOccurred())
	})

	When("No error is set", func() {
		It("should return without error", func() {
			err := operator.AddCluster(cluster)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	When("An error is set", func() {
		It("should return the error", func() {
			expectedErr := fmt.Errorf("Expected error")
			operator.SetError(expectedErr)
			err := operator.AddCluster(cluster)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(expectedErr))
		})
	})
})
