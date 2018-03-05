package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdSSL() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ssl",
		Short:             `Utility commands for SSL certificates`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdCreate())
	cmd.AddCommand(NewCmdGet())
	return cmd
}
