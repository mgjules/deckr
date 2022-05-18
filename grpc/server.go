package grpc

import (
	"fmt"
	"net"

	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
	grpc "google.golang.org/grpc"
)

// Server is a grpc server.
type Server struct {
	addr   string
	server *grpc.Server
	log    *logger.Logger
	repo   repo.Repository
}

// NewServer creates a new grpc server.
func NewServer(
	host string,
	port int,
	log *logger.Logger,
	repo repo.Repository,
) *Server {
	s := &Server{
		addr:   fmt.Sprintf("%s:%d", host, port),
		server: grpc.NewServer(),
		log:    log,
		repo:   repo,
	}

	return s
}

// RegisterServices registers services with the grpc server.
func (s *Server) RegisterServices() {
	RegisterDeckServiceServer(s.server, NewDeckService(s.log, s.repo))
}

// Start starts the grpc server.
func (s *Server) Start() error {
	s.log.Infof("Listening on tcp://%s...", s.addr)

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

// Stop stops the grpc server.
func (s *Server) Stop() {
	s.log.Info("Stopping server ...")

	s.server.GracefulStop()
}
