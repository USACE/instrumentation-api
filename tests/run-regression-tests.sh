#!/bin/bash

newman run \
    -e ./postman_environment.local.json \
    ./instrumentation-regression.postman_collection.json