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
      --alsologtostderr                              log to standard error as well as files
      --analytics                                    Send analytical events to Google Guard (default true)
      --cloud-provider-gce-lb-src-cidrs cidrs        CIDRS opened in GCE firewall for LB traffic proxy & health checks (default 35.191.0.0/16,209.85.152.0/22,209.85.204.0/22,130.211.0.0/22)
      --default-not-ready-toleration-seconds int     Indicates the tolerationSeconds of the toleration for notReady:NoExecute that is added by default to every pod that does not already have such a toleration. (default 300)
      --default-unreachable-toleration-seconds int   Indicates the tolerationSeconds of the toleration for unreachable:NoExecute that is added by default to every pod that does not already have such a toleration. (default 300)
      --ir-data-source string                        Data source used by InitialResources. Supported options: influxdb, gcm. (default "influxdb")
      --ir-dbname string                             InfluxDB database name which contains metrics required by InitialResources (default "k8s")
      --ir-hawkular string                           Hawkular configuration URL
      --ir-influxdb-host string                      Address of InfluxDB which contains metrics required by InitialResources (default "localhost:8080/api/v1/namespaces/kube-system/services/monitoring-influxdb:api/proxy")
      --ir-namespace-only                            Whether the estimation should be made only based on data from the same namespace.
      --ir-password string                           Password used for connecting to InfluxDB (default "root")
      --ir-percentile int                            Which percentile of samples should InitialResources use when estimating resources. For experiment purposes. (default 90)
      --ir-user string                               User used for connecting to InfluxDB (default "root")
      --log_backtrace_at traceLocation               when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                               If non-empty, write log files in this directory
      --loglevel int                                 Log level (0 = DEBUG, 5 = FATAL) (default 1)
      --logtostderr                                  log to standard error instead of files
      --stderrthreshold severity                     logs at or above this threshold go to stderr (default 2)
  -v, --v Level                                      log level for V logs
      --vmodule moduleSpec                           comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [pre-k merge](pre-k_merge.md)	 - Merge Kubeadm config

