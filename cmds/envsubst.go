package cmds

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/drone/envsubst"
	"github.com/spf13/cobra"
)

func NewCmdEnvsubst() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "envsubst",
		Short:             "Emulates bash environment variable substitution for input text",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				Fatal(fmt.Errorf("failed to read input. Reason: %v", err))
			}
			out, err := envsubst.EvalEnv(string(data))
			if err != nil {
				Fatal(fmt.Errorf("failed to decode input. Reason: %v", err))
			}
			fmt.Print(string(out))
		},
	}
	return cmd
}
