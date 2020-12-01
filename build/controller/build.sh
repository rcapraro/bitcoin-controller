#!/bin/bash
cd $GOPATH/src/io.saagie/bitcoin-controller/cmd/controller
CGO_ENABLED=0 go build
docker login
docker build --rm -f ../../build/controller/Dockerfile -t ubik74/bitcoin-controller:latest .
docker push ubik74/bitcoin-controller:latest