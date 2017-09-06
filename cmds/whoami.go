package cmds

import (
	"fmt"

	"github.com/appscode/pre-k/internal"
	"github.com/spf13/cobra"
)

func NewCmdWhoAmI() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "whoami",
		Short:             "Detect cloud provider",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(Detect())
		},
	}

	return cmd
}

func Detect() string {
	done := make(chan string)
	go internal.DetectAWS(done)
	go internal.DetectGCE(done)
	go internal.DetectDigitalOcean(done)
	go internal.DetectAzure(done)
	go internal.DetectVultr(done)
	go internal.DetectLinode(done)
	go internal.DetectSoftlayer(done)
	go internal.DetectScaleway(done)

	n := 8 // total number of go routines
	i := 0
	provider := ""
	for ; i < n; i++ {
		p := <-done
		if p != provider {
			provider = p
			break
		}
	}
	if i < n {
		// run drainer
		go func() {
			for ; i < n; i++ {
				<-done
			}
		}()
	}
	return provider
}
