package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/grtl/mysql-operator/pkg/apis/cr"
)

// SchemeGroupVersion with MySQL Cluster custom resource group
var SchemeGroupVersion = schema.GroupVersion{Group: cr.GroupName, Version: "v1"}

var (
	// SchemeBuilder with MySQL Cluster types added
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme exported for convenience
	AddToScheme = SchemeBuilder.AddToScheme
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		new(MySQLCluster),
		new(MySQLClusterList),
		new(MySQLBackupSchedule),
		new(MySQLBackupScheduleList),
		new(MySQLBackupInstance),
		new(MySQLBackupInstanceList),
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
