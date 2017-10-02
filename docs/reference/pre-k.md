## pre-k

Pre-k by AppsCode - Utilities for your cloud

### Synopsis


Pre-k by AppsCode - Utilities for your cloud

### Options

```
      --alsologtostderr                              log to standard error as well as files
      --analytics                                    Send analytical events to Google Guard (default true)
      --cloud-provider-gce-lb-src-cidrs cidrs        CIDRS opened in GCE firewall for LB traffic proxy & health checks (default 35.191.0.0/16,209.85.152.0/22,209.85.204.0/22,130.211.0.0/22)
      --default-not-ready-toleration-seconds int     Indicates the tolerationSeconds of the toleration for notReady:NoExecute that is added by default to every pod that does not already have such a toleration. (default 300)
      --default-unreachable-toleration-seconds int   Indicates the tolerationSeconds of the toleration for unreachable:NoExecute that is added by default to every pod that does not already have such a toleration. (default 300)
  -h, --help                                         help for pre-k
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
* [pre-k get](pre-k_get.md)	 - Get stuff
* [pre-k merge](pre-k_merge.md)	 - Merge Kubeadm config
* [pre-k version](pre-k_version.md)	 - Prints binary version number.
* [pre-k whoami](pre-k_whoami.md)	 - Detect cloud provider

