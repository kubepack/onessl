package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdWaitUntilReady() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "wait-until-ready",
		Short:             "Wait until resource is ready",
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdWaitUntilReadyCRD())
	cmd.AddCommand(NewCmdWaitUntilReadyAPIService())
	cmd.AddCommand(NewCmdWaitUntilReadyDeployment())
	return cmd
}
