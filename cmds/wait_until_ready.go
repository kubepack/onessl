package cmds

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewCmdWaitUntilReady(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "wait-until-ready",
		Short:             "Wait until resource is ready",
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdWaitUntilReadyCRD(clientGetter))
	cmd.AddCommand(NewCmdWaitUntilReadyAPIService(clientGetter))
	cmd.AddCommand(NewCmdWaitUntilReadyDeployment(clientGetter))
	return cmd
}
