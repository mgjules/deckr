package transport

import (
	context "context"
	"fmt"
	"net"

	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
	v1 "github.com/mgjules/deckr/transport/v1"
	grpc "google.golang.org/grpc"
)

// GRPCServer is a grpc server.
type GRPCServer struct {
	addr   string
	server *grpc.Server
	log    *logger.Logger
	repo   repo.Repository
}

// NewGRPCServer creates a new grpc server.
func NewGRPCServer(
	host string,
	port int,
	log *logger.Logger,
	repo repo.Repository,
) *GRPCServer {
	s := &GRPCServer{
		addr:   fmt.Sprintf("%s:%d", host, port),
		server: grpc.NewServer(),
		log:    log,
		repo:   repo,
	}

	s.registerServices()

	return s
}

// registerServices registers services with the grpc server.
func (s *GRPCServer) registerServices() {
	v1.RegisterDeckServiceServer(s.server, NewDeckService(s.log, s.repo))
}

// Start starts the grpc server.
func (s *GRPCServer) Start() error {
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
func (s *GRPCServer) Stop(context.Context) error {
	s.log.Info("Stopping server ...")

	s.server.GracefulStop()

	return nil
}
