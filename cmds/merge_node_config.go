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
	// ref: https://github.com/kubernetes/kubernetes/blob/0b9efaeb34a2fc51ff8e4d34ad9bc6375459c4a4/cmd/kubeadm/app/cmd/join.go#L122
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
	cmd.PersistentFlags().StringSliceVar(
		&cfg.DiscoveryTokenCACertHashes, "discovery-token-ca-cert-hash", []string{},
		"For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	cmd.PersistentFlags().BoolVar(
		&cfg.DiscoveryTokenUnsafeSkipCAVerification, "discovery-token-unsafe-skip-ca-verification", false,
		"For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")

	cmd.PersistentFlags().StringVar(
		&cfg.Token, "token", "",
		"Use this token for both discovery-token and tls-bootstrap-token")

	return cmd
}
