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
	cmd.AddCommand(NewCmdPublicIPs())
	cmd.AddCommand(NewCmdPrivateIPs())
	cmd.AddCommand(NewCmdLinodeHostname())
	cmd.AddCommand(NewCmdSwapoff())
	return cmd
}
