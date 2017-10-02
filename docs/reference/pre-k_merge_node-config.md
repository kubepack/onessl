## pre-k merge node-config

Merge Kubeadm node configuration

### Synopsis


Merge Kubeadm node configuration

```
pre-k merge node-config [flags]
```

### Options

```
      --config string                                 Path to kubeadm config file
      --discovery-file string                         A file or url from which to load cluster information
      --discovery-token string                        A token used to validate cluster information fetched from the master
      --discovery-token-ca-cert-hash stringSlice      For token-based discovery, validate that the root CA public key matches this hash (format: "<type>:<value>").
      --discovery-token-unsafe-skip-ca-verification   For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.
  -h, --help                                          help for node-config
      --node-name string                              Specify the node name
      --tls-bootstrap-token string                    A token used for TLS bootstrapping
      --token string                                  Use this token for both discovery-token and tls-bootstrap-token
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

