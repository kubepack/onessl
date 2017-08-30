package cmds

import (
	"fmt"
	"os"

	"github.com/appscode/go/net"
	"github.com/spf13/cobra"
)

func NewCmdNodeIP() *cobra.Command {
	var ifaces []string
	cmd := &cobra.Command{
		Use:   "node-ip",
		Short: `Prints a IPv4 address for current host`,
		Long: `Prints a IPv4 address for current host for a given set of interface names. It always prefers a private IP over a public IP.
If no interface name is given, all interfaces are checked.`,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			_, ip, err := net.NodeIP(ifaces...)
			if err != nil {
				Fatal(fmt.Errorf("Failed to detect node ip. Reason: %v", err))
			}
			fmt.Print(ip.String())
			os.Exit(0)
		},
	}
	cmd.Flags().StringSliceVar(&ifaces, "ifaces", ifaces, "Comma separated list of interface names. If empty, all interfaces are checked.")
	return cmd
}
