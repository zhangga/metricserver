package server

import (
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"net"
	"strconv"
	"strings"
	"sync"
)

const (
	mtuLimit = 1500
)

type Server struct {
	config	*config.Configure

	udpConn			*net.UDPConn
	influxClient	*InfluxClient

	recvQueue []*MetricData
	metrics	map[string]*MetricData

	die     chan struct{} // notify the server has closed
	dieOnce sync.Once
}

func NewServer(config *config.Configure) *Server {
	srv := &Server{
		config: config,
		influxClient: newInfluxClient(config),
		recvQueue: make([]*MetricData, 1000),
		metrics: make(map[string]*MetricData, 1000),
		die:    make(chan struct{}, 1),
	}
	return srv
}

func (s *Server) Start() {
	udpAddr, err := net.ResolveUDPAddr("udp", s.config.ServerConfig.Address)
	if err != nil {
		logs.Panicf("start server config=%v, err: %v", s.config, err)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		logs.Panicf("start server config=%v, err: %v", s.config, err)
	}
	logs.Infof("server start listened addr: %v", udpConn.LocalAddr().String())
	s.udpConn = udpConn
	go s.udpMonitor()
	select {
	case <-s.die:
		logs.Error("server stop!!!")
	}
}

func (s *Server) Stop() {
	s.dieOnce.Do(func() {
		close(s.die)
	})
}

func (s *Server) udpMonitor() {
	buf := make([]byte, mtuLimit)
	for {
		if n, from, err := s.udpConn.ReadFrom(buf); err == nil && n > 0 {
			s.onUdpRead(buf[:n], from)
		} else {
			logs.Errorf("udp read from: %v, err: %v", from, err)
		}
	}
}

func (s *Server) onUdpRead(data []byte, from net.Addr) {
	body := string(data)
	logs.Debugf("[onUdpRead] from=%v, data=%v", from.String(), body)
	firstIndex := strings.Index(body, "$")
	lastIndex := strings.LastIndex(body, "$")
	name := body[:firstIndex]
	value, err := strconv.ParseFloat(body[lastIndex+1:], 64)
	if err != nil {
		logs.Errorf("[onUdpRead] parse float64 value=%v, err: %v", body[lastIndex+1:], err)
		value = 0
	}
	tags := "{}"
	if (firstIndex < lastIndex) {
		tags = body[firstIndex+1:lastIndex]
	}
	metric := &MetricData{
		Name: name,
		Tags: tags,
		Value: value,
	}

	s.influxClient.WritePoint(metric)
}
