package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/grtl/mysql-operator/pkg/apis/cr"
)

var SchemeGroupVersion = schema.GroupVersion{Group: cr.GroupName, Version: "v1"}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&MySQLCluster{},
		&MySQLClusterList{},
	)
	return nil
}