package crd_test

import (
	. "github.com/grtl/mysql-operator/crd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/grtl/mysql-operator/util"
)

var _ = Describe("MySQL CRDs", func() {
	const (
		clusterCRD = "mysqlclusters.cr.mysqloperator.grtl.github.com"
		backupCRD  = "mysqlbackups.cr.mysqloperator.grtl.github.com"
	)

	var (
		clientset    *fake.Clientset
		crdInterface v1beta1.CustomResourceDefinitionInterface
	)

	BeforeEach(func() {
		clientset = fake.NewSimpleClientset()
		crdInterface = clientset.ApiextensionsV1beta1().CustomResourceDefinitions()
	})

	Describe("Create cluster CRD when it is not registered", func() {
		It("should register cluster CRD in the clientset", func(done Done) {
			go func() {
				defer GinkgoRecover()

				go func() {
					err := CreateClusterCRD(clientset)
					Expect(err).NotTo(HaveOccurred())
					crd, err := crdInterface.Get(clusterCRD, metav1.GetOptions{})
					Expect(err).NotTo(HaveOccurred())
					Expect(crd.Spec.Names.Kind).To(Equal("MySQLCluster"))
					close(done)
				}()

				// Manually update status in order for waitForCRDEstablish to succeed
				err := waitForCRDCreated(clientset, clusterCRD)
				Expect(err).NotTo(HaveOccurred())
				err = updateCRDStatus(clientset, clusterCRD)
				Expect(err).NotTo(HaveOccurred())
			}()
		})
	})

	Describe("Create cluster CRD when it is already registered", func() {
		JustBeforeEach(func() {
			crd := new(apiextensionsv1.CustomResourceDefinition)
			err := util.ObjectFromFile("artifacts/mysql-crd.yaml", crd)
			Expect(err).NotTo(HaveOccurred())
			_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should finish with no error", func() {
			err := CreateClusterCRD(clientset)
			Expect(err).NotTo(HaveOccurred())
			crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(clusterCRD, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(crd.Spec.Names.Kind).To(Equal("MySQLCluster"))
		})
	})

	Describe("Create backup CRD when it is not registered", func() {
		It("should register backup CRD in the clientset", func(done Done) {
			go func() {
				defer GinkgoRecover()

				go func() {
					err := CreateBackupCRD(clientset)
					Expect(err).NotTo(HaveOccurred())
					crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(backupCRD, metav1.GetOptions{})
					Expect(err).NotTo(HaveOccurred())
					Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackup"))
					close(done)
				}()

				// Manually update status in order for waitForCRDEstablish to succeed
				err := waitForCRDCreated(clientset, backupCRD)
				Expect(err).NotTo(HaveOccurred())
				err = updateCRDStatus(clientset, backupCRD)
				Expect(err).NotTo(HaveOccurred())
			}()
		})
	})

	Describe("Create backup CRD when it is already registered", func() {
		JustBeforeEach(func() {
			crd := new(apiextensionsv1.CustomResourceDefinition)
			err := util.ObjectFromFile("artifacts/backup-crd.yaml", crd)
			Expect(err).NotTo(HaveOccurred())
			_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should finish with no error", func() {
			err := CreateBackupCRD(clientset)
			Expect(err).NotTo(HaveOccurred())
			crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(backupCRD, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackup"))
		})
	})
})

// Waits until a CRD is registered in the clientset
func waitForCRDCreated(clientset *fake.Clientset, CRDName string) error {
	return wait.Poll(50*time.Millisecond, 5*time.Second, func() (bool, error) {
		_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CRDName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		return true, err
	})
}

// Updates a CRD status to Established
func updateCRDStatus(clientset *fake.Clientset, CRDName string) error {
	crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CRDName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	crd.Status.Conditions = append(crd.Status.Conditions, apiextensionsv1.CustomResourceDefinitionCondition{
		Type:   apiextensionsv1.Established,
		Status: apiextensionsv1.ConditionTrue,
	})
	_, err = clientset.ApiextensionsV1beta1().CustomResourceDefinitions().UpdateStatus(crd)
	return err
}
