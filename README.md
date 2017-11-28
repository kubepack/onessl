[![Go Report Card](https://goreportcard.com/badge/github.com/appscode/pre-k)](https://goreportcard.com/report/github.com/appscode/pre-k)

# Pre-k
A set of handy commands that you run to prepare a host for `kubeadm`. This is intended to used with [Pharmer by AppsCode](https://appscode.com/products/pharmer).


## Supported Versions
Kubernetes 1.8+


## Available Commands
* [pre-k create](/docs/reference/pre-k_create.md)	 - create PKI
  * [pre-k create ca-cert](/docs/reference/pre-k_create_ca-cert.md)	 - Create CA cert/key pair
  * [pre-k create client-cert](/docs/reference/pre-k_create_client-cert.md)	 - Generate client certificate pair
  * [pre-k create server-cert](/docs/reference/pre-k_create_server-cert.md)	 - Generate server certificate pair
* [pre-k get](/docs/reference/pre-k_get.md)	 - Get stuff
  * [pre-k get ca-cert](/docs/reference/pre-k_get_ca-cert.md)	 - Prints self-sgned CA certificate from PEM encoded RSA private key
* [pre-k linode](/docs/reference/pre-k_linode.md)  - linode commands
  * [pre-k linode hostname](/docs/reference/pre-k_linode_hostname.md)	 - Prints hostname based on public IP for current Linode host
* [pre-k machine](/docs/reference/pre-k_machine.md)  - machine commands
  * [pre-k machine cloud-provider](/docs/reference/pre-k_machine_cloud-provider.md)  - Detect cloud provider
  * [pre-k machine node-ip](/docs/reference/pre-k_machine_node-ip.md)  - Prints a IPv4 address for current host
  * [pre-k machine private-ips](/docs/reference/pre-k_machine_private-ips.md)  - Prints private ip(s) for current host
  * [pre-k machine public-ips](/docs/reference/pre-k_machine_public-ips.md)  - Prints public ip(s) for current host
  * [pre-k machine swapoff](/docs/reference/pre-k_machine_swapoff.md)  - Permanently disabled swap disks on a machine
* [pre-k merge](/docs/reference/pre-k_merge.md)	 - Merge Kubeadm config
  * [pre-k merge master-config](/docs/reference/pre-k_merge_master-config.md)	 - Merge Kubeadm master configuration
  * [pre-k merge node-config](/docs/reference/pre-k_merge_node-config.md)	 - Merge Kubeadm node configuration
* [pre-k vultr](/docs/reference/pre-k_vultr.md)  - vultr commands
  * [pre-k vultr private-ip](/docs/reference/pre-k_vultr_private-ip.md)  - Prints private IP of a Vultr host


## Contribution guidelines
Want to help improve pre-k? Please start [here](/CONTRIBUTING.md).

## Support
If you have any questions, [file an issue](https://github.com/appscode/pharmer/issues/new) or talk to us on the [Kubernetes Slack team](http://slack.kubernetes.io/) channel `#pharmer`.

---

**The pre-k binary collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---
