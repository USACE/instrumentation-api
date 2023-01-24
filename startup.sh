#!/bin/bash

# Script to conveniently git pull and rebuild images and (re)start docker stack
git pull;
docker compose --profile=local down;
docker builder prune --force --filter "label=instrumentation-api";
docker compose --profile=local up -d --build;
