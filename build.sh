#!/usr/bin/env bash

RUN_NAME="metricserver"

mkdir -p output/bin output/conf
cp conf/* output/conf
cp script/* output/
cat script/bootstrap.sh | sed 's/{{RUN_NAME}}/'${RUN_NAME}'/g' > output/bootstrap.sh
chmod +x output/bootstrap.sh

if [ "${IS_TEST_ENV}" != "1" ]; then
  go build -ldflags "-s -w" -o output/bin/${RUN_NAME}
else
  go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi