package cmds

import (
	"fmt"

	"github.com/appscode/pre-k/internal"
	"github.com/spf13/cobra"
)

func NewCmdCloudProvider() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cloud-provider",
		Short: "Detect cloud provider",
		Long: `
Kubernetes has the concept of a [Cloud Provider](https://kubernetes.io/docs/getting-started-guides/scratch/#cloud-provider),
which is a module which provides an interface for managing TCP Load Balancers, Nodes (Instances) and Networking Routes.
This library can be used to identify cloud provider based on various instance metadata without requiring user input.

**Supported Cloud Providers**

| Id          | Name                  | Technique                                                                                                          |
|-------------|-----------------------|--------------------------------------------------------------------------------------------------------------------|
|aws          | Amazon Web Services   | [Instance Identity Documents](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html) |
|azure        | Microsoft Azure       | [Instance info](https://azure.microsoft.com/en-us/blog/what-just-happened-to-my-vm-in-vm-metadata-service/) |
|digitalocean | DigitalOcan           | [Droplet metadata](https://developers.digitalocean.com/documentation/metadata/#metadata-in-json) |
|gce          | Google Cloud Platform | [GCE Instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#endpoints) |
|linode       | Linode                | Reverse domain name(PTR record) |
|scaleway     | Scaleway              | [Instance user data](https://github.com/scaleway/initrd/issues/84) |
|softlayer    | IBM Softlayer(Bluemix)| [Instance user metadata](https://github.com/bodenr/cci/wiki/SL-user-metadata) |
|vultr        | Vultr                 | [Instance metadata](https://www.vultr.com/metadata/) |
`,
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
