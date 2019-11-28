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

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

// kubectl config view --minify=true --flatten -o json | onessl jsonpath '{.clusters[0].cluster.certificate-authority-data}'
func NewCmdGetKubeCA(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "kube-ca",
		Short:             "Prints CA certificate for Kubernetes cluster from Kubeconfig",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := clientGetter.ToRESTConfig()
			if err != nil {
				Fatal(errors.Wrap(err, "failed to read kubeconfig"))
			}
			err = rest.LoadTLSFiles(cfg)
			if err != nil {
				Fatal(errors.Wrap(err, "failed to load tls files"))
			}
			if cfg.Insecure {
				Fatal(errors.New(`Kube apiserver uses "insecure-skip-tls-verify: true". Kube apiserver must not be accessed without verifying the TLS certificate.`))
			}
			fmt.Println(string(cfg.CAData))
			os.Exit(0)
		},
	}
	return cmd
}
