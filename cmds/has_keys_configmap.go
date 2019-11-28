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
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func NewCmdHasKeysConfigMap(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var (
		keys []string
	)
	cmd := &cobra.Command{
		Use:               "configmap",
		Short:             "Check a configmap has a set of given keys",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				Fatal(errors.Errorf("missing key"))
			}
			if len(args) > 1 {
				Fatal(errors.Errorf("multiple names found: %v", strings.Join(args, ",")))
			}

			namespace, _, err := clientGetter.ToRawKubeConfigLoader().Namespace()
			if err != nil {
				Fatal(err)
			}

			config, err := clientGetter.ToRESTConfig()
			if err != nil {
				Fatal(err)
			}
			client, err := kubernetes.NewForConfig(config)
			if err != nil {
				Fatal(err)
			}

			name := args[0]
			obj, err := client.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})
			if err != nil {
				Fatal(err)
			}

			for _, key := range keys {
				_, ok := obj.Data[key]
				if !ok {
					Fatal(fmt.Errorf("missing key %s", key))
				}
			}
		},
	}
	cmd.Flags().StringSliceVar(&keys, "keys", nil, "Keys to search for")
	return cmd
}
