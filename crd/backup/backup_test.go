package backup_test

import (
	. "github.com/grtl/mysql-operator/crd/backup"
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
	const backupCRD = "mysqlbackups.cr.mysqloperator.grtl.github.com"

	var (
		clientset    *fake.Clientset
		crdInterface v1beta1.CustomResourceDefinitionInterface
	)

	BeforeEach(func() {
		clientset = fake.NewSimpleClientset()
		crdInterface = clientset.ApiextensionsV1beta1().CustomResourceDefinitions()
	})

	Describe("Create backup CRD when it is not registered", func() {
		It("should register backup CRD in the clientset", func(done Done) {
			go func() {
				defer GinkgoRecover()

				go func() {
					err := CreateBackupCRD(clientset)
					Expect(err).NotTo(HaveOccurred())
					crd, err := crdInterface.Get(backupCRD, metav1.GetOptions{})
					Expect(err).NotTo(HaveOccurred())
					Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackup"))
					close(done)
				}()

				// Manually update status in order for waitForCRDEstablish to succeed
				err := waitForCRDCreated(crdInterface, backupCRD)
				Expect(err).NotTo(HaveOccurred())
				err = updateCRDStatus(crdInterface, backupCRD)
				Expect(err).NotTo(HaveOccurred())
			}()
		})
	})

	Describe("Create backup CRD when it is already registered", func() {
		JustBeforeEach(func() {
			crd := new(apiextensionsv1.CustomResourceDefinition)
			err := util.ObjectFromFile("artifacts/backup-crd.yaml", crd)
			Expect(err).NotTo(HaveOccurred())
			_, err = crdInterface.Create(crd)
			Expect(err).NotTo(HaveOccurred())
			err = updateCRDStatus(crdInterface, backupCRD)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should finish with no error", func() {
			err := CreateBackupCRD(clientset)
			Expect(err).NotTo(HaveOccurred())
			crd, err := crdInterface.Get(backupCRD, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(crd.Spec.Names.Kind).To(Equal("MySQLBackup"))
		})
	})
})

// Waits until a CRD is registered in the clientset
func waitForCRDCreated(crdInterface v1beta1.CustomResourceDefinitionInterface, crdName string) error {
	return wait.Poll(50*time.Millisecond, 15*time.Second, func() (bool, error) {
		_, err := crdInterface.Get(crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		return true, err
	})
}

// Updates a CRD status to Established
func updateCRDStatus(crdInterface v1beta1.CustomResourceDefinitionInterface, crdName string) error {
	crd, err := crdInterface.Get(crdName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	crd.Status.Conditions = append(crd.Status.Conditions, apiextensionsv1.CustomResourceDefinitionCondition{
		Type:   apiextensionsv1.Established,
		Status: apiextensionsv1.ConditionTrue,
	})
	_, err = crdInterface.UpdateStatus(crd)
	return err
}
