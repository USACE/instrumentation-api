#!/bin/bash

# Script to conveniently git pull and rebuild images and (re)start docker stack
git pull;

docker compose down;
docker compose -f ./docker-compose.swagger.yml down;

docker builder prune --force --filter "label=instrumentation-api";

docker compose up -d --build;
# suppress incorrect warning of orphan containers created in above line
COMPOSE_IGNORE_ORPHANS=1 docker compose -f ./docker-compose.swagger.yml up -d;
