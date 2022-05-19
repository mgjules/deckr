package transport

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/deckr/build"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
	"github.com/mgjules/deckr/transport/grpc"
	"github.com/mgjules/deckr/transport/http"
)

// Transporter is an interface to perform deck operations.
type Transporter interface {
	Start() error
	Stop(context.Context) error
}

// NewTransporter returns a new transporter.
func NewTransporter(
	debug bool,
	uri string,
	log *logger.Logger,
	build *build.Info,
	repo repo.Repository,
) (Transporter, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid transporter URI: %w", err)
	}

	switch u.Scheme {
	case "http":
		return http.NewServer(debug, u.Host, log, build, repo), nil
	case "tcp":
		return grpc.NewServer(u.Host, log, repo), nil
	default:
		return nil, fmt.Errorf("unknown transporter URI scheme: %s", u.Scheme)
	}
}
