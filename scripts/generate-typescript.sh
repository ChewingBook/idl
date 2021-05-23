#!/bin/bash

set -e

IDL_PATH=/go/src/github.com/ChewingBook/idl
if [[ -z "$PROTOS" ]]; then
  PROTOS=$(find ./protos -name '*.proto')
fi
DOCKER_PROTOC=$(cat "$(dirname "$0")"/docker-protoc-version)

docker run --rm -i --name protoc \
  -v "$(pwd)":$IDL_PATH -w $IDL_PATH \
  --env PROTOS="$(find ./protos -name '*.proto')" \
  --entrypoint /bin/sh \
  $DOCKER_PROTOC $IDL_PATH/scripts/protoc.sh gen_typescript 
