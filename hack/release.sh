#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

pushd "$(go env GOPATH)/src/github.com/pharmer/pre-k"
rm -rf dist
./hack/make.py build
./hack/make.py push
APPSCODE_ENV=prod ./hack/make.py push
./hack/make.py update_registry
popd
