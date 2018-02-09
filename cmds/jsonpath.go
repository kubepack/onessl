package cmds

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/jsonpath"
)

// ref:
// - https://github.com/kubernetes/kubernetes/pull/9296
// - https://kubernetes.io/docs/reference/kubectl/jsonpath/
// - https://github.com/kubernetes/client-go/blob/master/util/jsonpath/jsonpath_test.go#L137:2
func NewCmdJsonpath() *cobra.Command {
	var (
		name             string
		template         string
		allowMissingKeys bool
	)
	cmd := &cobra.Command{
		Use:               "jsonpath",
		Short:             "Print value of jsonpath for input text",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				Fatal(fmt.Errorf("multiple templates found: %+v", args))
			}
			if len(args) == 0 {
				Fatal(fmt.Errorf("missing templates"))
			}
			template = args[0]

			reader := bufio.NewReader(os.Stdin)
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				Fatal(fmt.Errorf("failed to read input. Reason: %v", err))
			}

			var input interface{}
			err = json.Unmarshal(data, &input)
			if err != nil {
				Fatal(err)
			}

			j := jsonpath.New(name)
			j.AllowMissingKeys(allowMissingKeys)
			err = j.Parse(template)
			if err != nil {
				Fatal(fmt.Errorf("in %s, parse %s error %v", name, template, err))
			}
			buf := new(bytes.Buffer)
			err = j.Execute(buf, input)
			if err != nil {
				Fatal(fmt.Errorf("in %s, execute error %v", name, err))
			}
			fmt.Print(buf.String())
		},
	}

	cmd.Flags().BoolVar(&allowMissingKeys, "allowMissingKeys", allowMissingKeys, "Allow missing keys")
	return cmd
}
