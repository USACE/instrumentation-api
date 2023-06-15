#!/bin/bash

if [ -z "$1" ]
  then
    echo "usage: ./build_ib.sh <tag>"
    exit 1
fi

(cd api; docker build -f Dockerfile.ib -t midas-api:$1 .)
(cd sql; docker build -f Dockerfile.ib -t midas-sql:$1 .)
(cd telemetry; go mod vendor; docker build -f Dockerfile.ib -t midas-telemetry:$1 .)
(cd alert; go mod vendor; docker build -f Dockerfile.ib -t midas-alert:$1 .)
exit 0
