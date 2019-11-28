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
