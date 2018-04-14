package cluster

import (
	"fmt"

	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/grtl/mysql-operator/pkg/cmd/util/config"
	"github.com/grtl/mysql-operator/pkg/cmd/util/fail"
	"github.com/grtl/mysql-operator/pkg/cmd/util/options"
)

var (
	portUpdate     int32
	replicasUpdate int32
)

var clusterUpdateCmd = &cobra.Command{
	Use:   "update [cluster name]",
	Short: "Updates a MySQL Cluster",
	Long: `Updates a MySQL Cluster.
Specify replicas number and port to update your MySQL Cluster:
msp cluster update "my-cluster" --replicas 4 --port 1337`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if replicasUpdate != 0 || portUpdate != 0 {
			options := options.ExtractOptions(cmd)
			err := updateCluster(args[0], options)
			if err != nil {
				fail.Error(err)
			}
		}
	},
}

func init() {
	Cmd.AddCommand(clusterUpdateCmd)

	clusterUpdateCmd.Flags().Int32Var(&replicasUpdate, "replicas", 0, "replicas number")
	clusterUpdateCmd.Flags().Int32Var(&portUpdate, "port", 0, "port number")
}

func updateCluster(clusterName string, options *options.Options) error {
	mySQLClusterInterface := config.GetConfig().Clientset().CrV1().MySQLClusters(options.Namespace)

	mySQLCluster, err := mySQLClusterInterface.Get(clusterName, *new(metav1.GetOptions))
	if err != nil {
		return err
	}

	if replicasUpdate != 0 {
		mySQLCluster.Spec.Replicas = replicasUpdate
	}
	if portUpdate != 0 {
		mySQLCluster.Spec.Port = portUpdate
	}

	fmt.Printf("Updating: %s/%s\n", options.Namespace, clusterName)
	_, err = mySQLClusterInterface.Update(mySQLCluster)

	return err
}
