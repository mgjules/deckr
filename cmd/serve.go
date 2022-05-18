package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mgjules/deckr/build"
	"github.com/mgjules/deckr/grpc"
	"github.com/mgjules/deckr/http"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
	"github.com/urfave/cli/v2"
)

var serve = &cli.Command{
	Name:  "serve",
	Usage: "Starts the REST API server.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Value:   false,
			Usage:   "whether running in PROD or DEBUG mode",
			EnvVars: []string{"DECKR_DEBUG"},
		},
		&cli.BoolFlag{
			Name:    "http",
			Value:   true,
			Usage:   "whether to start the HTTP server",
			EnvVars: []string{"DECKR_HTTP"},
		},
		&cli.StringFlag{
			Name:    "http-host",
			Value:   "localhost",
			Usage:   "host/IP for HTTP server",
			EnvVars: []string{"DECKR_HTTP_HOST"},
		},
		&cli.IntFlag{
			Name:    "http-port",
			Value:   9000,
			Usage:   "port for HTTP server",
			EnvVars: []string{"DECKR_HTTP_PORT"},
		},
		&cli.BoolFlag{
			Name:    "grpc",
			Value:   false,
			Usage:   "whether to start the GRPC server",
			EnvVars: []string{"DECKR_GRPC"},
		},
		&cli.StringFlag{
			Name:    "grpc-host",
			Value:   "localhost",
			Usage:   "host/IP for GRPC server",
			EnvVars: []string{"DECKR_GRPC_HOST"},
		},
		&cli.IntFlag{
			Name:    "grpc-port",
			Value:   9001,
			Usage:   "port for GRPC server",
			EnvVars: []string{"DECKR_GRPC_PORT"},
		},
		&cli.StringFlag{
			Name:    "storage-uri",
			Value:   "inmemory://",
			Usage:   "URI of storage",
			EnvVars: []string{"DECKR_STORAGE_URI"},
		},
	},
	Action: func(c *cli.Context) error {
		debug := c.Bool("debug")

		log, err := logger.New(debug)
		if err != nil {
			return fmt.Errorf("new logger: %w", err)
		}

		info, err := build.New()
		if err != nil {
			return fmt.Errorf("new build info: %w", err)
		}

		storageURI := c.String("storage-uri")

		repository, err := repo.NewRepository(storageURI, log)
		if err != nil {
			return fmt.Errorf("new repository: %w", err)
		}

		var httpServer *http.Server
		httpEnabled := c.Bool("http")
		if httpEnabled {
			httpHost := c.String("http-host")
			httpPort := c.Int("http-port")

			httpServer = http.NewServer(debug, httpHost, httpPort, log, info, repository)
			go func() {
				if err := httpServer.Start(); err != nil {
					log.Errorf("start http server: %v", err)
				}
			}()
		}

		var grpcServer *grpc.Server
		grpcEnabled := c.Bool("grpc")
		if grpcEnabled {
			grpcHost := c.String("grpc-host")
			grpcPort := c.Int("grpc-port")

			grpcServer = grpc.NewServer(grpcHost, grpcPort, log, repository)
			go func() {
				if err := grpcServer.Start(); err != nil {
					log.Errorf("start grpc server: %v", err)
				}
			}()
		}

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(c.Context, 5*time.Second)
		defer cancel()

		if httpEnabled {
			if err := httpServer.Stop(ctx); err != nil {
				return fmt.Errorf("stop http server: %w", err)
			}
		}

		if grpcEnabled {
			grpcServer.Stop()
		}

		return nil
	},
}
