package cmds

import (
	"strings"
	"time"

	"github.com/appscode/go/types"
	"github.com/appscode/kutil/tools/clientcmd"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

func NewCmdWaitUntilReadyDeployment() *cobra.Command {
	var (
		interval = 2 * time.Second
		timeout  = 3 * time.Minute

		kubeContext string
		kubeConfig  string
		namespace   string
	)
	cmd := &cobra.Command{
		Use:               "deployment",
		Short:             "Wait until a deployment is ready",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				Fatal(errors.Errorf("missing crd"))
			}
			if len(args) > 1 {
				Fatal(errors.Errorf("multiple crds found: %v", strings.Join(args, ",")))
			}

			clientConfig, err := clientcmd.BuildConfigFromContext(kubeConfig, kubeContext)
			if err != nil {
				Fatal(err)
			}

			client, err := kubernetes.NewForConfig(clientConfig)
			if err != nil {
				Fatal(err)
			}

			name := args[0]

			err = wait.PollImmediate(interval, timeout, func() (bool, error) {
				if obj, err := client.AppsV1beta1().Deployments(namespace).Get(name, metav1.GetOptions{}); err == nil {
					return types.Int32(obj.Spec.Replicas) == obj.Status.ReadyReplicas, nil
				}
				return false, nil
			})
			if err != nil {
				Fatal(err)
			}
		},
	}

	cmd.Flags().StringVar(&kubeConfig, "kubeconfig", "", "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	cmd.Flags().StringVar(&kubeContext, "kube-context", "", "Name of kube context")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", metav1.NamespaceDefault, "Namespace used to search for deployment")
	cmd.Flags().DurationVar(&interval, "interval", interval, "Interval between checks")
	cmd.Flags().DurationVar(&timeout, "timeout", timeout, "Timeout")
	return cmd
}
