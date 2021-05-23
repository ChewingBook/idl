#!/bin/bash

set -e

IDL_PATH=/go/src/github.com/ChewingBook/idl
if [[ -z "$PROTOS" ]]; then
  PROTOS=$(find ./protos -name '*.proto')
fi
DOCKER_PROTOC=$(cat "$(dirname "$0")"/docker-protoc-version)

docker run --rm --name protoc-go \
  -v "$(pwd)":$IDL_PATH -w $IDL_PATH \
  --env PROTOS="$PROTOS" \
  --entrypoint /bin/sh \
  "$DOCKER_PROTOC" $IDL_PATH/scripts/protoc.sh gen_go &
pid_gen_go=$!

wait $pid_gen_go 
exit $?
  
