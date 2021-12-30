package server

import (
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	mtuLimit = 1500
)

type Server struct {
	// 服务器配置
	config			*config.Configure

	// UDP连接
	udpConn			*net.UDPConn
	// influx client
	influxClient	*InfluxClient

	// 接收队列
	recvQueue		*RingQueue
	metrics	map[string]*MetricData

	die     chan struct{} // notify the server has closed
	dieOnce sync.Once
}

func NewServer(config *config.Configure) *Server {
	srv := &Server{
		config: config,
		influxClient: newInfluxClient(config),
		recvQueue: newRingQueue(config.ServerConfig.RecvSize),
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
	go s.workLoop()
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
	//logs.Debugf("[onUdpRead] from=%v, data=%v", from.String(), body)
	firstIndex := strings.Index(body, "$")
	secondIndex := strings.Index(body[firstIndex+1:], "$") + firstIndex + 1
	lastIndex := strings.LastIndex(body, "$")
	mtype, err := strconv.Atoi(body[:firstIndex])
	name := body[firstIndex+1:secondIndex]
	value, err := strconv.ParseFloat(body[lastIndex+1:], 64)
	if err != nil {
		logs.Errorf("[onUdpRead] parse float64 value=%v, err: %v", body[lastIndex+1:], err)
		value = 0
	}
	tags := "{}"
	if (secondIndex < lastIndex) {
		tags = body[secondIndex+1:lastIndex]
	}
	metric := &MetricData{
		Type: mtype,
		Name: name,
		Tags: tags,
		Value: value,
		Time: time.Now().Unix(),
	}
	s.influxClient.WritePoint(metric)
	//s.recvQueue.offer(metric)
}

func (s *Server) workLoop() {
	for {
		metricData := s.recvQueue.poll()
		s.influxClient.WritePoint(metricData)
	}
}
