#!/usr/bin/env bash

IMAGE_NAME="metricserver"
IMAGE_VERSION="v0.0.3"
IMAGE_ID="f9422f517dc4"

docker login --username=383523842@qq.com registry.cn-beijing.aliyuncs.com

docker tag ${IMAGE_ID} registry.cn-beijing.aliyuncs.com/zhangga/${IMAGE_NAME}:${IMAGE_VERSION}

docker push registry.cn-beijing.aliyuncs.com/zhangga/${IMAGE_NAME}:${IMAGE_VERSION}