[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/pre-k)](https://goreportcard.com/report/github.com/appscode/pre-k)

# Pre-k
Pre Kubeadm


## Available Commands
* [pre-k create](/docs/reference/pre-k_create.md)	 - create PKI
  * [pre-k create ca-cert](/docs/reference/pre-k_create_ca-cert.md)	 - Create CA cert/key pair
  * [pre-k create client-cert](/docs/reference/pre-k_create_client-cert.md)	 - Generate client certificate pair
  * [pre-k create server-cert](/docs/reference/pre-k_create_server-cert.md)	 - Generate server certificate pair
* [pre-k get](/docs/reference/pre-k_get.md)	 - Get stuff
  * [pre-k get ca-cert](/docs/reference/pre-k_get_ca-cert.md)	 - Prints self-sgned CA certificate from PEM encoded RSA private key
  * [pre-k get linode-hostname](/docs/reference/pre-k_get_linode-hostname.md)	 - Prints hostname based on public IP for current Linode host
  * [pre-k get node-ip](/docs/reference/pre-k_get_node-ip.md)	 - Prints a IPv4 address for current host
  * [pre-k get private-ips](/docs/reference/pre-k_get_private-ips.md)	 - Prints private ip(s) for current host
  * [pre-k get public-ips](/docs/reference/pre-k_get_public-ips.md)	 - Prints public ip(s) for current host
* [pre-k merge](/docs/reference/pre-k_merge.md)	 - Merge Kubeadm config
  * [pre-k merge master-config](/docs/reference/pre-k_merge_master-config.md)	 - Merge Kubeadm master configuration
  * [pre-k merge node-config](/docs/reference/pre-k_merge_node-config.md)	 - Merge Kubeadm node configuration
* [pre-k whoami](/docs/reference/pre-k_whoami.md)	 - Detect cloud provider


## Contribution guidelines
Want to help improve Guard? Please start [here](/CONTRIBUTING.md).

---

**The pre-k binary collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support
If you have any questions, you can reach out to us.
* [Slack](https://slack.appscode.com)
* [Twitter](https://twitter.com/AppsCodeHQ)
