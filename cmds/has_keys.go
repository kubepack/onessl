package cmds

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func NewCmdHasKeys(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "has-keys",
		Short:             "Checks configmap/secret has a set of given keys",
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdHasKeysConfigMap(clientGetter))
	cmd.AddCommand(NewCmdHasKeysSecret(clientGetter))
	return cmd
}
