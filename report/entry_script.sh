#!/bin/sh

if [ -z "${AWS_LAMBDA_RUNTIME_API}" ]; then
  exec ~/.aws-lambda-rie/aws-lambda-rie npx aws-lambda-ric $1
else
  exec npx aws-lambda-ric $1
fi
