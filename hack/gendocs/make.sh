#!/usr/bin/env bash

pushd $GOPATH/src/github.com/pharmer/pre-k/hack/gendocs
go run main.go
popd
