package cmds

import (
	"fmt"
	"os"
	"strings"

	"github.com/appscode/go/net"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/sets"
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
				Fatal(fmt.Errorf("Failed to detect host ips. Reason: %v", err))
			}
			ipset := sets.NewString()
			for _, ip := range ips {
				ipset.Insert(ip.String())
			}
			if !all && ipset.Len() > 0 {
				fmt.Print(ipset.List()[0])
			} else {
				fmt.Print(strings.Join(ipset.List(), ","))
			}
			os.Exit(0)
		},
	}
	cmd.Flags().BoolVar(&all, "all", all, "If true, returns all private ips.")
	return cmd
}
