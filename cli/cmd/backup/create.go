package backup

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/grtl/mysql-operator/cli/config"
	"github.com/grtl/mysql-operator/cli/options"
	"github.com/grtl/mysql-operator/cli/util"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

var (
	clusterName string
	backupName  string
	storage     string
	time        string
)

var backupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Schedules a new MySQL backup",
	Long: `Schedules a new MySQL backup for specified cluster. The backups will
be created according to cron-style time provided.
msp backup create --name "my-backup" --cluster "my-cluster"`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		options := options.ExtractOptions(cmd)

		if backupName == "" {
			backupName = clusterName
		}

		err := createMySQLBackup(options)
		if err != nil {
			util.FailWithError(err)
		}
	},
}

func init() {
	Cmd.AddCommand(backupCreateCmd)

	backupCreateCmd.Flags().StringVarP(&clusterName, "cluster", "c", "",
		"name of the cluster to be backed up")
	backupCreateCmd.MarkFlagRequired("cluster")
	backupCreateCmd.Flags().StringVarP(&time, "time", "t", "", "cron-style time")
	backupCreateCmd.MarkFlagRequired("time")
	backupCreateCmd.Flags().StringVarP(&storage, "storage", "s", "1Gi", "storage value")
	backupCreateCmd.Flags().StringVar(&backupName, "name", "", "backup name")
}

func createMySQLBackup(options *options.Options) error {
	fmt.Printf("Creating: %s/%s for %s\n", options.Namespace, backupName, clusterName)

	var (
		storageQuantity resource.Quantity
		err             error
	)

	if storage != "" {
		storageQuantity, err = resource.ParseQuantity(storage)
		if err != nil {
			return err
		}
	}

	mySQLBackupInterface := config.GetConfig().Clientset().CrV1().MySQLBackups(options.Namespace)
	_, err = mySQLBackupInterface.Create(&crv1.MySQLBackup{
		ObjectMeta: metav1.ObjectMeta{
			Name: backupName,
		},
		Spec: crv1.MySQLBackupSpec{
			Cluster: clusterName,
			Time:    time,
			Storage: storageQuantity,
		},
	})

	return err
}
