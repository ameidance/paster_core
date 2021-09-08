#!/usr/bin/env bash
RUN_NAME="paster_core"

mkdir -p output/bin output/conf
cp script/* output/
cp conf/* output/conf/
chmod +x output/bootstrap.sh

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -o output/bin/${RUN_NAME}
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi