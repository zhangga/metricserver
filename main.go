package main

import (
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/zhangga/logs"
	"time"
)

func main() {
	logs.WithLoggerConf("./conf/logger.yaml")
	logs.Debug("metric server start...")

	org := "jossy"
	bucket := "test"
	url := "http://192.144.167.243:8086/"
	token := "f46BqkFbgCQbXVt5H4p1f-hqp8-v274KMx8ThLaoUB5nb2dmF7Y-Lmc-2EF2tqfTmZsgPUvksQJ1qtOgncS2bw=="
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, bucket)

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	writeAPI.WritePoint(p)
	client.Close()
}
