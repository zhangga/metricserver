FROM alpine:latest
MAINTAINER Jossy Zhang<383523842@qq.com>
ENV VERSION 1.0

WORKDIR /apps
VOLUME ["/apps/conf"]

COPY output/bin/metricserver /apps/metricserver
COPY output/conf/* apps/conf/

ENV LANG C.UTF-8

EXPOSE 8000

ENTRYPOINT ["/apps/metricserver"]

