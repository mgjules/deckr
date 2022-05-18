package repo

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo/inmemory"
)

// Repository is an interface to get and save a deck.
type Repository interface {
	Get(context.Context, string) (*deck.Deck, error)
	Save(context.Context, *deck.Deck) error
}

// NewRepository returns a new repository.
func NewRepository(uri string, log *logger.Logger) (Repository, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid repository URI: %w", err)
	}

	switch u.Scheme {
	case "inmemory":
		return inmemory.NewRepository(log), nil
	default:
		return nil, fmt.Errorf("unknown repository URI scheme: %s", u.Scheme)
	}
}
