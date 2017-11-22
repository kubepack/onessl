## pre-k machine swapoff

Permanently disabled swap disks on a machine

### Synopsis


Runs "swapoff -a" and comments out all swap devices on a machine found in /etc/fstab file.

```
pre-k machine swapoff [flags]
```

### Options

```
      --filename string   Path to volume list file. (default "/etc/fstab")
  -h, --help              help for swapoff
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

