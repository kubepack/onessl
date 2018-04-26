package cmds

import (
	"fmt"
	"os"

	"github.com/appscode/kutil/tools/backup"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

func NewCmdBackup(clientConfig clientcmd.ClientConfig) *cobra.Command {
	var (
		ClusterName string
		BackupDir   string
		Sanitize    bool
	)
	cmd := &cobra.Command{
		Use:               "backup",
		Short:             "Backup cluster objects",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			restConfig, err := clientConfig.ClientConfig()
			if err != nil {
				Fatal(err)
			}
			mgr := backup.NewBackupManager(ClusterName, restConfig, Sanitize)
			filename, err := mgr.BackupToTar(BackupDir)
			if err != nil {
				Fatal(err)
			}
			fmt.Printf("Cluster objects are stored in %s", filename)
			os.Exit(0)
		},
	}
	cmd.Flags().BoolVar(&Sanitize, "sanitize", Sanitize, " Sanitize fields in YAML")
	cmd.Flags().StringVar(&BackupDir, "backup-dir", BackupDir, "Directory where yaml files will be saved")
	return cmd
}
