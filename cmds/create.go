package cmds

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

var (
	certDir = func() string {
		if v, ok := os.LookupEnv("ONESSL_PKI_DIR"); ok {
			return v
		}
		dir, err := os.Getwd()
		if err != nil {
			dir = homedir.HomeDir()
		}
		return dir
	}()
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "create",
		Short:             `create PKI`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdCreateCA(certDir))
	cmd.AddCommand(NewCmdCreateServer(certDir))
	cmd.AddCommand(NewCmdCreatePeer(certDir))
	cmd.AddCommand(NewCmdCreateClient(certDir))
	return cmd
}
