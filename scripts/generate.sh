#!/bin/bash

set -e

IDL_PATH=/go/src/github.com/ChewingBook/idl
DOCKER_PROTOC_VERSION=thethingsindustries/protoc:3.1.30

docker run --rm -i --name protoc \
  -v "$(pwd)":$IDL_PATH -w $IDL_PATH \
  --env PROTOS="$(find ./protos -name '*.proto')" \
  --entrypoint /bin/sh \
  $DOCKER_PROTOC_VERSION $IDL_PATH/scripts/protoc-typescript.sh
