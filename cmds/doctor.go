package cmds

import (
	"fmt"
	"os"

	"github.com/appscode/kutil/tools/doctor"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

// kubectl config view --minify=true --flatten -o json | onessl jsonpath '{.clusters[0].cluster.certificate-authority-data}'
func NewCmdDoctor(clientConfig clientcmd.ClientConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "doctor",
		Short:             "Diagnoses Kubernetes Cluster Setup issues",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := clientConfig.ClientConfig()
			if err != nil {
				Fatal(errors.Wrap(err, "failed to read kubeconfig"))
			}

			d, err := doctor.New(cfg)
			if err != nil {
				Fatal(err)
			}
			info, err := d.GetClusterInfo()
			if err != nil {
				Fatal(errors.Wrap(err, "failed to get cluster info"))
			}
			err = info.Validate()
			if err != nil {
				fmt.Println(info)
				Fatal(errors.Wrap(err, "cluster info is not valid"))
			}
			os.Exit(0)
		},
	}
	return cmd
}
