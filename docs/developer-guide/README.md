## Development Guide
This document is intended to be the canonical source of truth for things like supported toolchain versions for building Pack.
If you find a requirement that this doc does not capture, please submit an issue on github.

This document is intended to be relative to the branch in which it is found. It is guaranteed that requirements will change over time
for the development branch, but release branches of Pre-k should not change.

### Build Pack
Some of the Pre-k development helper scripts rely on a fairly up-to-date GNU tools environment, so most recent Linux distros should
work just fine out-of-the-box.

#### Setup GO
Pre-k is written in Google's GO programming language. Currently, Pre-k is developed and tested on **go 1.8.3**. If you haven't set up a GO
development environment, please follow [these instructions](https://golang.org/doc/code.html) to install GO.

#### Download Source

```console
$ go get github.com/pharmer/pre-k
$ cd $(go env GOPATH)/src/github.com/pharmer/pre-k
```

#### Install Dev tools
To install various dev tools for Pack, run the following command:

```console
# setting up dependencies for compiling pre-k...
$ ./hack/builddeps.sh
```

#### Build Binary
```
$ ./hack/make.py
$ pre-k version
```

#### Dependency management
Pre-k uses [Glide](https://github.com/Masterminds/glide) to manage dependencies. Dependencies are already checked in the `vendor` folder.
If you want to update/add dependencies, run:
```console
$ glide slow
```

#### Generate CLI Reference Docs
```console
$ ./hack/gendocs/make.sh
```
