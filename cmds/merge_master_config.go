package cmds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/appscode/mergo"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1"
	kubeadmcmd "k8s.io/kubernetes/cmd/kubeadm/app/cmd"
	"k8s.io/kubernetes/cmd/kubeadm/app/features"
	kubeadmutil "k8s.io/kubernetes/cmd/kubeadm/app/util"
)

func NewCmdMergeMasterConfig() *cobra.Command {
	cfg := &kubeadmapi.MasterConfiguration{}
	var cfgPath string
	var featureGatesString string
	cmd := &cobra.Command{
		Use:               "master-config",
		Short:             `Merge Kubeadm master configuration`,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if cfg.FeatureGates, err = features.NewFeatureGate(&features.InitFeatureGates, featureGatesString); err != nil {
				kubeadmutil.CheckErr(err)
			}

			if cfgPath != "" {
				data, err := ioutil.ReadFile(cfgPath)
				if err != nil {
					Fatal(err)
				}
				data, err = yaml.YAMLToJSON(data)
				if err != nil {
					Fatal(err)
				}

				var in kubeadmapi.MasterConfiguration
				err = json.Unmarshal(data, &in)
				if err != nil {
					Fatal(err)
				}

				err = mergo.MergeWithOverwrite(cfg, &in)
				if err != nil {
					Fatal(err)
				}
			}

			cfg.APIVersion = "kubeadm.k8s.io/v1alpha1"
			cfg.Kind = "MasterConfiguration"
			data, err := yaml.Marshal(cfg)
			if err != nil {
				Fatal(err)
			}
			fmt.Println(string(data))
			os.Exit(0)
		},
	}
	// ref: https://github.com/kubernetes/kubernetes/blob/0b9efaeb34a2fc51ff8e4d34ad9bc6375459c4a4/cmd/kubeadm/app/cmd/init.go#L141
	kubeadmcmd.AddInitConfigFlags(cmd.Flags(), cfg, &featureGatesString)
	cmd.Flags().StringVar(&cfgPath, "config", cfgPath, "Path to kubeadm config file (WARNING: Usage of a configuration file is experimental)")

	return cmd
}
