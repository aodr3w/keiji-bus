package main

import (
	"log"
	"net"

	"github.com/aodr3w/keiji-bus/core"
	"github.com/aodr3w/keiji-core/paths"
	"github.com/aodr3w/keiji-core/utils"
	"github.com/aodr3w/logger"
)

type Server struct {
	address string
	logger  *logger.Logger
}

func NewServer(address string, logger *logger.Logger) *Server {

	return &Server{
		address: address,
		logger:  logger,
	}
}

func (s *Server) start(f func(net.Conn, *logger.Logger)) {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		s.logger.Fatal("Failed to start server: %v", err)
	}
	defer listener.Close()
	log.Printf("Server started at %s", s.address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("failed to accept connection: %v", err)
			continue
		}
		go f(conn, s.logger)

	}
}

func main() {
	mq := core.NewMessageQueue(100)
	lg, err := logger.NewFileLogger(paths.TCP_BUS_LOGS)
	if err != nil {
		log.Fatal(err)
	}
	push := NewServer(
		core.PUSH_PORT,
		lg,
	)

	pull := NewServer(
		core.PULL_PORT,
		lg,
	)

	go push.start(func(c net.Conn, l *logger.Logger) {
		core.HandlePush(mq, c, push.logger)
	})

	go pull.start(func(c net.Conn, l *logger.Logger) {
		core.HandlePull(mq, c, pull.logger)
	})

	utils.HandleStopSignal(func() {
		log.Println("stopping tcp bus")
	})
}
