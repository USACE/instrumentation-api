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