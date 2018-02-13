package operator

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

const mySQLPortNumber = 3306

func AddCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) {
	createServiceForCluster(cluster, kubeClientset)
	createStatefulSetForCluster(cluster, kubeClientset)
}

func createServiceForCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) {
	servicesInterface := kubeClientset.CoreV1().Services(cluster.ObjectMeta.Namespace)

	newService := serviceForCluster(cluster)
	_, err := servicesInterface.Create(&newService)

	if err != nil && !apierrors.IsAlreadyExists(err) {
		panic(err)
	}
}

func createStatefulSetForCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) {
	statefulSetsInterface := kubeClientset.AppsV1beta2().StatefulSets(cluster.ObjectMeta.Namespace)

	newStatefulSet := statefulSetForCluster(cluster)

	_, err := statefulSetsInterface.Create(&newStatefulSet)

	if err != nil && !apierrors.IsAlreadyExists(err) {
		panic(err)
	}
}

func serviceForCluster(cluster *crv1.MySQLCluster) v1.Service {
	return v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: cluster.Spec.Name,
		},
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Ports: []v1.ServicePort{
				v1.ServicePort{
					Port: mySQLPortNumber,
				},
			},
			Selector: map[string]string{
				"mysql-cluster": cluster.Spec.Name,
			},
		},
	}
}

func statefulSetForCluster(cluster *crv1.MySQLCluster) v1beta2.StatefulSet {
	namespace := cluster.ObjectMeta.Namespace

	labels := map[string]string{
		"mysql-cluster": cluster.Spec.Name,
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
		Name:  cluster.Spec.Name,
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

	stsSpec := v1beta2.StatefulSetSpec{
		Replicas:    &replicas,
		ServiceName: cluster.Spec.Name,
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
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Spec.Name,
			Namespace: namespace,
		},
		Spec: stsSpec,
	}
}
