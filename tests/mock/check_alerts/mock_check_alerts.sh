#!/bin/bash

# mock eventbridge trigger
# use seconds instead of minutes for testing
while true
    do
        curl -o /dev/null -s -w "%{http_code}" -XPOST "http://api:80/check_alerts" -d '{}'
        echo -e '\nsleeping 15 seconds...'
        sleep 15
done
