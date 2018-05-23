[![Go Report Card](https://goreportcard.com/badge/github.com/kubepack/onessl)](https://goreportcard.com/report/github.com/kubepack/onessl)
[![Build Status](https://travis-ci.org/kubepack/onessl.svg?branch=master)](https://travis-ci.org/kubepack/onessl)
[![codecov](https://codecov.io/gh/kubepack/onessl/branch/master/graph/badge.svg)](https://codecov.io/gh/kubepack/onessl)
[![Slack](https://slack.appscode.com/badge.svg)](https://slack.appscode.com)
[![mailing list](https://img.shields.io/badge/mailing_list-join-blue.svg)](https://groups.google.com/forum/#!forum/kubepack)
[![Twitter](https://img.shields.io/twitter/follow/appscodehq.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=AppsCodeHQ)

# onessl
Simple CLI to generate SSL certificates on any platform. Just download the pre-built binaries from this project's release page and you are ready to go!


## Available Commands
* [onessl create](/docs/reference/onessl_create.md)	 - create PKI
  * [onessl create ca-cert](/docs/reference/onessl_create_ca-cert.md)	 - Create CA cert/key pair
  * [onessl create server-cert](/docs/reference/onessl_create_server-cert.md)	 - Generate server certificate pair
  * [onessl create client-cert](/docs/reference/onessl_create_client-cert.md)	 - Generate client certificate pair
* [onessl get](/docs/reference/onessl_get.md)	 - Get stuff
  * [onessl get ca-cert](/docs/reference/onessl_get_ca-cert.md)	 - Prints self-signed CA certificate from PEM encoded RSA private key
  * [onessl get kube-ca](/docs/reference/onessl_get_kube-ca.md)	 - Prints CA certificate for Kubernetes cluster from Kubeconfig
* [onessl base64](/docs/reference/onessl_base64.md)	 - Base64 encode/decode input text
* [onessl envsubst](/docs/reference/onessl_envsubst.md)	 - Emulates bash environment variable substitution for input text
* [onessl jsonpath](/docs/reference/onessl_jsonpath.md)	 - Print value of jsonpath for input text
* [onessl semver](/docs/reference/onessl_semver.md)	 - Print sanitized semver version


## Contribution guidelines
Want to help improve onessl? Please start [here](/CONTRIBUTING.md).

---

**onessl binary collects anonymous usage statistics to help us learn how the software is being used and how we can improve it. To disable stats collection, run the operator with the flag** `--analytics=false`.

---

## Support

We use Slack for public discussions. To chit chat with us or the rest of the community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C0XQFLGRM/details/) channel `#general`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

If you have found a bug with Voyager or want to request for new features, please [file an issue](https://github.com/appscode/voyager/issues/new).
