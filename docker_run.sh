#!/usr/bin/env bash

IMAGE_NAME="metricserver"
IMAGE_VERSION="v0.0.3"

echo Runing Docker

docker run -d -p 127.0.0.1:9110:9110/udp --name ${IMAGE_NAME} ${IMAGE_NAME}:${IMAGE_VERSION}

#docker run -d -p 127.0.0.1:9110:9110/udp registry.cn-beijing.aliyuncs.com/zhangga/metricserver:v0.0.3
