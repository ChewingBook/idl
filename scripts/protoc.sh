#!/bin/sh

set -e

IDL_PATH=/go/src/github.com/banksalad/idl

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

gen_swift() {
  mkdir -p "$IDL_PATH/gen/swift/"
  SWIFT_BUILT_IN_PROTO_FILE_PATH="/protos/google/protobuf"
  for FILE in $PROTOS; do
    if echo "$FILE" | grep -q "$SWIFT_BUILT_IN_PROTO_FILE_PATH"; then
      echo "[SWIFT EXCEPT FILE] $FILE"
      continue
    fi
    echo "[SWIFT] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --swift_opt=FileNaming=PathToUnderscores \
      --swift_opt=Visibility=Public \
      --swift_out=gen/swift/ \
      --plugin=protoc-gen-grpc-swift=/usr/bin/protoc-gen-grpc-swift \
      --grpc-swift_opt=Visibility=Public \
      --grpc-swift_opt=Server=false \
      --grpc-swift_opt=FileNaming=PathToUnderscores \
      --grpc-swift_out=gen/swift \
      "$IDL_PATH${FILE#.}"
  done
}

gen_python() {
  mkdir -p "$IDL_PATH/gen/python/"
  for FILE in $PROTOS; do
    echo "[PYTHON] $FILE"
    protoc-wrapper \
      -I$IDL_PATH/protos \
      -I/go/src \
      --python_out=gen/python/ \
      --grpc-python_out=gen/python/ \
      "$IDL_PATH${FILE#.}"
  done
}

gen_java() {
  mkdir -p "$IDL_PATH/gen/java/"
  JAVA_BUILT_IN_PROTO_FILE_PATH="/protos/google/rpc"
  for FILE in $PROTOS; do
    if echo "$FILE" | grep -q "$JAVA_BUILT_IN_PROTO_FILE_PATH"; then
      echo "[JAVA EXCEPT FILE] $FILE"
      continue
    fi
    echo "[JAVA] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --java_out=$IDL_PATH/gen/java \
      --plugin=protoc-gen-grpc-java=/usr/bin/protoc-gen-grpc-java \
      --grpc-java_out=$IDL_PATH/gen/java \
      "$IDL_PATH${FILE#.}"
  done
}

gen_javalite() {
  mkdir -p "$IDL_PATH/gen/javalite/"
  for FILE in $PROTOS; do
    echo "[JAVALITE] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --java_out=lite:$IDL_PATH/gen/javalite \
      --plugin=protoc-gen-grpc-java=/usr/bin/protoc-gen-grpc-java \
      --grpc-java_out=lite:$IDL_PATH/gen/javalite \
      "$IDL_PATH${FILE#.}"
  done
}

gen_go() {
  mkdir -p "$IDL_PATH/gen/go/"
  for FILE in $PROTOS; do
    echo "[GO] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --go_out=plugins=grpc:/go/src/ \
      --grpc-gateway_out=logtostderr=true:/go/src/ \
      "$IDL_PATH${FILE#.}"
  done
}

gen_swagger() {
  mkdir -p "$IDL_PATH/gen/swagger/"
  for FILE in $PROTOS; do
    echo "[SWAGGER] $FILE"
    protoc \
      -I$IDL_PATH/protos \
      -I/go/src \
      --swagger_out=logtostderr=true:gen/swagger/ \
      "$IDL_PATH${FILE#.}"
  done
}

# execute function by argument
# e.g. ./protoc.sh gen_python
$1