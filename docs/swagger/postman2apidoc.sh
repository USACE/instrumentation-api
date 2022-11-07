#!/bin/bash

# Runs in swagger_init container, converts most recent postman collection to apidoc
if ({
    (npm list -g postman-to-openapi || npm i postman-to-openapi -g)

    p2o /tests/instrumentation-regression.postman_collection.json \
        -f /docs/swagger/apidoc.json \
        -o /docs/swagger/postman-compose.env.json;
} &> /dev/null) ; then
    echo 'SUCCESS'
    exit 0;
else
    echo 'ERROR'
    exit 1;
fi
