#!/bin/bash

if [ -z "$1" ]
  then
    echo "usage: ./build_ib.sh <tag>"
    exit 1
fi

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

(./compose.sh mkdocs)

for BUILD_TARGET in midas-api midas-telemetry midas-alert midas-dcs-loader
do
  docker build \
    --platform=linux/amd64 \
    --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" \
    --build-arg="GO_VERSION=1.22" \
    --build-arg="BUILD_TAG=$1" \
    --build-arg="BUILD_TARGET=${BUILD_TARGET}" \
    -t $BUILD_TARGET:"$1" api
done

docker build \
  --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/opensource/alpinelinux/alpine:3.19.1" \
  -t midas-report:$1 report

docker build \
  --platform=linux/amd64 \
  --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/flyway/flyway-docker:v10.9.1" \
  -t midas-sql:$1 migrate

exit 0
