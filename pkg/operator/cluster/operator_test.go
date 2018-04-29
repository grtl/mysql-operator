package cluster_test

import (
	. "github.com/grtl/mysql-operator/pkg/operator/cluster"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"

	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	versioned "github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
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
		err := factory.Build(testingFactory.MySQLClusterFactory,
			factory.WithTraits("ChangeDefaults"),
		).To(cluster)
		Expect(err).NotTo(HaveOccurred())

		clientset = versioned.NewSimpleClientset()
		kubeClientset = fake.NewSimpleClientset()

		clusters := clientset.CrV1().MySQLClusters("default")
		_, err = clusters.Create(cluster)
		Expect(err).NotTo(HaveOccurred())

		operator = NewClusterOperator(clientset, kubeClientset)

		services = kubeClientset.CoreV1().Services(metav1.NamespaceDefault)
		statefulSets = kubeClientset.AppsV1().StatefulSets(metav1.NamespaceDefault)
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

		It("creates the appropriate read-write service", func() {
			name := cluster.Name
			svcs, err := services.List(metav1.ListOptions{})

			Expect(err).NotTo(HaveOccurred())

			var readSvc *apicorev1.Service = nil
			for _, svc := range svcs.Items {
				if svc.Name == ServiceName(name) {
					readSvc = &svc
				}
			}

			Expect(readSvc).NotTo(BeNil())
			Expect(readSvc.OwnerReferences[0].UID).To(Equal(cluster.UID))
			Expect(readSvc.Spec.Ports[0].Port).To(Equal(cluster.Spec.Port))
			Expect(readSvc.Spec.Selector["app"]).To(Equal(name))
		})

		It("creates the appropriate read service", func() {
			name := cluster.Name
			svcs, err := services.List(metav1.ListOptions{})

			Expect(err).NotTo(HaveOccurred())

			var readSvc *apicorev1.Service = nil
			for _, svc := range svcs.Items {
				if svc.Name == ReadServiceName(name) {
					readSvc = &svc
				}
			}

			Expect(readSvc).NotTo(BeNil())
			Expect(readSvc.OwnerReferences[0].UID).To(Equal(cluster.UID))
			Expect(readSvc.Spec.Ports[0].Port).To(Equal(cluster.Spec.Port))
			Expect(readSvc.Spec.Selector["app"]).To(Equal(name))
		})

		It("creates the appropriate StatefulSet", func() {
			name := cluster.Name
			sets, err := statefulSets.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(sets.Items).To(HaveLen(1))

			sts := sets.Items[0]

			Expect(sts.Name).To(Equal(StatefulSetName(name)))
			Expect(sts.OwnerReferences[0].UID).To(Equal(cluster.UID))
			Expect(sts.Spec.Selector.MatchLabels).To(Equal(map[string]string{
				"app": name,
			}))
			Expect(*sts.Spec.Replicas).To(Equal(cluster.Spec.Replicas))
			Expect(sts.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests).To(
				Equal(apicorev1.ResourceList{
					"storage": cluster.Spec.Storage,
				}))
		})
	})

	When("a cluster is updated", func() {
		var updatedCluster *crv1.MySQLCluster

		BeforeEach(func() {
			updatedCluster = new(crv1.MySQLCluster)
			updatedCluster = cluster.DeepCopy()
			updatedCluster.Spec.Port = updatedCluster.Spec.Port + 1
			updatedCluster.Spec.Replicas = updatedCluster.Spec.Replicas + 1
		})

		JustBeforeEach(func() {
			err := operator.UpdateCluster(updatedCluster)
			Expect(err).NotTo(HaveOccurred())
		})

		It("updates the StatefulSet", func() {
			sets, err := statefulSets.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(sets.Items).To(HaveLen(1))
			sts := sets.Items[0]
			Expect(*sts.Spec.Replicas).To(Equal(cluster.Spec.Replicas + 1))
		})

		It("updates the StatefulSet", func() {
			sets, err := statefulSets.List(metav1.ListOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(sets.Items).To(HaveLen(1))
			sts := sets.Items[0]
			Expect(*sts.Spec.Replicas).To(Equal(cluster.Spec.Replicas + 1))
		})

		It("updates the Services", func() {
			svcs, err := services.List(metav1.ListOptions{})

			Expect(err).NotTo(HaveOccurred())

			for _, svc := range svcs.Items {
				Expect(svc.Spec.Ports[0].Port).To(Equal(cluster.Spec.Port + 1))
			}
		})
	})
})
