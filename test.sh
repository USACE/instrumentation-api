docker run \
    --rm \
    -v $(pwd)/tests:/etc/newman --network=instrumentation-api_default \
    -t postman/newman run /etc/newman/instrumentation-regression.postman_collection.json \
    --environment=/etc/newman/postman_environment.docker-compose.json