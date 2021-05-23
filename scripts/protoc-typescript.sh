#!/bin/sh

set -e

IDL_PATH=/go/src/github.com/ChewingBook/idl

gen_typescript() {
  npm install ts-proto@1.79.8
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
gen_typescript &
pid_gen_typescript=$!

wait $pid_gen_typescript
exit $?
