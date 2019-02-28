#!/usr/bin/env bash

source ../scripts/platform.sh

rootDir=$(getRootDir)

protoc \
    -I ${rootDir}/protobuf \
    --go_out=plugins=grpc:${rootDir}/internal/pkg/pb  \
    imgtrip.proto

protoc_go_inject_tag=$(suffix "${rootDir}/build/bin/protoc-go-inject-tag")

${protoc_go_inject_tag} -input=${rootDir}/internal/pkg/pb/imgtrip.pb.go
