#!/bin/bash

GREEN='\33[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

if ({
    ((npm list -g postman-to-openapi || npm i postman-to-openapi -g) ||
    (echo "Node must be installed to run this script" && exit 1)) &&

    p2o $(pwd)/tests/instrumentation-regression.postman_collection.json \
        -f $(pwd)/docs/swagger/apidoc.json \
        -o $(pwd)/docs/swagger/compose-env.json;
} &> /dev/null) ; then
    printf "${GREEN}SUCCESS:${NC} OpenAPI doc written to $(pwd)/docs/swagger/apidoc.json\n\n"
else
    printf "${RED}ERROR:${NC} Invalid file location(s). Ensure these files are named and located as follows:"
    printf "\n\tinput file: $(pwd)/tests/instrumentation-regression.postman_collection.json"
    printf "\n\tconfig file: $(pwd)/docs/compose-env.json\n\n"
fi
