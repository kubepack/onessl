package cmds

import (
	"flag"

	v "github.com/appscode/go/version"
	"github.com/appscode/kutil/tools/cli"
	"github.com/spf13/cobra"
	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "onessl [command]",
		Short:             `onessl by AppsCode - Simple CLI to generate SSL certificates on any platform`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			cli.SendAnalytics(c, v.Version.Version)
		},
	}

	flags := rootCmd.PersistentFlags()
	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)

	kubeConfigFlags := genericclioptions.NewConfigFlags()
	kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags := cmdutil.NewMatchVersionFlags(kubeConfigFlags)
	matchVersionKubeConfigFlags.AddFlags(flags)

	flags.AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	flag.CommandLine.Parse([]string{})
	flags.BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send analytical events to Google Analytics")

	flags.BoolVar(&cli.EnableAnalytics, "analytics", cli.EnableAnalytics, "Send usage events to Google Analytics")
	flags.MarkDeprecated("analytics", "use --enable-analytics")

	flag.Set("stderrthreshold", "ERROR")

	rootCmd.AddCommand(NewCmdBase64())
	rootCmd.AddCommand(NewCmdCreate())
	rootCmd.AddCommand(NewCmdEnvsubst())
	rootCmd.AddCommand(NewCmdGet(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdJsonpath())
	rootCmd.AddCommand(NewCmdSemver())
	rootCmd.AddCommand(NewCmdHasKeys(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdWaitUntilHas(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdWaitUntilReady(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdBackup(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(NewCmdDoctor(matchVersionKubeConfigFlags))
	rootCmd.AddCommand(v.NewCmdVersion())
	return rootCmd
}
