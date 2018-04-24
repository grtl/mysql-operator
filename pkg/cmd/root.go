package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/grtl/mysql-operator/pkg/cmd/backup"
	"github.com/grtl/mysql-operator/pkg/cmd/cluster"
	"github.com/grtl/mysql-operator/pkg/cmd/util/config"
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

	rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "kubeconfig file location (default is $HOME/.kube/config)")
	rootCmd.PersistentFlags().StringP("namespace", "n", v1.NamespaceDefault, "Select namespace to modify objects in. (default is \"default\")")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "Ignore errors.")

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
