#!/bin/bash

if [ -z "$1" ]
  then
    echo "usage: ./build_ib.sh <tag>"
    exit 1
fi

(cd api && swag init --pd -g cmd/core/main.go --parseInternal true --dir internal)
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target core -t midas-api:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target telemetry -t midas-telemetry:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target alert -t midas-alert:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank" --target dcs-loader -t midas-dcs-loader:$1 api
docker build --build-arg="BASE_IMAGE=registry1.dso.mil/ironbank/flyway/flyway-docker:v8.5.9" -t midas-sql:$1 migrate

exit 0
