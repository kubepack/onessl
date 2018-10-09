package cmds

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func NewCmdWaitUntilHas(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "wait-until-has",
		Short:             "Wait until has some attribute",
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdWaitUntilHasLabel(clientGetter))
	cmd.AddCommand(NewCmdWaitUntilHasAnnotation(clientGetter))
	return cmd
}
