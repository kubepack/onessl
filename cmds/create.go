package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "create",
		Short:             `create PKI`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdCreateCA())
	cmd.AddCommand(NewCmdCreateServer())
	cmd.AddCommand(NewCmdCreateClient())
	return cmd
}
