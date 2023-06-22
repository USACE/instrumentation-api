#!/bin/bash

if [ "$1" = "up" ]; then
    (cd telemetry; go mod vendor);
    (cd alert; go mod vendor);
    if [ "$2" = "mock" ]; then
        env DOCKER_BUILDKIT=1 docker-compose --profile=local --profile=mock up -d --build;
    else
        env DOCKER_BUILDKIT=1 docker-compose --profile=local up -d --build;
    fi
elif [ "$1" = "down" ]; then
    if [ "$2" = "mock" ]; then
        docker-compose --profile=local --profile=mock down;
    else
        docker-compose --profile=local down;
    fi
elif [ "$1" = "clean" ]; then
    if [ "$2" = "mock" ]; then
        docker-compose --profile=local --profile=mock down -v;
    else
        docker-compose --profile=local down -v;
    fi
elif [ "$1" = "test" ]; then
    (cd telemetry; go mod vendor);
    (cd alert; go mod vendor);
    docker-compose up -d --build;
    if [ "$REPORT" = true ]; then
        docker run \
            -v $(pwd)/tests/postman:/etc/newman --network=instrumentation-api_default \
            --rm \
            --entrypoint /bin/sh \
            -t postman/newman \
            -c "npm i -g newman newman-reporter-htmlextra; \
                newman run /etc/newman/instrumentation-regression.postman_collection.json \
                --environment=/etc/newman/postman_environment.docker-compose.json \
                --reporter-htmlextra-browserTitle 'Instrumentation' \
                --reporter-htmlextra-title 'Instrumentation Regression Tests' \
                --reporter-htmlextra-titleSize 4 \
                -r htmlextra --reporter-htmlextra-export /etc/newman/instrumentation.html"
    else
        docker run \
            --rm \
            -v $(pwd)/tests/postman:/etc/newman --network=instrumentation-api_default \
            -t postman/newman run /etc/newman/instrumentation-regression.postman_collection.json \
            --environment=/etc/newman/postman_environment.docker-compose.json
    fi
    docker-compose down;
else
    echo -e "usage:\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean\n\t./compose.sh test"
fi
