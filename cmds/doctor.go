package cmds

import (
	"fmt"
	"os"

	"github.com/appscode/kutil/tools/doctor"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewCmdDoctor(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "doctor",
		Short:             "Diagnoses Kubernetes Cluster Setup issues",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := clientGetter.ToRESTConfig()
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
				Fatal(errors.Wrapf(err, "cluster info is not valid.\n%s", info))
			} else if glog.V(10) {
				fmt.Println("cluster info:")
				fmt.Println(info)
			}

			os.Exit(0)
		},
	}
	return cmd
}
