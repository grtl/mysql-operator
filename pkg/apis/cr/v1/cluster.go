package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Default values for optional fields.
const (
	DefaultReplicas int32 = 2
	DefaultPort     int32 = 3306
	DefaultImage          = "mysql:latest"
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
	// Secret is the name of Kubernetes secret containing the password.
	Secret string `json:"secret"`
	// Storage indicates the size of the Persistent Volume Claim for each replica.
	Storage resource.Quantity `json:"storage"`
	// Number of mysql instances in the cluster.
	Replicas int32 `json:"replicas,omitempty"`
	// Port specifies port for MySQL server.
	Port int32 `json:"port,omitempty"`
	// Image allows to specify mysql image
	Image string `json:"image,omitempty"`
	// FromBackup lets you specify the backup name to restore the cluster from.
	FromBackup BackupInstance `json:"fromBackup,omitempty"`
}

// MySQLClusterStatus represents a cluster's status.
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

// BackupInstance represents a single backupschedule instance.
type BackupInstance struct {
	BackupName string `json:"backupName"`
	Instance   string `json:"instance"`
}

// WithDefaults fills cluster missing fields with their default values.
func (c *MySQLCluster) WithDefaults() {
	if c.Spec.Replicas == 0 {
		c.Spec.Replicas = DefaultReplicas
	}

	if c.Spec.Port == 0 {
		c.Spec.Port = DefaultPort
	}

	if c.Spec.Image == "" {
		c.Spec.Image = DefaultImage
	}
}
