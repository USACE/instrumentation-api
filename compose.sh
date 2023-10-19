#!/bin/bash

if [ "$1" = "up" ]; then
    if [ "$2" = "mock" ]; then
        env DOCKER_BUILDKIT=1 docker-compose --profile=local --profile=mock up -d --build
    else
        env DOCKER_BUILDKIT=1 docker-compose --profile=local up -d --build
    fi

elif [ "$1" = "down" ]; then
    if [ "$2" = "mock" ]; then
        docker-compose --profile=local --profile=mock down
    else
        docker-compose --profile=local down
    fi

elif [ "$1" = "clean" ]; then
    if [ "$2" = "mock" ]; then
        docker-compose --profile=local --profile=mock down -v
    else
        docker-compose --profile=local down -v
    fi

elif [ "$1" = "test" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml build --build-arg dev
    docker-compose run -d alert
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml run --entrypoint="go test github.com/USACE/instrumentation-api/api/internal/handler" api

elif [ "$1" = "mkdocs" ]; then
    # TODO: this could possibly be added in CI, just run locally for now
    (cd api && swag init --pd -g cmd/core/main.go --parseInternal true --dir internal)

else
    echo -e "usage:\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean\n\t./compose.sh test\n\t./compose.sh mkdocs"
fi
