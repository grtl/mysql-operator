package cmd

import (
	"fmt"
	"os"

	"github.com/grtl/mysql-operator/cli/cmd/backup"
	"github.com/grtl/mysql-operator/cli/cmd/cluster"

	"github.com/grtl/mysql-operator/cli/cmd/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "msp",
	Short: "MySQL Operator",
	Long: `MySQL Operator is a command line interface that allows
you to manage and create MySQL Clusters and Backups`,
}

var kubeconfig string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.EnablePrefixMatching = true

	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "kubeconfig file (default is $HOME/.kube/config)")

	rootCmd.AddCommand(cluster.Cmd)
	rootCmd.AddCommand(backup.Cmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if kubeconfig == "" {
		if value, ok := os.LookupEnv("HOME"); ok {
			kubeconfig = fmt.Sprintf("%s/.kube/config", value)
		}
	}

	if err := config.InitBaseConfig(kubeconfig); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
