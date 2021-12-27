#!/usr/bin/env bash

CUR_DIR=$(cd $(dirname $0); pwd)

exec "${CUR_DIR}/bin/{{RUN_NAME}}"
