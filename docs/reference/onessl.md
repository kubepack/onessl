## onessl

onessl by AppsCode - Simple CLI to generate SSL certificates on any platform

### Synopsis

onessl by AppsCode - Simple CLI to generate SSL certificates on any platform

### Options

```
      --alsologtostderr                  log to standard error as well as files
      --as string                        Username to impersonate for the operation
      --as-group stringArray             Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string                 Default HTTP cache directory (default "$HOME/.kube/http-cache")
      --certificate-authority string     Path to a cert file for the certificate authority
      --client-certificate string        Path to a client certificate file for TLS
      --client-key string                Path to a client key file for TLS
      --cluster string                   The name of the kubeconfig cluster to use
      --context string                   The name of the kubeconfig context to use
  -h, --help                             help for onessl
      --insecure-skip-tls-verify         If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string                Path to the kubeconfig file to use for CLI requests.
      --log-backtrace-at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log-dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --match-server-version             Require server version to match client version
  -n, --namespace string                 If present, the namespace scope for this CLI request
      --request-timeout string           The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                    The address and port of the Kubernetes API server
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
      --token string                     Bearer token for authentication to the API server
      --user string                      The name of the kubeconfig user to use
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [onessl backup](onessl_backup.md)	 - Backup cluster objects
* [onessl base64](onessl_base64.md)	 - Base64 encode/decode input text
* [onessl create](onessl_create.md)	 - create PKI
* [onessl doctor](onessl_doctor.md)	 - Diagnoses Kubernetes Cluster Setup issues
* [onessl envsubst](onessl_envsubst.md)	 - Emulates bash environment variable substitution for input text
* [onessl get](onessl_get.md)	 - Get stuff
* [onessl has-keys](onessl_has-keys.md)	 - Checks configmap/secret has a set of given keys
* [onessl jsonpath](onessl_jsonpath.md)	 - Print value of jsonpath for input text
* [onessl semver](onessl_semver.md)	 - Print sanitized semver version
* [onessl version](onessl_version.md)	 - Prints binary version number.
* [onessl wait-until-has](onessl_wait-until-has.md)	 - Wait until has some attribute
* [onessl wait-until-ready](onessl_wait-until-ready.md)	 - Wait until resource is ready

