package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ErrDeckNotFound is the error returned when a deck is not found.
var ErrDeckNotFound = errors.New("deck not found")

// Repository is a PostgreSQL implementation of the deckr.Repository interface.
type Repository struct {
	log *logger.Logger
	db  *gorm.DB
}

// NewRepository creates a new PostgreSQL repository.
func NewRepository(uri string, log *logger.Logger) (*Repository, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: uri,
	}))
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}

	return &Repository{
		log: log,
		db:  db,
	}, nil
}

// Get returns the deck with the given id.
func (r *Repository) Get(_ context.Context, id string) (*deck.Deck, error) {
	var saved Deck

	if err := r.db.First(&saved, id); err != nil {
		return nil, fmt.Errorf("deck '%s': %w", id, ErrDeckNotFound)
	}

	d, err := DeckToDomainDeck(&saved)
	if err != nil {
		return nil, fmt.Errorf("deck '%s': %w", id, err)
	}

	r.log.Debugf("get deck '%s'", d.ID)

	return d, nil
}

// Save saves the given deck.
func (r *Repository) Save(_ context.Context, d *deck.Deck) error {
	save := DomainDeckToDeck(d)

	r.db.Save(save)

	r.log.Debugf("saved deck '%s'", save.ID)

	return nil
}

// Migrate migratess the deck model.
func (r *Repository) Migrate(_ context.Context) error {
	if err := r.db.AutoMigrate(&Deck{}); err != nil {
		return fmt.Errorf("migrate deck model: %w", err)
	}

	r.log.Debug("migrated deck model")

	return nil
}
