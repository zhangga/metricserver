package server

import (
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	jsoniter "github.com/json-iterator/go"
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"time"
)

type InfluxClient struct {
	client		influxdb2.Client
	writeAPI	api.WriteAPI
}

func newInfluxClient(config *config.Configure) *InfluxClient {
	cli := influxdb2.NewClient(config.DBConfig.Url, config.DBConfig.Token)
	influxCli := &InfluxClient{
		client: cli,
		writeAPI: cli.WriteAPI(config.DBConfig.Org, config.DBConfig.Bucket),
	}
	return influxCli
}

func (cli *InfluxClient) WritePoint(data *MetricData) {
	tags := make(map[string]string)
	err := jsoniter.Unmarshal([]byte(data.Tags), &tags)
	if err != nil {
		logs.Errorf("json unmarshal err: %v", err)
	}
	p := influxdb2.NewPoint(data.Name,
		tags,
		map[string]interface{}{"value": data.Value},
		time.Now())
	cli.writeAPI.WritePoint(p)
}

func (cli *InfluxClient) Close() {
	cli.client.Close()
}
