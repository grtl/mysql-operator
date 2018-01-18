package operator

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1beta2client "k8s.io/client-go/kubernetes/types/apps/v1beta2"
)

const ContainerPortNumber = 1234

func AddCluster(cluster *crv1.MySQLCluster, corev1_client corev1.CoreV1Interface, v1beta2_client v1beta2client.AppsV1beta2Interface) {
	createServiceForCluster(cluster, corev1_client)
	createStatefulSetForCluster(cluster, v1beta2_client)
}

func createServiceForCluster(cluster *crv1.MySQLCluster, corev1_client corev1.CoreV1Interface) {
	servicesInterface := corev1_client.Services(cluster.ObjectMeta.Namespace)

	newService := serviceForCluster(cluster)
	_, err := servicesInterface.Create(&newService)

	if err != nil && !apierrors.IsAlreadyExists(err) {
		panic(err)
	}
}

func createStatefulSetForCluster(cluster *crv1.MySQLCluster, v1beta2_client v1beta2client.AppsV1beta2Interface) {
	statefulSetsInterface := v1beta2_client.StatefulSets(cluster.ObjectMeta.Namespace)

	newStatefulSet := statefulSetForCluster(cluster)
	_, err := statefulSetsInterface.Create(&newStatefulSet)

	if err != nil && !apierrors.IsAlreadyExists(err) {
		panic(err)
	}
}

func serviceForCluster(cluster *crv1.MySQLCluster) v1.Service {
	namespace := cluster.ObjectMeta.Namespace

	port := cluster.Spec.Port

	return v1.Service{
		TypeMeta: cluster.TypeMeta,
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.ObjectMeta.Name,
			Namespace: namespace,
			Labels:    cluster.ObjectMeta.Labels,
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Ports: []v1.ServicePort{
				v1.ServicePort{
					Port: int32(port),
					TargetPort: intstr.IntOrString{
						IntVal: ContainerPortNumber,
						Type:   intstr.Int,
					},
				},
			},
			Selector: map[string]string{
				"app": cluster.Spec.App,
			},
		},
	}
}

func serviceForCluster(cluster *crv1.MySQLCluster) v1beta2.StatefulSet {
	namespace := cluster.ObjectMeta.Namespace

	return v1beta2.StatefulSet{
		TypeMeta: cluster.TypeMeta,
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.ObjectMeta.Name,
			Namespace: namespace,
		},
		Spec: v1beta2.StatefulSetSpec{
			Replicas:    cluster.Spec.Replicas,
			ServiceName: cluster.ObjectMeta.Name,
			Ports: []v1.ServicePort{
				v1.ServicePort{
					Port: int32(port),
					TargetPort: intstr.IntOrString{
						IntVal: ContainerPortNumber,
						Type:   intstr.Int,
					},
				},
			},
			Selector: map[string]string{
				"app": cluster.Spec.App,
			},
		},
	}
}
