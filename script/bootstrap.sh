#!/usr/bin/env bash

cd `dirname $0`
CUR_DIR=`pwd`

exec "${CUR_DIR}/bin/{{RUN_NAME}}"
