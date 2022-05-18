package repo

import (
	"context"

	"github.com/mgjules/deckr/deck"
)

// Repository is an interface to get and save a deck.
type Repository interface {
	Get(context.Context, string) (*deck.Deck, error)
	Save(context.Context, *deck.Deck) error
}
