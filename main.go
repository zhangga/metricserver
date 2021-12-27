package main

import (
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"github.com/zhangga/metricserver/server"
)

func main() {
	logs.WithLoggerConf("conf/logger.yaml")
	conf, err := config.Load("conf/server.yaml")
	if err != nil {
		logs.Panicf("metric server stop!, err: %v", err)
	}

	srv := server.NewServer(conf)
	srv.Start()
}
