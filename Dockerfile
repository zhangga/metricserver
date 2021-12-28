FROM golang:1.17-alpine as builder

# 工作目录
WORKDIR /tmp/github.com/metricserver

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN apk --no-cache add ca-certificates tzdata upx

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
ADD . .
# 构建并使用upx压缩
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o metricserver &&\
    upx --best metricserver -o _upx_metricserver && \
    mv -f _upx_metricserver metricserver

#TODO 把日志volume挂载出去

FROM alpine as runner
# 时区
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 配置文件
COPY --from=builder /tmp/github.com/metricserver/conf ./conf
# 可执行程序
COPY --from=builder /tmp/github.com/metricserver/metricserver /opt/apps/
CMD ["/opt/apps/metricserver"]
