package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mgjules/deckr/build"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
	"github.com/mgjules/deckr/transport"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

var serve = &cli.Command{
	Name:  "serve",
	Usage: "Starts the REST/gRPC API server.",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Value:   false,
			Usage:   "whether running in PROD or DEBUG mode",
			EnvVars: []string{"DECKR_DEBUG"},
		},
		&cli.StringFlag{
			Name:    "server-host",
			Value:   "localhost",
			Usage:   "HOST of server",
			EnvVars: []string{"DECKR_SERVER_HOST"},
		},
		&cli.IntFlag{
			Name:    "server-port",
			Value:   9000,
			Usage:   "PORT of server",
			EnvVars: []string{"DECKR_SERVER_PORT"},
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

		repository, err := repo.NewRepository(c.Context, storageURI, log)
		if err != nil {
			return fmt.Errorf("new repository: %w", err)
		}
		defer func() {
			if err = repository.Close(c.Context); err != nil {
				log.Errorf("close repository: %w", err)
			}
		}()

		if err = repository.Migrate(c.Context); err != nil {
			return fmt.Errorf("migrate repository: %w", err)
		}

		host := c.String("server-host")
		port := c.Int("server-port")
		grpcServer := transport.NewGRPCServer(host, port, log, repository)
		httpServer := transport.NewHTTPServer(debug, host, port+1, log, info)

		go func() {
			if err = grpcServer.Start(); err != nil {
				log.Errorf("start grpc server: %w", err)
			}
		}()

		go func() {
			if err = httpServer.Start(); err != nil {
				log.Errorf("start http server: %w", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(c.Context, 5*time.Second)
		defer cancel()

		g, ctx := errgroup.WithContext(ctx)

		g.Go(func() error {
			return grpcServer.Stop(ctx)
		})

		g.Go(func() error {
			return httpServer.Stop(ctx)
		})

		if err = g.Wait(); err != nil {
			return fmt.Errorf("wait for servers to stop: %w", err)
		}

		return err
	},
}
