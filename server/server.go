package server

import (
	"github.com/zhangga/logs"
	"github.com/zhangga/metricserver/config"
	"net"
	"sync"
)

type Server struct {
	config *config.Configure

	die     chan struct{} // notify the server has closed
	dieOnce sync.Once
}

func NewServer(config *config.Configure) *Server {
	srv := &Server{
		config: config,
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
