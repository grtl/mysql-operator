package cluster

import (
	"fmt"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/cli/config"
	"github.com/grtl/mysql-operator/cli/options"
	"github.com/grtl/mysql-operator/cli/util"
)

var removePVC bool

var clusterDeleteCmd = &cobra.Command{
	Use:   "delete [cluster names]",
	Short: "Deletes MySQL clusters",
	Long: `Deletes MySQL clusters and resources associated with them.
Unless explicitly specified, does not remove PersistentVolumeClaims.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		options := options.ExtractOptions(cmd)

		for _, arg := range args {
			err := deleteCluster(arg, options)
			util.FailOnErrorOrForceContinue(err, options)
		}
	},
}

func init() {
	clusterDeleteCmd.PersistentFlags().BoolVarP(&removePVC, "remove-pvc", "r", false, "remove PersistentVolumeClaims along with the cluster")
	Cmd.AddCommand(clusterDeleteCmd)
}

func deleteCluster(clusterName string, options *options.Options) error {
	fmt.Printf("Deleting: %s/%s\n", options.Namespace, clusterName)

	clustersInterface := config.GetConfig().Clientset().CrV1().MySQLClusters(options.Namespace)
	err := clustersInterface.Delete(clusterName, &v1.DeleteOptions{})

	if removePVC && (err == nil || options.Force) {
		deleteErr := deletePVC(clusterName, options)
		return errors.NewAggregate([]error{err, deleteErr})
	}

	return err
}

func deletePVC(clusterName string, options *options.Options) error {
	fmt.Printf("Deleting PVCs for: %s/%s\n", options.Namespace, clusterName)

	pvcInterface := config.GetConfig().KubeClientset().CoreV1().PersistentVolumeClaims(options.Namespace)
	return pvcInterface.DeleteCollection(&v1.DeleteOptions{}, v1.ListOptions{
		LabelSelector: labels.Set{"app": clusterName}.AsSelector().String(),
	})
}
