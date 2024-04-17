#!/bin/bash

set -o pipefail

COMPOSECMD="env DOCKER_BUILDKIT=1 docker-compose -f docker-compose.yml --profile=local"
mkdocs() {
    (
        cd api && swag init --pd $1 -g cmd/core/main.go --parseInternal true --dir internal;
        find ./docs -type f -exec sed -i '' -e 's/github_com_USACE_instrumentation-api_api_internal_model.//g' {} \;
    )
}

if [ "$1" = "watch" ]; then
    mkdocs -q
    if [ "$2" = "mock" ]; then
        $COMPOSECMD -f docker-compose.dev.yml --profile=mock watch
    else
        $COMPOSECMD -f docker-compose.dev.yml watch
    fi

elif [ "$1" = "up" ]; then
    mkdocs -q
    if [ "$2" = "mock" ]; then
        $COMPOSECMD --profile=mock up -d --build
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
        docker-compose --profile=local --profile=mock down -v
    fi

elif [ "$1" = "mkdocs" ]; then
    mkdocs

else
    echo -e "usage:\n\t./compose.sh watch\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean\n\t./compose.sh test\n\t./compose.sh mkdocs"
fi
