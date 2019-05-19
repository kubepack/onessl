package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kmodules.xyz/client-go/discovery"
)

func NewCmdHasAPI(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var (
		gv   string
		kind string
	)
	cmd := &cobra.Command{
		Use:               "has-api",
		Short:             "Check an api GroupVersionKind exists",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			dc, err := clientGetter.ToDiscoveryClient()
			if err != nil {
				Fatal(err)
			}

			if !discovery.IsPreferredAPIResource(dc, gv, kind) {
				Fatal(fmt.Errorf("mising apiVersion=%s kind=%s", gv, kind))
			}
		},
	}
	cmd.Flags().StringVar(&gv, "apiVersion", "", "apiVersion to search for (eg: apps/v1)")
	cmd.Flags().StringVar(&kind, "kind", "", "kind to search for (eg: Deployment)")
	return cmd
}
