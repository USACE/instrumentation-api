#!/bin/bash

# conveniently listen for compose logs, exit on ctrl-c

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "$parent_path"

trap "exit" SIGINT

while :
do
    echo "$(date +%Y-%m-%d_%H:%M:%S) Listening for compose logs... $1"
    docker-compose logs -f $1
    sleep 5
done
