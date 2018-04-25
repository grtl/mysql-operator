package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupInstance represents an already created backup.
type MySQLBackupInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec MySQLBackupInstanceSpec `json:"spec"`
}

// MySQLBackupInstanceSpec stores the properties of a backup.
type MySQLBackupInstanceSpec struct {
	Schedule string `json:"schedule"`
	Cluster  string `json:"cluster"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupInstanceList represents a list of MySQLBackupInstances.
type MySQLBackupInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLBackupInstance `json:"items"`
}
