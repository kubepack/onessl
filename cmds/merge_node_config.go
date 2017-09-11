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
)

func NewCmdMergeNodeConfig() *cobra.Command {
	cfg := &kubeadmapi.NodeConfiguration{}
	var cfgPath string
	var skipPreFlight bool
	cmd := &cobra.Command{
		Use:               "node-config",
		Short:             `Merge Kubeadm node configuration`,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if cfgPath != "" {
				data, err := ioutil.ReadFile(cfgPath)
				if err != nil {
					Fatal(err)
				}
				data, err = yaml.YAMLToJSON(data)
				if err != nil {
					Fatal(err)
				}

				var in kubeadmapi.NodeConfiguration
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
			cfg.Kind = "NodeConfiguration"
			data, err := yaml.Marshal(cfg)
			if err != nil {
				Fatal(err)
			}
			fmt.Println(string(data))
			os.Exit(0)
		},
	}

	cmd.PersistentFlags().StringVar(
		&cfgPath, "config", cfgPath,
		"Path to kubeadm config file")

	cmd.PersistentFlags().StringVar(
		&cfg.DiscoveryFile, "discovery-file", "",
		"A file or url from which to load cluster information")
	cmd.PersistentFlags().StringVar(
		&cfg.DiscoveryToken, "discovery-token", "",
		"A token used to validate cluster information fetched from the master")
	cmd.PersistentFlags().StringVar(
		&cfg.NodeName, "node-name", "",
		"Specify the node name")
	cmd.PersistentFlags().StringVar(
		&cfg.TLSBootstrapToken, "tls-bootstrap-token", "",
		"A token used for TLS bootstrapping")
	cmd.PersistentFlags().StringVar(
		&cfg.Token, "token", "",
		"Use this token for both discovery-token and tls-bootstrap-token")

	cmd.PersistentFlags().BoolVar(
		&skipPreFlight, "skip-preflight-checks", false,
		"Skip preflight checks normally run before modifying the system",
	)

	return cmd
}
