package cmds

import (
	"fmt"
	"os"

	gv "github.com/JamesClonk/vultr/lib"
	"github.com/spf13/cobra"
)

func NewCmdVultrPrivateIP() *cobra.Command {
	var (
		token      string
		instanceID string
	)
	cmd := &cobra.Command{
		Use:               "private-ip",
		Short:             "Prints private IP of a Vultr host",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			client := gv.NewClient(token, nil)
			ips, err := client.ListIPv4(instanceID)
			if err != nil {
				Fatal(fmt.Errorf("failed to detect host ips. Reason: %v", err))
			}
			for _, ip := range ips {
				if ip.Type == "private" {
					fmt.Print(ip.IP)
					os.Exit(0)
				}
			}
			os.Exit(1)
		},
	}
	cmd.Flags().StringVar(&token, "token", "", "Vultr api token")
	cmd.Flags().StringVar(&instanceID, "instance-id", "", "Instance id of Vultr host")
	return cmd
}
