#!/bin/bash
set -e
export GOPATH=`pwd`
if [[ ! $(which go-bindata) ]]; then
	go install github.com/jteeuwen/go-bindata/...
fi
if [[ ! -d src ]]; then
	mkdir -p src/github.com/gen2brain
	pushd src/github.com/gen2brain
	ln -s ../../../ acra-go
	popd
fi
if [[ ! -d bin ]]; then
	mkdir bin
fi
pushd server
../bin/go-bindata -o bindata.go -pkg server assets/... app/...
popd
go build -o bin/acra-go src/github.com/gen2brain/acra-go/acra-go/main_leveldb.go
