package grpcx

import (
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	addr       string
	tcpServer  net.Listener
}

func NewServer(server *grpc.Server, addr string) *Server {
	return &Server{
		grpcServer: server,
		addr:       addr,
	}
}

func (s *Server) ListenAndServe() error {
	var err error
	s.tcpServer, err = net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	if err := s.grpcServer.Serve(s.tcpServer); err != nil {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	return s.tcpServer.Close()
}
