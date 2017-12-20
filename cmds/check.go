package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "check",
		Short:             `Check stuff`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdCheckMasterStatus())
	return cmd
}
