package startup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/grtl/mysql-operator/pkg/crd/backupinstance"
	"github.com/grtl/mysql-operator/pkg/crd/backupschedule"
	"github.com/grtl/mysql-operator/pkg/crd/cluster"
)

var _ = Describe("On operator startup", func() {
	It("should register MySQLCluster crd", func() {
		crdInterface := operator.ExtClientset().ApiextensionsV1beta1().CustomResourceDefinitions()
		crd, err := crdInterface.Get(cluster.CustomResourceName, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(crd.Spec.Names.Kind).To(Equal("MySQLCluster"))
	})

	It("should register MySQLBackupSchedule crd", func() {
		crdInterface := operator.ExtClientset().ApiextensionsV1beta1().CustomResourceDefinitions()
		crd, err := crdInterface.Get(backupschedule.CustomResourceName, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackupSchedule"))
	})

	It("should register MySQLBackupInstance crd", func() {
		crdInterface := operator.ExtClientset().ApiextensionsV1beta1().CustomResourceDefinitions()
		crd, err := crdInterface.Get(backupinstance.CustomResourceName, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackupInstance"))
	})
})
