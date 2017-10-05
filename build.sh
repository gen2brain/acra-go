#!/bin/bash
#
# This build file will build acra-go by creating a src/ directory that 
# symlinks the github/gen2brain/acra-go directory back to this directory,
# and then does all the required go-installs to get the necessary packages
# for vendor management, and the other required tools.
# It's rather hacky, but it works.
set -e
export GOPATH=`pwd`
if [[ ! -d src ]]; then
	mkdir -p src/github.com/gen2brain
	pushd src/github.com/gen2brain
	ln -s ../../../ acra-go
	popd
fi
if [[ ! -d bin ]]; then
	mkdir bin
fi
if [[ ! -x bin/go-bindata ]]; then
	go get -u github.com/jteeuwen/go-bindata/...
fi
if [[ ! -x bin/govendor ]]; then
	go get -u github.com/kardianos/govendor/...
	pushd src/github.com/gen2brain/acra-go
	bin/govendor sync
	popd
fi
pushd server
../bin/go-bindata -o bindata.go -pkg server assets/... app/...
popd
go build -o bin/acra-go src/github.com/gen2brain/acra-go/acra-go/main_leveldb.go
