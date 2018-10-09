package cmds

import (
	"time"

	dynamic_util "github.com/appscode/kutil/dynamic"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func NewCmdHasAnnotation(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var (
		apiVersion string
		kind       string
		name       string
		key        string
		value      string
		timeout    time.Duration
	)
	cmd := &cobra.Command{
		Use:               "has-annotation",
		Short:             "Check an object has a annotation Optionally with a given value",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			namespace, _, err := clientGetter.ToRawKubeConfigLoader().Namespace()
			if err != nil {
				Fatal(err)
			}

			config, err := clientGetter.ToRESTConfig()
			if err != nil {
				Fatal(err)
			}

			var v *string
			f := cmd.Flags().Lookup("value")
			if f.Changed {
				v = &value
			}
			_, err = dynamic_util.HasAnnotation(config, schema.FromAPIVersionAndKind(apiVersion, kind), namespace, name, key, v, timeout)
			if err != nil {
				Fatal(err)
			}
		},
	}
	cmd.Flags().StringVar(&apiVersion, "apiVersion", apiVersion, "api version of object")
	cmd.Flags().StringVar(&kind, "kind", kind, "Kind of object")
	cmd.Flags().StringVar(&name, "name", name, "Name of object")
	cmd.Flags().StringVar(&key, "key", key, "Key to check for value in object")
	cmd.Flags().StringVar(&value, "value", value, "Value of annotation of object")
	cmd.Flags().DurationVar(&timeout, "timeout", timeout, "Timeout for detecting annotation")
	return cmd
}
