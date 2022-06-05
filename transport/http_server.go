package transport

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mgjules/deckr/build"
	"github.com/mgjules/deckr/logger"
	v1 "github.com/mgjules/deckr/transport/v1"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	_readTimeout       = 2 * time.Second
	_writeTimeout      = 2 * time.Second
	_idleTimeout       = 30 * time.Second
	_readHeaderTimeout = 2 * time.Second
)

// HTTPServer is the main HTTP server.
type HTTPServer struct {
	router   *gin.Engine
	http     *http.Server
	log      *logger.Logger
	build    *build.Info
	addr     string
	grpcAddr string
}

// NewHTTPServer creates a new Server.
func NewHTTPServer(
	debug bool,
	host string,
	port int,
	logger *logger.Logger,
	build *build.Info,
) *HTTPServer {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	w := logger.Writer()
	gin.DefaultWriter = w
	gin.DefaultErrorWriter = w

	s := HTTPServer{
		router:   gin.Default(),
		addr:     fmt.Sprintf("%s:%d", host, port),
		grpcAddr: fmt.Sprintf("%s:%d", host, port-1),
		log:      logger,
		build:    build,
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

func (s *HTTPServer) registerRoutes() {
	// Health Check
	s.router.GET("/", s.handleHealthCheck())

	// Version
	s.router.GET("/version", s.handleVersion())

	// Swagger
	s.router.StaticFS("/docs", http.FS(docsFS))
	s.router.GET("/swagger/*any", s.handleSwagger())

	// Decks
	mux := runtime.NewServeMux()
	if err := v1.RegisterDeckServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		s.grpcAddr,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	); err != nil {
		s.log.Errorf("register grpc gateway: %v", err)
	}

	s.router.Group("/v1/*{grpc_gateway}").Any("", gin.WrapH(mux))
}

// Start starts the server.
// It blocks until the server stops.
func (s *HTTPServer) Start() error {
	s.log.Infof("Listening on http://%s...", s.addr)

	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

// Stop stops the server.
func (s *HTTPServer) Stop(ctx context.Context) error {
	s.log.Info("Stopping server ...")

	if err := s.http.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}

	return nil
}
