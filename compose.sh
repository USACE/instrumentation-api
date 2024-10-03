#!/bin/bash

set -Eeo pipefail

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

COMPOSECMD="docker compose -f docker-compose.yml"

mkdocs() {
    (
        DOCKER_BUILDKIT=1 docker build --file api/Dockerfile.openapi --output api/internal/server/docs api
        cd report && npm run generate >/dev/null;
    )
}


if [ "$1" = "watch" ]; then
    mkdocs -q
    if [ "$2" = "mock" ]; then
        DOCKER_BUILDKIT=1 $COMPOSECMD -f docker-compose.dev.yml --profile=mock watch
    else
        DOCKER_BUILDKIT=1 $COMPOSECMD -f docker-compose.dev.yml watch
    fi


elif [ "$1" = "up" ]; then
    mkdocs -q
    if [ "$2" = "mock" ]; then
        DOCKER_BUILDKIT=1 $COMPOSECMD --profile=mock up -d --build
    else
        DOCKER_BUILDKIT=1 $COMPOSECMD up -d --build
    fi
    

elif [ "$1" = "build" ]; then
    if [ "$2" = "local" ] || [ "$2" = "develop" ] || [ "$2" = "test" ] || [ "$2" = "prod" ]; then
        SCRATCH_BASE_IMAGE=scratch
        ALPINE_BASE_IMAGE=alpine:3.19
        FLYWAY_BASE_IMAGE=flyway/flyway:10.9.1

        if [ "$2" = "test" ] || [ "$2" = "prod" ]; then
            SCRATCH_BASE_IMAGE=registry1.dso.mil/ironbank/docker/scratch:ironbank
            ALPINE_BASE_IMAGE=registry1.dso.mil/ironbank/opensource/alpinelinux/alpine:3.19.1
            FLYWAY_BASE_IMAGE=registry1.dso.mil/ironbank/flyway/flyway-docker:v10.9.1
            AMD64_TARGET_PLATFORM=true
        fi

        for BUILD_TARGET in midas-api midas-sql midas-telemetry midas-alert midas-dcs-loader
        do
          docker build \
            ${AMD64_TARGET_PLATFORM:+--platform=linux/amd64} \
            --build-arg="BASE_IMAGE=${SCRATCH_BASE_IMAGE}" \
            --build-arg="GO_VERSION=1.23" \
            --build-arg="BUILD_TAG=$2" \
            --build-arg="BUILD_TARGET=${BUILD_TARGET}" \
            -t $BUILD_TARGET:"$2" api
        done

        docker build \
          --build-arg="BASE_IMAGE=${ALPINE_BASE_IMAGE}" \
          -t midas-report:$2 report
    else
        echo -e "usage:\n\t./compose.sh build [local,develop,test,prod]"
        exit 1
    fi

    if [ "$3" = "push" ]; then
        if [ -z "$4" ]; then
            echo -e "usage:\n\t./compose.sh build [local,develop,test,prod] push <image_registry>"
            exit 1
        fi

        declare -a REGISTRIES=("midas-api" "midas-telemetry" "midas-alert" "midas-dcs-loader" "midas-sql")

        # tag
        for IMAGE in "${REGISTRIES[@]}"
        do
            docker tag $IMAGE:"$2" $4/$IMAGE:"$2"
        done
        if [ "$2" = "develop" ]; then
            docker tag midas-report:"$2" $4/midas-report:"$2"
        fi

        # push
        for IMAGE in "${REGISTRIES[@]}"
        do
            docker push $4/$IMAGE:"$2"
        done
        if [ "$2" = "develop" ]; then
            docker push $4/midas-report:"$2"
        fi
    fi


elif [ "$1" = "authdbdump" ]; then
    $COMPOSECMD exec authdb pg_dump postgres > auth/initdb/init2.sql


elif [ "$1" = "down" ]; then
    mkdocs -q
    $COMPOSECMD -f docker-compose.dev.yml --profile=mock down


elif [ "$1" = "clean" ]; then
    $COMPOSECMD -f docker-compose.dev.yml --profile=mock down -v


elif [ "$1" = "test" ]; then
    docker compose build
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
        docker compose run -e INSTRUMENTATION_AUTH_JWT_MOCKED=true --entrypoint="$GOCMD" api > $(pwd)/test.log
    else
        docker compose run -e INSTRUMENTATION_AUTH_JWT_MOCKED=true --entrypoint="$GOCMD" api
    fi

    if [ $TEARDOWN = true ]; then
        docker compose --profile=mock down -v
    fi


elif [ "$1" = "mkdocs" ]; then
    mkdocs


else
    echo -e "usage:\n\t./compose.sh watch\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean\n\t./compose.sh test\n\t./compose.sh mkdocs"
fi
