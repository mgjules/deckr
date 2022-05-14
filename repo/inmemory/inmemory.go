package inmemory

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo"
)

// ErrDeckNotFound is the error returned when a deck is not found.
var ErrDeckNotFound = errors.New("deck not found")

// Repository is an in-memory implementation of the deckr.Repository interface.
type Repository struct {
	log *logger.Logger

	items map[string]*repo.Deck
	mu    sync.Mutex
}

// NewRepository creates a new in-memory repository.
func NewRepository(log *logger.Logger) *Repository {
	return &Repository{
		log:   log,
		items: make(map[string]*repo.Deck),
	}
}

// Get returns the deck with the given id.
func (r *Repository) Get(_ context.Context, id string) (*repo.Deck, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	d, ok := r.items[id]
	if !ok {
		return nil, fmt.Errorf("deck '%s': %w", id, ErrDeckNotFound)
	}

	r.log.Debugf("get deck '%s'", d.ID)

	return d, nil
}

// Save saves the given deck.
func (r *Repository) Save(_ context.Context, d *repo.Deck) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items[d.ID] = d

	r.log.Debugf("saved deck '%s'", d.ID)

	return nil
}
