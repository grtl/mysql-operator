package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLCluster is a representation of MySQL Cluster.
type MySQLCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   MySQLClusterSpec   `json:"spec"`
	Status MySQLClusterStatus `json:"status,omitempty"`
}

// MySQLClusterSpec stores the properties of a MySQL Cluster.
type MySQLClusterSpec struct {
	Password   string            `json:"password"`
	Storage    resource.Quantity `json:"storage"`
	Replicas   int32             `json:"replicas"`
	FromBackup BackupInstance    `json:"fromBackup,omitempty"`
}

type MySQLClusterStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLClusterList represents a list of MySQL Clusters
type MySQLClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLCluster `json:"items"`
}

type BackupInstance struct {
	BackupName string `json:"backupName"`
	Instance   string `json:"instance"`
}
