#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/cloudid/hack/gendocs
go run main.go
popd
