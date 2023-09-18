#!/bin/bash

if [ -z "$1" ]
  then
    echo "usage: ./build_ib.sh <tag>"
    exit 1
fi

(
cd api
docker build --file Dockerfile.ib --target core --tag midas-api:$1 .
docker build --file Dockerfile.ib --target telemetry --tag midas-telemetry:$1 .
docker build --file Dockerfile.ib --target alert --tag midas-alert:$1 .
)
(
cd migrate
docker build --file Dockerfile.ib --tag midas-sql:$1 .
)
exit 0
