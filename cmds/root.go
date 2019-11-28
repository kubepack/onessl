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
	"flag"

	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cliflag "k8s.io/component-base/cli/flag"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"kmodules.xyz/client-go/logs"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "onessl [command]",
		Short:             `onessl by AppsCode - Simple CLI to generate SSL certificates on any platform`,
		DisableAutoGenTag: true,
	}

	flags := rootCmd.PersistentFlags()
	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	kubeConfigFlags := genericclioptions.NewConfigFlags(true)
	kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	matchVersionKubeConfigFlags.AddFlags(flags)

	flags.AddGoFlagSet(flag.CommandLine)
	logs.ParseFlags()

	rootCmd.AddCommand(NewCmdBase64())
	rootCmd.AddCommand(NewCmdCreate())
	rootCmd.AddCommand(NewCmdEnvsubst())
	rootCmd.AddCommand(NewCmdGet(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdJsonpath())
	rootCmd.AddCommand(NewCmdSemver())
	rootCmd.AddCommand(NewCmdHasAPI(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdHasKeys(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdWaitUntilHas(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdWaitUntilReady(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdBackup(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdDoctor(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(v.NewCmdVersion())
	return rootCmd
}
