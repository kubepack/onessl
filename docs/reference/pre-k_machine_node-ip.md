## pre-k machine node-ip

Prints a IPv4 address for current host

### Synopsis


Prints a IPv4 address for current host for a given set of interface names. It always prefers a private IP over a public IP.
If no interface name is given, all interfaces are checked.

```
pre-k machine node-ip [flags]
```

### Options

```
  -h, --help                 help for node-ip
      --ifaces stringSlice   Comma separated list of interface names. If empty, all interfaces are checked.
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
* [pre-k machine](pre-k_machine.md)	 - machine commands

