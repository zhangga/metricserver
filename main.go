package main

import (
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"github.com/zhangga/metricserver/server"
	"net/http"
	"time"
)

func main() {
	logs.WithLoggerConf("./conf/logger.yaml")
	conf, err := config.Load("./conf/server.yaml")
	if err != nil {
		logs.Panicf("metric server stop!, err: %v", err)
	}
	// test timezone
	logs.Infof("time now is: %v", time.Now().Format(time.RFC1123Z))
	// test http ssl
	_, err = http.Get("https://api.github.com/")
	if err != nil {
		logs.Errorf("Error! http ssl cert is not found.")
	}

	srv := server.NewServer(conf)
	srv.Start()
}
