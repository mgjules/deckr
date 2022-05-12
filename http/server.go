package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mgjules/deckr/build"
	"github.com/mgjules/deckr/logger"
)

const (
	_readTimeout       = 2 * time.Second
	_writeTimeout      = 2 * time.Second
	_idleTimeout       = 30 * time.Second
	_readHeaderTimeout = 2 * time.Second
)

// Server is the main HTTP server.
type Server struct {
	router *gin.Engine
	http   *http.Server
	logger *logger.Logger
	build  *build.Info
	addr   string
}

// NewServer creates a new Server.
func NewServer(
	prod bool,
	host string,
	port int,
	logger *logger.Logger,
	build *build.Info,
) *Server {
	if prod {
		gin.SetMode(gin.ReleaseMode)
	}

	w := logger.Writer()
	gin.DefaultWriter = w
	gin.DefaultErrorWriter = w

	s := Server{
		router: gin.Default(),
		addr:   fmt.Sprintf("%s:%d", host, port),
		logger: logger,
		build:  build,
	}

	desugared := logger.Desugar()
	s.router.Use(ginzap.Ginzap(desugared.Logger, time.RFC3339, true))
	s.router.Use(ginzap.RecoveryWithZap(desugared.Logger, true))

	s.http = &http.Server{
		Addr:              s.addr,
		Handler:           s.router,
		ReadTimeout:       _readTimeout,
		WriteTimeout:      _writeTimeout,
		IdleTimeout:       _idleTimeout,
		ReadHeaderTimeout: _readHeaderTimeout,
	}

	s.registerRoutes()

	return &s
}

func (s *Server) registerRoutes() {
	// Health Check
	s.router.GET("/", s.handleHealthCheck())

	// Version
	s.router.GET("/version", s.handleVersion())

	// Swagger
	s.router.GET("/swagger/*any", s.handleSwagger())
}

// Start starts the server.
// It blocks until the server stops.
func (s *Server) Start() error {
	s.logger.Infof("Listening on http://%s...", s.addr)

	if err := s.http.ListenAndServe(); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

// Stop stops the server.
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Info("Stopping server ...")

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
