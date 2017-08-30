[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/cloudid)](https://goreportcard.com/report/github.com/appscode/cloudid)

[Website](https://appscode.com) • [Slack](https://slack.appscode.com) • [Twitter](https://twitter.com/AppsCodeHQ)

# CloudID
Detect Cloud Provider from Instance Metadata

## Motivation
Kubernetes has the concept of a [Cloud Provider](https://kubernetes.io/docs/getting-started-guides/scratch/#cloud-provider),
which is a module which provides an interface for managing TCP Load Balancers, Nodes (Instances) and Networking Routes.
This library can be used to identify cloud provider based on various instance metadata without requiring user input.

## Supported Cloud Providers
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

Please file an issue if you would have ideas to improve detection technique or add support for additional cloud providers. Pull requests are welcome.
