package backup

import (
	"fmt"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/pkg/cmd/util/config"
	"github.com/grtl/mysql-operator/pkg/cmd/util/fail"
	"github.com/grtl/mysql-operator/pkg/cmd/util/options"
	"github.com/grtl/mysql-operator/pkg/operator/backupschedule"
)

var removePVC bool

var backupDeleteCmd = &cobra.Command{
	Use:   "delete [backupschedule name]",
	Short: "Deletes MySQL Backup schedule.",
	Long: `Deletes MySQL Backup schedule and resources associated with them.
Unless explicitly specified, does not remove Persistent Volume Claim.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		opts := options.ExtractOptions(cmd)

		for _, arg := range args {
			err := deleteBackup(arg, opts)
			fail.ErrorOrForceContinue(err, opts)
		}
	},
}

func deleteBackup(backupName string, opts *options.Options) error {
	fmt.Printf("Deleting: %s/%s\n", opts.Namespace, backupName)

	backupsInterface := config.GetConfig().Clientset().CrV1().MySQLBackupSchedules(opts.Namespace)
	err := backupsInterface.Delete(backupName, &v1.DeleteOptions{})

	if removePVC && (err == nil || opts.Force) {
		deleteErr := deletePVC(clusterName, opts)
		return errors.NewAggregate([]error{err, deleteErr})
	}

	return err
}

func deletePVC(backupName string, opts *options.Options) error {
	fmt.Printf("Deleting PVC for: %s/%s\n", opts.Namespace, clusterName)

	pvcInterface := config.GetConfig().KubeClientset().CoreV1().PersistentVolumeClaims(opts.Namespace)
	return pvcInterface.Delete(backupschedule.PVCName(backupName), &v1.DeleteOptions{})
}

func init() {
	backupDeleteCmd.PersistentFlags().BoolVarP(&removePVC, "remove-pvc", "r", false,
		"remove Persistent Volume Claim along with the backup schedule. This will remove all the backups!")
	Cmd.AddCommand(backupDeleteCmd)
}
