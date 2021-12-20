package main

import (
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"github.com/zhangga/metricserver/server"
)

func main() {
	logs.WithLoggerConf("./conf/logger.yaml")
	conf, err := config.Load("./conf/server.yaml")
	if err != nil {
		logs.Panicf("metric server stop!, err: %v", err)
	}

	srv := server.NewServer(conf)
	srv.Start()

	//org := "jossy"
	//bucket := "test"
	//url := "http://192.144.167.243:8086/"
	//token := "f46BqkFbgCQbXVt5H4p1f-hqp8-v274KMx8ThLaoUB5nb2dmF7Y-Lmc-2EF2tqfTmZsgPUvksQJ1qtOgncS2bw=="
	//client := influxdb2.NewClient(url, token)
	//writeAPI := client.WriteAPI(org, bucket)
	//
	//p := influxdb2.NewPoint("stat",
	//	map[string]string{"unit": "temperature"},
	//	map[string]interface{}{"avg": 24.5, "max": 45},
	//	time.Now())
	//writeAPI.WritePoint(p)
	//client.Close()
}
