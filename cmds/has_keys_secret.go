package cmds

import (
	"fmt"
	"strings"

	"github.com/appscode/kutil/tools/clientcmd"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func NewCmdHasysSecret() *cobra.Command {
	var (
		kubeContext string
		kubeConfig  string
		namespace   string
		keys        []string
	)
	cmd := &cobra.Command{
		Use:               "secret",
		Short:             "Check a secret has a set of given keys",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				Fatal(errors.Errorf("missing key"))
			}
			if len(args) > 1 {
				Fatal(errors.Errorf("multiple names found: %v", strings.Join(args, ",")))
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
			obj, err := client.CoreV1().Secrets(namespace).Get(name, metav1.GetOptions{})
			if err != nil {
				Fatal(err)
			}

			for _, key := range keys {
				_, ok := obj.Data[key]
				if !ok {
					Fatal(fmt.Errorf("missing key %s", key))
				}
			}
		},
	}

	cmd.Flags().StringVar(&kubeConfig, "kubeconfig", "", "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	cmd.Flags().StringVar(&kubeContext, "kube-context", "", "Name of kube context")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", metav1.NamespaceDefault, "Namespace used to search for deployment")
	cmd.Flags().StringSliceVar(&keys, "keys", nil, "Keys to search for")
	return cmd
}
