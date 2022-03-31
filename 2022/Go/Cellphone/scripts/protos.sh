#! /usr/bin/bash

set -e

CUR_DIR=$(pwd)
BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

if [[ ! $(which protoc) ]]; then
    if [[ ! $(which brew) ]]; then
        log "ERROR" "protoc and brew aren't in path. Exiting..."
        return 1
    fi

    brew install protobuf
fi

if [[ ! $(which protoc-gen-go) ]]; then
  go install github.com/golang/protobuf/protoc-gen-go
fi
if [[ ! $(which protoc-gen-go-grpc) ]]; then
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
fi
if [[ ! $(which protoc-gen-grpc-gateway) ]]; then
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
fi
if [[ ! $(which protoc-gen-openapiv2) ]]; then
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
fi

echo "[protos.sh] Generating Proto Files"

protoc -I=./ \
  -I=${BASE_DIR}/protos \
  --go_out=paths=source_relative:${BASE_DIR}/protos/go \
  ${BASE_DIR}/protos/*.proto

cd $CUR_DIR

echo "[protos.sh] Done"
