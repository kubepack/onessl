## pre-k merge master-config

Merge Kubeadm master configuration

### Synopsis

Merge Kubeadm master configuration

```
pre-k merge master-config [flags]
```

### Options

```
      --apiserver-advertise-address string      The IP address the API Server will advertise it's listening on. 0.0.0.0 means the default network interface's address.
      --apiserver-bind-port int32               Port for the API Server to bind to
      --apiserver-cert-extra-sans stringSlice   Optional extra altnames to use for the API Server serving cert. Can be both IP addresses and dns names.
      --cert-dir string                         The path where to save and store the certificates
      --config string                           Path to kubeadm config file (WARNING: Usage of a configuration file is experimental)
      --feature-gates string                    A set of key=value pairs that describe feature gates for various features. Options are:
CoreDNS=true|false (ALPHA - default=false)
DynamicKubeletConfig=true|false (ALPHA - default=false)
SelfHosting=true|false (ALPHA - default=false)
StoreCertsInSecrets=true|false (ALPHA - default=false)
  -h, --help                                    help for master-config
      --kubernetes-version string               Choose a specific Kubernetes version for the control plane
      --node-name string                        Specify the node name
      --pod-network-cidr string                 Specify range of IP addresses for the pod network; if set, the control plane will automatically allocate CIDRs for every node
      --service-cidr string                     Use alternative range of IP address for service VIPs
      --service-dns-domain string               Use alternative domain for services, e.g. "myorg.internal"
      --token string                            The token to use for establishing bidirectional trust between nodes and masters.
      --token-ttl duration                      The duration before the bootstrap token is automatically deleted. 0 means 'never expires'.
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Guard (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [pre-k merge](pre-k_merge.md)	 - Merge Kubeadm config

