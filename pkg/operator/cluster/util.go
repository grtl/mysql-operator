package cluster

import (
	"fmt"
	"text/template"
)

// FuncMap can be used to execute templates with the helper functions from
// the cluster operator.
var FuncMap = template.FuncMap{
	"StatefulSetName": StatefulSetName,
	"ServiceName":     ServiceName,
	"ReadServiceName": ReadServiceName,
}

// StatefulSetName returns a name for the stateful set associated with the
// given clusterName.
func StatefulSetName(clusterName string) string {
	return clusterName
}

// ServiceName returns a name for the service associated with the given
// clusterName.
func ServiceName(clusterName string) string {
	return clusterName
}

// ReadServiceName returns a name for the read service associated with the
// given clusterName.
func ReadServiceName(clusterName string) string {
	return fmt.Sprintf("%s-read", clusterName)
}
