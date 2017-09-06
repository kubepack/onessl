package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdMerge() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "merge",
		Short:             `Merge Kubeadm config`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdMergeMasterConfig())
	cmd.AddCommand(NewCmdMergeNodeConfig())
	return cmd
}
