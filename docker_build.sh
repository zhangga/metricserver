#!/usr/bin/env bash

IMAGE_NAME="metricserver"
IMAGE_VERSION="v0.0.3"

echo Building Dockerfile: ${IMAGE_NAME}:${IMAGE_VERSION}

docker build --no-cache -t ${IMAGE_NAME}:${IMAGE_VERSION} . -f Dockerfile