## Development Guide
This document is intended to be the canonical source of truth for things like supported toolchain versions for building onessl.
If you find a requirement that this doc does not capture, please submit an issue on github.

This document is intended to be relative to the branch in which it is found. It is guaranteed that requirements will change over time
for the development branch, but release branches of onessl should not change.

### Build onessl
Some of the onessl development helper scripts rely on a fairly up-to-date GNU tools environment, so most recent Linux distros should
work just fine out-of-the-box.

#### Setup GO
onessl is written in Google's GO programming language. Currently, onessl is developed and tested on **go 1.9.3**. If you haven't set up a GO
development environment, please follow [these instructions](https://golang.org/doc/code.html) to install GO.

#### Download Source

```console
$ go get kubepack.dev/onessl
$ cd $(go env GOPATH)/src/kubepack.dev/onessl
```

#### Install Dev tools
To install various dev tools for onessl, run the following command:

```console
# setting up dependencies for compiling onessl...
$ ./hack/builddeps.sh
```

#### Build Binary
```
$ ./hack/make.py
$ onessl version
```

#### Dependency management
onessl uses [Glide](https://github.com/Masterminds/glide) to manage dependencies. Dependencies are already checked in the `vendor` folder.
If you want to update/add dependencies, run:

```console
$ glide slow
```

#### Generate CLI Reference Docs
```console
$ ./hack/gendocs/make.sh
```
