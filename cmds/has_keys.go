package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdHasys() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "has-keys",
		Short:             "Checks configmap/secret has a set of given keys",
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdHasysConfigMap())
	cmd.AddCommand(NewCmdHasysSecret())
	return cmd
}
