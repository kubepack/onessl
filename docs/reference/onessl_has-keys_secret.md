## onessl has-keys secret

Check a secret has a set of given keys

### Synopsis

Check a secret has a set of given keys

```
onessl has-keys secret [flags]
```

### Options

```
  -h, --help                  help for secret
      --keys stringSlice      Keys to search for
      --kube-context string   Name of kube context
      --kubeconfig string     Path to kubeconfig file with authorization information (the master location is set by the master flag).
  -n, --namespace string      Namespace used to search for deployment (default "default")
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

* [onessl has-keys](onessl_has-keys.md)	 - Checks configmap/secret has a set of given keys

