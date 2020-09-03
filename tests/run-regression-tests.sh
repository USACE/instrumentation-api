#!/bin/bash

newman run \
    -e ./postman_environment.local.json \
    ./hhd-regression.postman_collection.json