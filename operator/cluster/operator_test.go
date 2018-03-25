package cluster_test

import (
	. "github.com/grtl/mysql-operator/operator/cluster"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	versioned "github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

var _ = Describe("Cluster Operator", func() {
	logrus.SetOutput(ioutil.Discard)

	var (
		operator      Operator
		cluster       *crv1.MySQLCluster
		kubeClientset *fake.Clientset
		clientset     *versioned.Clientset
		services      corev1.ServiceInterface
		statefulSets  appsv1.StatefulSetInterface
	)

	BeforeEach(func() {
		cluster = new(crv1.MySQLCluster)
		err := factory.Build(testingFactory.MySQLClusterFactory).To(cluster)
		Expect(err).NotTo(HaveOccurred())

		clientset = versioned.NewSimpleClientset()
		kubeClientset = fake.NewSimpleClientset()

		operator = NewClusterOperator(clientset, kubeClientset)

		services = kubeClientset.CoreV1().Services("default")
		statefulSets = kubeClientset.AppsV1().StatefulSets("default")
	})

	JustBeforeEach(func() {
		err := operator.AddCluster(cluster)
		Expect(err).NotTo(HaveOccurred())
	})

	When("a cluster is added", func() {
		It("creates two Services", func() {
			svcs, err := services.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(svcs.Items).To(HaveLen(2))
		})

		It("creates the appropriate StatefulSet", func() {
			sets, err := statefulSets.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(sets.Items).To(HaveLen(1))
		})
	})
})
