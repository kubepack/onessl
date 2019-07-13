#!/usr/bin/env bash

pushd $GOPATH/src/kubepack.dev/onessl/hack/gendocs
go run main.go
popd
