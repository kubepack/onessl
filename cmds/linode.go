package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdLinode() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "linode",
		Short:             `linode commands`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdLinodeHostname())
	return cmd
}
