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
	"github.com/mgjules/deckr/repo"
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
	log    *logger.Logger
	build  *build.Info
	repo   repo.Repository
	addr   string
}

// NewServer creates a new Server.
func NewServer(
	debug bool,
	host string,
	port int,
	logger *logger.Logger,
	build *build.Info,
	repo repo.Repository,
) *Server {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	w := logger.Writer()
	gin.DefaultWriter = w
	gin.DefaultErrorWriter = w

	s := Server{
		router: gin.Default(),
		addr:   fmt.Sprintf("%s:%d", host, port),
		log:    logger,
		build:  build,
		repo:   repo,
	}

	desugared := logger.Desugar()
	s.router.Use(ginzap.Ginzap(desugared, time.RFC3339, true))
	s.router.Use(ginzap.RecoveryWithZap(desugared, true))

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

	deck := s.router.Group("/decks")
	{
		deck.POST("", s.handleCreateDeck())
		deck.GET("/:id", s.handleOpenDeck())
		deck.GET("/:id/draw", s.handleDrawCards())
		deck.POST("/:id/shuffle", s.handleShuffleDeck())
	}
}

// Start starts the server.
// It blocks until the server stops.
func (s *Server) Start() error {
	s.log.Infof("Listening on http://%s...", s.addr)

	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

// Stop stops the server.
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("Stopping server ...")

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
