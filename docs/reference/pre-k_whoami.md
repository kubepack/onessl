## pre-k whoami

Detect cloud provider

### Synopsis



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


```
pre-k whoami [flags]
```

### Options

```
  -h, --help   help for whoami
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
* [pre-k](pre-k.md)	 - Pre-k by AppsCode - Utilities for your cloud

