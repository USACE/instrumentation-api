#!/bin/bash

# convenience scipt for starting all containers
if [ "$1" = "up" ]; then
    env DOCKER_BUILDKIT=1 docker compose --profile='local' up -d --build;
elif [ "$1" = "down" ]; then
    docker compose --profile=local down;
elif [ "$1" = "clean" ]; then
    docker compose --profile=local down -v;
else
    echo -e "usage:\n\t./compose.sh up\n\t./compose.sh down\n\t./compose.sh clean"
fi
