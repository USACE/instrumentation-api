docker run \
    -v $(pwd)/tests:/etc/newmann --network=instrumentation-api_default \
    -t postman/newman run /etc/newmann/instrumentation-regression.postman_collection.json \
    --environment=/etc/newmann/postman_environment.docker-compose.json