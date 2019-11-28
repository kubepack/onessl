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
	"os"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kmodules.xyz/client-go/tools/doctor"
)

func NewCmdDoctor(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "doctor",
		Short:             "Diagnoses Kubernetes Cluster Setup issues",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := clientGetter.ToRESTConfig()
			if err != nil {
				Fatal(errors.Wrap(err, "failed to read kubeconfig"))
			}

			d, err := doctor.New(cfg)
			if err != nil {
				Fatal(err)
			}
			info, err := d.GetClusterInfo()
			if err != nil {
				Fatal(errors.Wrap(err, "failed to get cluster info"))
			}
			err = info.Validate()
			if err != nil {
				Fatal(errors.Wrapf(err, "cluster info is not valid.\n%s", info))
			} else if glog.V(10) {
				fmt.Println("cluster info:")
				fmt.Println(info)
			}

			os.Exit(0)
		},
	}
	return cmd
}
