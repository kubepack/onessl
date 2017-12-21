package cmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/appscode/go/net"
	"github.com/spf13/cobra"
)

func NewCmdPrivateIPs() *cobra.Command {
	all := true
	cmd := &cobra.Command{
		Use:               "private-ips",
		Short:             "Prints private ip(s) for current host",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			_, ips, err := net.HostIPs()
			if err != nil {
				Fatal(fmt.Errorf("failed to detect host ips. Reason: %v", err))
			}
			if !all && len(ips) > 0 {
				fmt.Print(ips[0])
			} else {
				fmt.Print(strings.Join(ips, ","))
			}
			os.Exit(0)
		},
	}
	cmd.Flags().BoolVar(&all, "all", all, "If true, returns all private ips.")
	return cmd
}
