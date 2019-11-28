/*
Copyright The Kubepack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
