package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	pb "gogrpc/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server is a gRPC server.
type Server struct {
	pb.DemoServiceServer

	grpcServer *grpc.Server
	listener   net.Listener
}

// New creates a new gRPC server.
func New(listen string, opt ...grpc.ServerOption) (*Server, error) {
	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	time.Sleep(time.Millisecond * 100) // give the listener time to start

	grpcServer := grpc.NewServer(opt...)
	pb.RegisterDemoServiceServer(grpcServer, &Server{})

	return &Server{
		grpcServer: grpcServer,
		listener:   listener,
	}, nil
}

// NewTLS creates a new gRPC server with TLS.
func NewTLS(listen string, cert tls.Certificate) (*Server, error) {
	cred := credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{cert}})

	return New(listen, grpc.Creds(cred))
}

// Start starts the gRPC server.
func (s *Server) Start() error {
	if err := s.grpcServer.Serve(s.listener); err != nil {
		return fmt.Errorf("failed to serve gRPC: %v", err)
	}

	return nil
}

// Close closes the gRPC server.
func (s *Server) Close() {
	if s.listener != nil {
		s.listener.Close()
	}
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
}
