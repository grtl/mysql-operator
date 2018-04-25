package cluster

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	errorsAgg "k8s.io/apimachinery/pkg/util/errors"

	mysqlv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/cmd/util/config"
	"github.com/grtl/mysql-operator/pkg/cmd/util/fail"
	"github.com/grtl/mysql-operator/pkg/cmd/util/options"
)

var (
	replicas       int32
	password       string
	storage        string
	secretName     string
	backupName     string
	backupInstance string
	fromSecret     string
	port           int32
	image          string
)

var clusterCreateCmd = &cobra.Command{
	Use:   "create [cluster name]",
	Short: "Creates a MySQL cluster",
	Long: `Creates a new MySQL cluster.
You can specify your own secret using the from-secret flag:
msp cluster create "my-cluster" --from-secret "my-secret"
or create a new secret using the secret and password flags:
msp cluster create "my-cluster" --secret "your-new-secret" --password "mysql-password"
In order to create a cluster from a backupschedule use the backup and instance flags:
msp cluster create "my-cluster" --backup "backup-name" --instance "instance"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		options := options.ExtractOptions(cmd)

		if fromSecret == "" {
			if secretName == "" || password == "" {
				fmt.Fprintln(os.Stderr, "You must provide a secret and a password (or create"+
					" a cluster from an existing secret using the from-secret flag).")
				fmt.Fprintln(os.Stderr, "Use msp help cluster create for more information.")
				os.Exit(1)
			}
			err := createSecret(options)
			if err != nil {
				fail.Error(err)
			}
		} else {
			secretName = fromSecret
		}

		err := createMySQLCluster(args[0], options)
		if err != nil {
			fail.Error(err)
		}
	},
}

func init() {
	Cmd.AddCommand(clusterCreateCmd)

	clusterCreateCmd.Flags().StringVarP(&storage, "storage", "s", "1Gi", "storage value")
	clusterCreateCmd.Flags().StringVarP(&password, "password", "p",
		"", "password your-password")
	clusterCreateCmd.Flags().Int32Var(&replicas, "replicas", mysqlv1.DefaultReplicas, "replicas number")
	clusterCreateCmd.Flags().StringVar(&secretName, "secret", "", "secret secrete-name")
	clusterCreateCmd.Flags().StringVar(&backupName, "backupschedule", "", "backupschedule backupschedule-name")
	clusterCreateCmd.Flags().StringVar(&backupInstance, "instance",
		"", "instance backupschedule-instance")
	clusterCreateCmd.Flags().StringVar(&fromSecret, "from-secret", "", "from-secret secret-name")
	clusterCreateCmd.Flags().StringVarP(&image, "image", "i",
		mysqlv1.DefaultImage, "image your-image")
	clusterCreateCmd.Flags().Int32Var(&port, "port", mysqlv1.DefaultPort, "port number")
}

func createSecret(options *options.Options) error {
	fmt.Printf("Creating: %s/%s\n", options.Namespace, secretName)

	data := make(map[string]string)
	data["password"] = password

	secreteInterface := config.GetConfig().KubeClientset().CoreV1().Secrets(options.Namespace)
	_, err := secreteInterface.Create(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: secretName,
		},
		StringData: data,
	})

	return err
}

func createMySQLCluster(clusterName string, options *options.Options) error {
	fmt.Printf("Creating: %s/%s using %s secret\n", options.Namespace, clusterName, secretName)

	storageQuantity, err := resource.ParseQuantity(storage)
	if err != nil {
		removeErr := removeSecret(options.Namespace)
		return errorsAgg.NewAggregate([]error{err, removeErr})
	}

	mySQLClusterInterface := config.GetConfig().Clientset().CrV1().MySQLClusters(options.Namespace)

	_, err = mySQLClusterInterface.Create(&mysqlv1.MySQLCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterName,
		},
		Spec: mysqlv1.MySQLClusterSpec{
			Secret:   secretName,
			Storage:  storageQuantity,
			Replicas: replicas,
			Port:     port,
			Image:    image,
			FromBackup: mysqlv1.BackupInstance{
				BackupName: backupName,
				Instance:   backupInstance,
			},
		},
	})

	if err != nil {
		removeErr := removeSecret(options.Namespace)
		return errorsAgg.NewAggregate([]error{err, removeErr})
	}

	return nil
}

func removeSecret(namespace string) error {
	if fromSecret != "" {
		return nil
	}
	secretInterface := config.GetConfig().KubeClientset().CoreV1().Secrets(namespace)
	return secretInterface.Delete(secretName, new(metav1.DeleteOptions))
}
