package cmds

import (
	"flag"

	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cliflag "k8s.io/component-base/cli/flag"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
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
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})

	flag.Set("stderrthreshold", "ERROR")

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
