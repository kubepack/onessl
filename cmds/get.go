package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get",
		Short:             `Get stuff`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdGetCACert())
	cmd.AddCommand(NewCmdNodeIP())
	cmd.AddCommand(NewCmdPublicIPs())
	cmd.AddCommand(NewCmdPrivateIPs())
	cmd.AddCommand(NewCmdLinodeHostname())
	return cmd
}
