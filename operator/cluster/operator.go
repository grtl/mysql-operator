package cluster

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

const mySQLPortNumber = 3306

// AddCluster creates the Kubernetes API objects necessary for a MySQL cluster.
func AddCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) {
	createServiceForCluster(cluster, kubeClientset)
	createStatefulSetForCluster(cluster, kubeClientset)
}

func createServiceForCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) {
	servicesInterface := kubeClientset.CoreV1().Services(cluster.ObjectMeta.Namespace)

	newService := serviceForCluster(cluster)
	_, err := servicesInterface.Create(&newService)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		logrus.Panic(err)
	} else if apierrors.IsAlreadyExists(err) {
		logrus.WithFields(logrus.Fields{
			"cluster": cluster.Name,
		}).Info("Service for cluster already exists")
	}
}

func createStatefulSetForCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) {
	statefulSetsInterface := kubeClientset.AppsV1().StatefulSets(cluster.ObjectMeta.Namespace)

	newStatefulSet := statefulSetForCluster(cluster)
	_, err := statefulSetsInterface.Create(&newStatefulSet)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		logrus.Panic(err)
	} else if apierrors.IsAlreadyExists(err) {
		logrus.WithFields(logrus.Fields{
			"cluster": cluster.Name,
		}).Info("StatefulSet for cluster already exists")
	}
}

func serviceForCluster(cluster *crv1.MySQLCluster) corev1.Service {
	return corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: cluster.Spec.Name,
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{
				{
					Port: mySQLPortNumber,
				},
			},
			Selector: map[string]string{
				"mysql-cluster": cluster.Spec.Name,
			},
		},
	}
}

func statefulSetForCluster(cluster *crv1.MySQLCluster) appsv1.StatefulSet {
	namespace := cluster.ObjectMeta.Namespace

	stsSpec := statefulSetSpecForCluster(cluster)

	return appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Spec.Name,
			Namespace: namespace,
		},
		Spec: stsSpec,
	}
}

func statefulSetSpecForCluster(cluster *crv1.MySQLCluster) appsv1.StatefulSetSpec {
	labels := map[string]string{
		"mysql-cluster": cluster.Spec.Name,
	}

	var replicas int32 = 1

	selector := metav1.LabelSelector{
		MatchLabels: labels,
	}

	container := containerForCluster(cluster)

	pvc := pvcForCluster(cluster)

	return appsv1.StatefulSetSpec{
		Replicas:    &replicas,
		ServiceName: cluster.Spec.Name,
		Selector:    &selector,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: labels,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					container,
				},
			},
		},
		VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
			pvc,
		},
	}
}

func containerForCluster(cluster *crv1.MySQLCluster) corev1.Container {
	envVars := []corev1.EnvVar{
		{
			Name:  "MYSQL_ROOT_PASSWORD",
			Value: cluster.Spec.Password,
		},
	}

	return corev1.Container{
		Name:  cluster.Spec.Name,
		Image: "mysql:8",
		Env:   envVars,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "mysql",
				MountPath: "/var/lib/mysql",
			},
		},
	}
}

func pvcForCluster(cluster *crv1.MySQLCluster) corev1.PersistentVolumeClaim {
	return corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mysql",
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				"ReadWriteOnce",
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					"storage": cluster.Spec.Storage,
				},
			},
		},
	}
}
