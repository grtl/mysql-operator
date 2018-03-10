package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackup is a representation of the MySQL cluster backup.
type MySQLBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec MySQLBackupSpec `json:"spec"`
}

// MySQLBackupSpec stores the properties of a MySQL cluster backup.
type MySQLBackupSpec struct {
	Cluster string `json:"cluster"`
	Time    string `json:"time"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupList represents a list of MySQL cluster backups
type MySQLBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLBackup `json:"items"`
}
