#!/bin/bash

# mock eventbridge trigger
# use seconds instead of minutes for testing
while true
    do
        curl -o /dev/null -s -w "%{http_code}" -XPOST 'http://alert:8080/2015-03-31/functions/function/invocations' -d '{"name": "mock trigger"}'
        echo -e '\nsleeping 15 seconds...'
        sleep 15
done
