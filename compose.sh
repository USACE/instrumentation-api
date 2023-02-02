#!/bin/bash

if [ "$1" = "up" ]; then
    env DOCKER_BUILDKIT=1 docker-compose --profile=local up -d --build;
elif [ "$1" = "down" ]; then
    docker-compose --profile=local down;
elif [ "$1" = "clean" ]; then
    docker-compose --profile=local down -v;
elif [ "$1" = "test" ]; then
    docker-compose up -d --build;
    if [ "$REPORT" = true ]; then
        docker run \
            -v $(pwd)/tests:/etc/newman --network=instrumentation-api_default \
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
            -v $(pwd)/tests:/etc/newman --network=instrumentation-api_default \
            -t postman/newman run /etc/newman/instrumentation-regression.postman_collection.json \
            --environment=/etc/newman/postman_environment.docker-compose.json
    fi
    docker-compose down;
else
    echo -e "usage:\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean\n\t./compose.sh test"
fi
