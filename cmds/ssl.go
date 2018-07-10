package cmds

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func NewCmdSSL(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ssl",
		Short:             `Utility commands for SSL certificates`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdCreate())
	cmd.AddCommand(NewCmdGet(clientGetter))
	return cmd
}
