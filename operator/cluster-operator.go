package operator

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

const ContainerPortNumber = 1234

func AddCluster(cluster *crv1.MySQLCluster, corev1_client corev1.CoreV1Interface) {
	createServiceForCluster(cluster, corev1_client)
}

func createServiceForCluster(cluster *crv1.MySQLCluster, corev1_client corev1.CoreV1Interface) {
	servicesInterface := corev1_client.Services(cluster.ObjectMeta.Namespace)

	newService := serviceForCluster(cluster)
	_, err := servicesInterface.Create(&newService)

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
