package cmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/appscode/go/net"
	"github.com/spf13/cobra"
)

func NewCmdLinodeHostname() *cobra.Command {
	var (
		clusterName string
	)
	cmd := &cobra.Command{
		Use:               "hostname",
		Short:             "Prints hostname based on public IP for current Linode host",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			ips, _, err := net.HostIPs()
			if err != nil {
				Fatal(fmt.Errorf("failed to detect host ips. Reason: %v", err))
			}
			if len(ips) == 0 {
				os.Exit(1)
			}
			parts := strings.SplitN(ips[0], ".", 4)
			fmt.Printf("%s-%03s-%03s-%03s-%03s", clusterName, parts[0], parts[1], parts[2], parts[3])
			os.Exit(0)
		},
	}
	cmd.Flags().StringVarP(&clusterName, "cluster", "k", "", "Name of cluster")
	return cmd
}
