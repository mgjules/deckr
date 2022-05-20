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
			Name:    "server-uri",
			Value:   "http://localhost:9000",
			Usage:   "URI of server",
			EnvVars: []string{"DECKR_SERVER_URI"},
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
		defer func() {
			if err = repository.Close(c.Context); err != nil {
				log.Errorf("close repository: %w", err)
			}
		}()

		if err = repository.Migrate(c.Context); err != nil {
			return fmt.Errorf("migrate repository: %w", err)
		}

		transporter, err := transport.NewTransporter(debug, c.String("server-uri"), log, info, repository)
		if err != nil {
			return fmt.Errorf("new transporter: %w", err)
		}
		go func() {
			if err := transporter.Start(); err != nil {
				log.Errorf("start transporter: %w", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(c.Context, 5*time.Second)
		defer cancel()

		if err := transporter.Stop(ctx); err != nil {
			return fmt.Errorf("stop transporter: %w", err)
		}

		return nil
	},
}
