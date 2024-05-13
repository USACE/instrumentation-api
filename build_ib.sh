#!/bin/bash

if [ -z "$1" ]
  then
    echo "usage: ./build_ib.sh <tag>"
    exit 1
fi

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

(./compose.sh mkdocs)
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target core -t midas-api:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target telemetry -t midas-telemetry:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target alert -t midas-alert:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target dcs-loader -t midas-dcs-loader:$1 api
# TODO add report node serice build
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/flyway/flyway-docker:v10.9.1" -t midas-sql:$1 migrate

exit 0
