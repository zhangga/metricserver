package server

import (
	"github.com/bytedance/sonic"
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"net"
	"sync"
)

const (
	mtuLimit = 1500
)

type Server struct {
	// 服务器配置
	config *config.Configure

	// UDP连接
	udpConn *net.UDPConn
	// influx client
	influxClient *InfluxClient

	// 接收队列
	recvQueue *RingQueue
	metrics   map[string]*MetricData

	die     chan struct{} // notify the server has closed
	dieOnce sync.Once
}

func NewServer(config *config.Configure) *Server {
	srv := &Server{
		config:       config,
		influxClient: newInfluxClient(config),
		recvQueue:    newRingQueue(config.ServerConfig.RecvSize),
		metrics:      make(map[string]*MetricData, 1000),
		die:          make(chan struct{}, 1),
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
	go s.udpLoop()
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

func (s *Server) udpLoop() {
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
	var metricData MetricData
	err := sonic.Unmarshal(data, &metricData)
	if err != nil {
		logs.Errorf("[onUdpRead] from=%s, data=$s", from.String(), string(data))
		return
	}
	//logs.Debugf("[onUdpRead] from=%v, data=%v", from.String(), metricData)

	s.recvQueue.offer(&metricData)
}

func (s *Server) workLoop() {
	for {
		metricData := s.recvQueue.poll()
		s.influxClient.WritePoint(metricData)
	}
}
