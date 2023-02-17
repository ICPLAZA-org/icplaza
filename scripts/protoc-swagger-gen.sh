#!/usr/bin/env bash

set -eo pipefail

SDK_VERSION=v0.45.4
ETHERMINT_VERSION=v0.15.0

ROOTPATH=/root/github.com/jonathan/thirds/icplaza

rm -rf ./tmp-swagger-gen ./tmp && mkdir -p ./tmp-swagger-gen ./tmp/proto ./tmp/third_party

cp -r ./proto ./tmp
cp -r ./third_party/proto ./tmp/third_party
cp -r ${ROOTPATH}/cosmos-sdk-icplaza@${SDK_VERSION}/proto ./tmp/third_party

proto_dirs=$(find ./tmp/proto ./tmp/third_party/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  # TODO: migrate to `buf build`

  if [[ $dir =~ "cosmos" ]]; then
      query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  fi
  if [[ ! -z "$query_file" ]]; then
    protoc  \
    -I "tmp/proto" \
    -I "tmp/third_party/proto" \
    "$query_file" \
    --swagger_out=./tmp-swagger-gen \
    --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
  fi
done

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen

# clean proto files
rm -rf ./tmp

# generate binary for static server
statik -src=./client/docs/swagger-ui -dest=./client/docs
