#!/usr/bin/env bash

pushd $GOPATH/src/github.com/kubepack/onessl/hack/gendocs
go run main.go
popd
