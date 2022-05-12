package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mgjules/deckr/build"
	"github.com/mgjules/deckr/http"
	"github.com/mgjules/deckr/logger"
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
		&cli.StringFlag{
			Name:    "host",
			Value:   "localhost",
			Usage:   "host/IP for HTTP server",
			EnvVars: []string{"DECKR_HOST"},
		},
		&cli.IntFlag{
			Name:    "port",
			Value:   9000,
			Usage:   "port for HTTP server",
			EnvVars: []string{"DECKR_PORT"},
		},
	},
	Action: func(c *cli.Context) error {
		debug := c.Bool("debug")

		log, err := logger.New(debug)
		if err != nil {
			return fmt.Errorf("new logger: %w", err)
		}

		host := c.String("host")
		port := c.Int("port")

		info, err := build.New()
		if err != nil {
			return fmt.Errorf("new build info: %w", err)
		}

		server := http.NewServer(debug, host, port, log, info)
		go func() {
			if err := server.Start(); err != nil {
				log.Errorf("start server: %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			return fmt.Errorf("stop server: %w", err)
		}

		return nil
	},
}
