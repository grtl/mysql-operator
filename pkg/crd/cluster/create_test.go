package cluster_test

import (
	. "github.com/grtl/mysql-operator/pkg/crd/cluster"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/grtl/mysql-operator/pkg/util"
)

var _ = Describe("CRD Cluster Create", func() {
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
				err := CreateClusterCRD(corev1.NamespaceAll, clientset)
				Expect(err).NotTo(HaveOccurred())
				crd, err := crdInterface.Get(CustomResourceName, metav1.GetOptions{})
				Expect(err).NotTo(HaveOccurred())
				Expect(crd.Spec.Names.Kind).To(Equal("MySQLCluster"))
				close(done)
			}()

			// Manually update status in order for waitForCRDEstablish to succeed
			err := waitForCRDCreated(crdInterface, CustomResourceName)
			Expect(err).NotTo(HaveOccurred())
			err = updateCRDStatus(crdInterface, CustomResourceName)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Create cluster CRD when it is already registered", func() {
		JustBeforeEach(func() {
			crd := new(apiextensionsv1.CustomResourceDefinition)
			err := util.ObjectFromFile("artifacts/cluster-crd.yaml", crd)
			Expect(err).NotTo(HaveOccurred())
			_, err = crdInterface.Create(crd)
			Expect(err).NotTo(HaveOccurred())
			err = updateCRDStatus(crdInterface, CustomResourceName)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should finish with no fail", func() {
			err := CreateClusterCRD(corev1.NamespaceAll, clientset)
			Expect(err).NotTo(HaveOccurred())
			crd, err := crdInterface.Get(CustomResourceName, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(crd.Spec.Names.Kind).To(Equal("MySQLCluster"))
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
