## onessl wait-until-ready crd

Wait until a CRD is ready

### Synopsis

Wait until a CRD is ready

```
onessl wait-until-ready crd [flags]
```

### Options

```
  -h, --help                  help for crd
      --interval duration     Interval between checks (default 2s)
      --kube-context string   Name of kube context
      --kubeconfig string     Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --timeout duration      Timeout (default 3m0s)
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

* [onessl wait-until-ready](onessl_wait-until-ready.md)	 - Wait until resource is ready

