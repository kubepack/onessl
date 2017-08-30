## cloudid get node-ip

Prints a IPv4 address for current host

### Synopsis


Prints a IPv4 address for current host for a given set of interface names. It always prefers a private IP over a public IP.
If no interface name is given, all interfaces are checked.

```
cloudid get node-ip [flags]
```

### Options

```
  -h, --help                 help for node-ip
      --ifaces stringSlice   Comma separated list of interface names. If empty, all interfaces are checked.
```

### Options inherited from parent commands

```
      --analytics   Send analytical events to Google Guard (default true)
```

### SEE ALSO
* [cloudid get](cloudid_get.md)	 - Get stuff

