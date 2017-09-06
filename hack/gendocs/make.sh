#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/pre-k/hack/gendocs
go run main.go
popd
