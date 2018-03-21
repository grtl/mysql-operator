package startup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("On operator startup", func() {
	It("should register MySQLCluster crd", func() {
		const clusterCRD = "mysqlclusters.cr.mysqloperator.grtl.github.com"
		crd, err := operator.ExtClientset().ApiextensionsV1beta1().
			CustomResourceDefinitions().Get(clusterCRD, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(crd.Spec.Names.Kind).To(Equal("MySQLCluster"))
	})

	It("should register MySQLBackup crd", func() {
		const backupCRD = "mysqlbackups.cr.mysqloperator.grtl.github.com"
		crd, err := operator.ExtClientset().ApiextensionsV1beta1().
			CustomResourceDefinitions().Get(backupCRD, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackup"))
	})
})
