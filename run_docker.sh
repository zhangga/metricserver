#!/usr/bin/env bash

IMAGE_NAME="metricserver"

echo Runing Docker

docker run -p 8000:9110 --name ${IMAGE_NAME} ${IMAGE_NAME}