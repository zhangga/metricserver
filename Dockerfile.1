#基础镜像,构建环境
FROM golang as builder

#环境变量
ENV GO111MODULE=on
# 配置模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/zhangga/metricserver
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mreticserver .

FROM alpine as runner
RUN apk --no-cache add ca-certificates

WORKDIR .
#COPY --from=0 $GOPATH/src/github.com/zhangga/metricserver .

EXPOSE 8000
CMD ["./metricserver"]
