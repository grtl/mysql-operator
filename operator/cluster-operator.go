package operator

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	v1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

const MysqlPortNumber = 3306

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
			Name:      cluster.Spec.ServiceName,
			Namespace: namespace,
			Labels:    cluster.ObjectMeta.Labels,
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Ports: []v1.ServicePort{
				v1.ServicePort{
					Port: int32(port),
					TargetPort: intstr.IntOrString{
						IntVal: MysqlPortNumber,
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

func statefulSetForCluster(cluster *crv1.MySQLCluster) v1beta2.StatefulSet {
	namespace := cluster.ObjectMeta.Namespace

	labels := map[string]string{
		"app": cluster.Spec.App,
	}

	var replicas int32 = 1

	selector := metav1.LabelSelector{
		MatchLabels: labels,
	}

	envVars := []v1.EnvVar{
		v1.EnvVar{
			Name:  "MYSQL_ROOT_PASSWORD",
			Value: cluster.Spec.Password,
		},
	}

	container := v1.Container{
		Name:  cluster.ObjectMeta.Name,
		Image: "mysql:8",
		Env:   envVars,
		VolumeMounts: []v1.VolumeMount{
			v1.VolumeMount{
				Name:      "mysql",
				MountPath: "/var/lib/mysql",
			},
		},
	}

	pvc := v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mysql",
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				"ReadWriteOnce",
			},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					"storage": cluster.Spec.Storage,
				},
			},
		},
	}

	sts_spec := v1beta2.StatefulSetSpec{
		Replicas:    &replicas,
		ServiceName: cluster.Spec.ServiceName,
		Selector:    &selector,
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: labels,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					container,
				},
			},
		},
		VolumeClaimTemplates: []v1.PersistentVolumeClaim{
			pvc,
		},
	}

	return v1beta2.StatefulSet{
		TypeMeta: cluster.TypeMeta,
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.ObjectMeta.Name,
			Namespace: namespace,
		},
		Spec: sts_spec,
	}
}
