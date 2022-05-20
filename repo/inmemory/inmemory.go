package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/logger"
	"github.com/mgjules/deckr/repo/errs"
)

// Repository is an in-memory implementation of the deckr.Repository interface.
type Repository struct {
	log *logger.Logger

	items map[string]*Deck
	mu    sync.Mutex
}

// NewRepository creates a new in-memory repository.
func NewRepository(log *logger.Logger) *Repository {
	return &Repository{
		log:   log,
		items: make(map[string]*Deck),
	}
}

// Get returns the deck with the given id.
func (r *Repository) Get(_ context.Context, id string) (*deck.Deck, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	saved, ok := r.items[id]
	if !ok {
		return nil, fmt.Errorf("deck '%s': %w", id, errs.ErrDeckNotFound)
	}

	d, err := DeckToDomainDeck(saved)
	if err != nil {
		return nil, fmt.Errorf("deck '%s': %w", id, err)
	}

	r.log.Debugf("get deck '%s'", d.ID)

	return d, nil
}

// Save saves the given deck.
func (r *Repository) Save(_ context.Context, d *deck.Deck) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	save := DomainDeckToDeck(d)

	r.items[save.ID] = save

	r.log.Debugf("saved deck '%s'", save.ID)

	return nil
}

// Migrate migrates the deck model.
func (r *Repository) Migrate(_ context.Context) error {
	r.log.Debug("in-memory repository does not need migration")

	return nil
}

// Close closes any external connection in the repository.
func (r *Repository) Close(_ context.Context) error {
	r.log.Debug("in-memory repository has no external connection to close")

	return nil
}
