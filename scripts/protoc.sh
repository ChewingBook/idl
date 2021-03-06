#!/bin/sh

set -e

IDL_PATH=$(pwd)

if [[ -z "$PROTOS" ]]; then
  PROTOS=$(find ./protos -name '*.proto')
fi


gen_typescript() {
  npm install ts-proto@1.81.3
  mkdir -p "$IDL_PATH/gen/typescript/"
  for FILE in $PROTOS; do
    echo "[TYPESCRIPT] $FILE"
    protoc \
      -I$IDL_PATH/protos \
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
	  --go_out=$IDL_PATH/gen/go\
	  --go_opt=paths=source_relative\
	  --go-grpc_out=$IDL_PATH/gen/go\
	  --go-grpc_opt=paths=source_relative\
	  "$IDL_PATH${FILE#.}"
  done
}

# execute function by argument
# e.g. ./protoc.sh gen_python
$1
