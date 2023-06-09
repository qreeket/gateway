# bin/sh

PROTOS_DIR=protos
OUT_DIR=gen

# Create generated directories
mkdir -p $OUT_DIR

# Generate Go protobufs for announcements
protoc -I=$PROTOS_DIR --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative -I=$OUT_DIR --go-grpc_out=$OUT_DIR --go_out=$OUT_DIR \
  $(find $PROTOS_DIR -name '*.proto')