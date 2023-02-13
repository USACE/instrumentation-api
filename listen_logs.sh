#!/bin/bash

# conveniently listen for compose logs, exit on ctrl-c

trap "exit" SIGINT

while :
do
    echo "$(date +%Y-%m-%d_%H:%M:%S) Listening for compose logs... $1"
    docker-compose logs -f $1
    sleep 5
done
