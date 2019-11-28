/*
Copyright The Kubepack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmds

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
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
				Fatal(errors.Errorf("multiple templates found: %v", strings.Join(args, ",")))
			}
			if len(args) == 0 {
				Fatal(errors.Errorf("missing templates"))
			}
			template = args[0]

			reader := bufio.NewReader(os.Stdin)
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				Fatal(errors.Wrap(err, "failed to read input"))
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
				Fatal(errors.Errorf("in %s, parse %s error %v", name, template, err))
			}
			buf := new(bytes.Buffer)
			err = j.Execute(buf, input)
			if err != nil {
				Fatal(errors.Errorf("in %s, execute error %v", name, err))
			}
			fmt.Print(buf.String())
		},
	}

	cmd.Flags().BoolVar(&allowMissingKeys, "allowMissingKeys", allowMissingKeys, "Allow missing keys")
	return cmd
}
