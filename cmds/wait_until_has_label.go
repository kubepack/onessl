package cmds

import (
	"fmt"
	"os"
	"time"

	dynamic_util "github.com/appscode/kutil/dynamic"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func NewCmdWaitUntilHasLabel(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var (
		apiVersion string
		kind       string
		name       string
		key        string
		value      string
		timeout    time.Duration
	)
	cmd := &cobra.Command{
		Use:               "label",
		Short:             "Wait until an object has a label optionally with a given value",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			config, err := clientGetter.ToRESTConfig()
			if err != nil {
				Fatal(err)
			}

			var namespace string
			f := cmd.Flags().Lookup("namespace")
			if f != nil {
				namespace = f.Value.String()
			}

			var v *string
			f = cmd.Flags().Lookup("value")
			if f != nil && f.Changed {
				v = &value
			}

			out, err := dynamic_util.UntilHasLabel(config, schema.FromAPIVersionAndKind(apiVersion, kind), namespace, name, key, v, timeout)
			if err != nil {
				Fatal(err)
			}
			fmt.Print(out)
			os.Exit(0)
		},
	}
	cmd.Flags().StringVar(&apiVersion, "apiVersion", apiVersion, "api version of object")
	cmd.Flags().StringVar(&kind, "kind", kind, "Kind of object")
	cmd.Flags().StringVar(&name, "name", name, "Name of object")
	cmd.Flags().StringVar(&key, "key", key, "Key to check for value in object")
	cmd.Flags().StringVar(&value, "value", value, "Value of label of object")
	cmd.Flags().DurationVar(&timeout, "timeout", timeout, "Timeout for detecting label")
	return cmd
}
