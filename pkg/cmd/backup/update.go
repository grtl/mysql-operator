package backup

import (
	"fmt"

	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/grtl/mysql-operator/pkg/cmd/util/config"
	"github.com/grtl/mysql-operator/pkg/cmd/util/fail"
	options "github.com/grtl/mysql-operator/pkg/cmd/util/options"
)

var timeUpd string

var backupUpdateCmd = &cobra.Command{
	Use:   "update [backupschedule name]",
	Short: "Updates a MySQL Backup Schedule",
	Long: `Updates a MySQL Backup Schedule. Specify time in CRON style to update your backup schedule
msp backup update "my-backup" --time "59 23 31 DEC Fri *"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if timeUpd != "" {
			options := options.ExtractOptions(cmd)
			err := updateBackup(args[0], options)
			if err != nil {
				fail.Error(err)
			}
		}
	},
}

func init() {
	Cmd.AddCommand(backupUpdateCmd)

	backupUpdateCmd.Flags().StringVarP(&timeUpd, "time", "t", "", "CRON style time ")
}

func updateBackup(backupName string, options *options.Options) error {
	mySQLBackupInterface := config.GetConfig().Clientset().CrV1().MySQLBackupSchedules(options.Namespace)

	mySQLBackup, err := mySQLBackupInterface.Get(backupName, *new(metav1.GetOptions))
	if err != nil {
		return err
	}

	mySQLBackup.Spec.Time = timeUpd

	fmt.Printf("Updating: %s/%s\n", options.Namespace, backupName)
	_, err = mySQLBackupInterface.Update(mySQLBackup)

	return err
}
