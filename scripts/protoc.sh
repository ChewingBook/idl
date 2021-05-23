#!/bin/sh

set -e

IDL_PATH=/go/src/github.com/ChewingBook/idl

gen_typescript() {
  npm install ts-proto@1.79.8
  mkdir -p "$IDL_PATH/gen/typescript/"
  for FILE in $PROTOS; do
    echo "[TYPESCRIPT] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --plugin=protoc-gen-grpc=/usr/bin/protoc-gen-grpc-js \
      --plugin=./node_modules/.bin/protoc-gen-ts_proto \
      --ts_proto_out=gen/typescript \
      "$IDL_PATH${FILE#.}"
  done
  rm -rf node_modules package*.json
}

gen_go() {
  mkdir -p "$IDL_PATH/gen/go/"
  for FILE in $PROTOS; do
    echo "[GO] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --go_out=plugins=grpc:/go/src/ \
	  --go_out=$(pwd)/gen/go \
      --grpc-gateway_out=logtostderr=true:/go/src/ \
      "$IDL_PATH${FILE#.}"
  done
}

# execute function by argument
# e.g. ./protoc.sh gen_python
$1
