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

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gomodules.xyz/cert/certstore"
)

func NewCmdCreateCA(certDir string) *cobra.Command {
	var (
		org       []string
		prefix    string
		overwrite bool
	)
	cmd := &cobra.Command{
		Use:               "ca-cert",
		Short:             "Create CA cert/key pair",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			store, err := certstore.NewCertStore(afero.NewOsFs(), certDir, org...)
			if err != nil {
				fmt.Printf("Failed to create certificate store. Reason: %v.", err)
				os.Exit(1)
			}

			var p []string
			if prefix != "" {
				p = append(p, prefix)
			}
			if store.IsExists("ca", p...) && !overwrite {
				fmt.Printf("CA certificate found at %s.", store.Location())
				os.Exit(1)
			}

			err = store.NewCA(p...)
			if err != nil {
				fmt.Printf("Failed to init ca. Reason: %v.", err)
				os.Exit(1)
			}
			fmt.Println("Wrote ca certificates in ", store.Location())
			os.Exit(0)
		},
	}

	cmd.Flags().StringVar(&certDir, "cert-dir", certDir, "Path to directory where pki files are stored.")
	cmd.Flags().StringSliceVarP(&org, "organization", "o", org, "Name of client organizations.")
	cmd.Flags().StringVarP(&prefix, "prefix", "p", prefix, "Prefix added to certificate files")
	cmd.Flags().BoolVar(&overwrite, "overwrite", overwrite, "Overwrite existing cert/key pair")
	return cmd
}
