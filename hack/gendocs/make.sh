#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/onessl/hack/gendocs
go run main.go
popd
