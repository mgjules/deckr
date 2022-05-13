package repo

import (
	"context"
)

// Repository is an interface to get and save a deck.
type Repository interface {
	Get(context.Context, string) (*Deck, error)
	Save(context.Context, *Deck) error
}
