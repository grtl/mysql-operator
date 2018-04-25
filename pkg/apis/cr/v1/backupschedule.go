package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupSchedule represents a backup schedule for a MySQL cluster.
type MySQLBackupSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec MySQLBackupScheduleSpec `json:"spec"`
}

// MySQLBackupScheduleSpec stores the properties of a backup schedule.
type MySQLBackupScheduleSpec struct {
	Cluster string            `json:"cluster"`
	Time    string            `json:"time"`
	Storage resource.Quantity `json:"storage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MySQLBackupScheduleList represents a list of MySQLBackupSchedules.
type MySQLBackupScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MySQLBackupSchedule `json:"items"`
}
