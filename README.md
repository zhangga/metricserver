# metricserver
Metric Server

## 服务器环境部署
首先根据 [文档](http://kwaibook.com/influxdb-grafana/) 在服务器上部署好所需环境：
1. Docker
2. InfluxDB
3. InfluxCLI
4. Grafana

## 创建InfluxDB Token
[Create a token](https://docs.influxdata.com/influxdb/v2.1/security/tokens/create-token/)

启动UDP端口接收数据，每30s处理后将数据存储在influxDB中。

### dockerfile文档
https://juejin.cn/post/6844904174396637197