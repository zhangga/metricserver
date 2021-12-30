# Metrics 使用手册
Metrics提供一套完整的监控服务解决方案，它是基于时序性数据库 [InfluxDB](https://docs.influxdata.com/influxdb/v2.1/) 作为数据存储，使用 [Grafana](https://grafana.com/grafana/) 作为数据展现，方便易用。数据接入的SDK在项目 [MetricSDK](https://github.com/zhangga/metricsdk) 中。

## 环境部署
首先你需要找一台服务器来部署InfluxDB和Grafana来提供数据的存储和展现服务。下面提供了几种服务部署的方案，你可以根据实际生产环境来选择。

* InfluxDB
  1. 【推荐】云主机安装方式，可以参考我的[文章](http://kwaibook.com/influxdb-grafana/)。这样可以将InfluxDB的机器部署在其他服务器相同的内网环境。
  2. [阿里云产品](https://www.aliyun.com/product/hitsdb_influxdb_pre)等。

* Grafana
  1. 【推荐】使用Grafana官网提供的服务，可以参考我的[文章](http://kwaibook.com/influxdb-grafana/)。
  2. 云主机安装方式，可以参考我的[文章](http://kwaibook.com/influxdb-grafana/)。

## 服务部署

部署好服务所需的InfluxDB和Grafana后，监控服务本身只需要在每台服务器起一个进程即可。

有两种方式可选，这里推荐Docker方式启动运行。

* Docker部署【推荐】

  1. 服务器安装Docker，可以参考我的[文章](http://kwaibook.com/influxdb-grafana/)。

  2. 运行以下命令，拉去最新镜像启动

     ```
     docker run -d -p 127.0.0.1:9110:9110/udp registry.cn-beijing.aliyuncs.com/zhangga/metricserver:v0.0.3
     ```

* 可执行程序部署

  1. 安装服务运行需要的golang1.17环境。
  2. 执行项目根目录下`build.sh`脚本打包最终产物，在`output`目录下，执行`bootstrap.sh`脚本启动服务。

## 服务配置
服务配置文件在根目录的`conf`文件夹下

1. 修改服务端口。

   文件：`server.yaml`，配置：`Server:Address`

   可以修改服务器端口，如果以docker方式运行的话，没必要修改这里的配置。可以启动容器的时候做端口映射，比如：`docker run IP:端口:9110/udp`

2. 修改InfluxDB配置。

   文件：`server.yaml`，配置：`DB`

   修改成你数据库的Url、Org、Bucket、Token。这些信息在安装InfluxDB的时候配置过。

   [Create a token](https://docs.influxdata.com/influxdb/v2.1/security/tokens/create-token/)

## Dockerfile文档
扩展部分，想修改dockerfile的可以参考下这个文档：

https://juejin.cn/post/6844904174396637197