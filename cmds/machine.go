package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdMachine() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "machine",
		Short:             `machine commands`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdNodeIP())
	cmd.AddCommand(NewCmdPrivateIPs())
	cmd.AddCommand(NewCmdPublicIPs())
	cmd.AddCommand(NewCmdSwapoff())
	return cmd
}
