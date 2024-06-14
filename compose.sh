#!/bin/bash

set -o pipefail

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

COMPOSECMD="DOCKER_BUILDKIT=1 docker-compose -f docker-compose.yml"
mkdocs() {
    (
        DOCKER_BUILDKIT=1 docker build --file api/Dockerfile.openapi --output api/internal/server/docs api
        cd report && npm run generate >/dev/null;
    )
}

if [ "$1" = "watch" ]; then
    mkdocs -q
    if [ "$2" = "mock" ]; then
        $COMPOSECMD -f docker-compose.dev.yml --profile=mock watch
    elif [ "$2" = "auth" ]; then
        AUTH_JWT_MOCKED=false $COMPOSECMD -f docker-compose.dev.yml watch
    else
        $COMPOSECMD -f docker-compose.dev.yml watch
    fi

elif [ "$1" = "up" ]; then
    mkdocs -q
    if [ "$2" = "mock" ]; then
        $COMPOSECMD --profile=mock up -d --build
    elif [ "$2" = "auth" ]; then
        AUTH_JWT_MOCKED=false $COMPOSECMD up -d --build
    else
        $COMPOSECMD up -d --build
    fi

elif [ "$1" = "down" ]; then
    mkdocs -q
    $COMPOSECMD --profile=mock down

elif [ "$1" = "clean" ]; then
    $COMPOSECMD --profile=mock down -v

elif [ "$1" = "test" ]; then
    docker-compose build
    shift

    TEARDOWN=false
    REST_ARGS=()

    while [[ $# -gt 0 ]]; do
        case $1 in
            -rm)
                TEARDOWN=true
                shift
                ;;
            *)
                REST_ARGS+=("$1")
                shift
                ;;
        esac
    done

    GOCMD="go test ${REST_ARGS[@]} github.com/USACE/instrumentation-api/api/internal/handler"

    if [ "$REPORT" = true ]; then
        docker-compose run --entrypoint="$GOCMD" api > $(pwd)/test.log
    else
        docker-compose run --entrypoint="$GOCMD" api
    fi

    if [ $TEARDOWN = true ]; then
        docker-compose --profile=mock down -v
    fi

elif [ "$1" = "mkdocs" ]; then
    mkdocs

else
    echo -e "usage:\n\t./compose.sh watch\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean\n\t./compose.sh test\n\t./compose.sh mkdocs"
fi
