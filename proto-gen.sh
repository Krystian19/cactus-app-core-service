#!/bin/bash

# Generates the respective protobuf go files
# Installs golang https://github.com/stefanmaric/g
curl -sSL https://git.io/g-install | sh -s -- bash -y
source ~/.bashrc
g install 1.13.8
go get -u github.com/golang/protobuf/protoc-gen-go

# Generate Go protobuf files
protoc --go_out=plugins=grpc,paths=source_relative:. proto/*.proto