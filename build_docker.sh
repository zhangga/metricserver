#!/usr/bin/env bash

IMAGE_NAME="metricserver"
IMAGE_VERSION="0.0.2"

echo Building Dockerfile

docker build --no-cache -t ${IMAGE_NAME}:${IMAGE_VERSION} . -f Dockerfile