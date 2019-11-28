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
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

var (
	certDir = func() string {
		if v, ok := os.LookupEnv("ONESSL_PKI_DIR"); ok {
			return v
		}
		dir, err := os.Getwd()
		if err != nil {
			dir = homedir.HomeDir()
		}
		return dir
	}()
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "create",
		Short:             `create PKI`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdCreateCA(certDir))
	cmd.AddCommand(NewCmdCreateServer(certDir))
	cmd.AddCommand(NewCmdCreatePeer(certDir))
	cmd.AddCommand(NewCmdCreateClient(certDir))
	return cmd
}
