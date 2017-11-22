package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdVultr() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "vultr",
		Short:             `vultr commands`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdVultrPrivateIP())
	return cmd
}
