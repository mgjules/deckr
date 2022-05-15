package composition

import (
	"errors"
	"strings"

	"github.com/mgjules/deckr/card"
)

var (
	// ErrUnknownComposition is returned when a composition is not found.
	ErrUnknownComposition = errors.New("unknown composition")

	// compositions is a dictionary of card compositions.
	compositions = make(map[string]*Composition)
)

// FromString returns a card composition given a string.
func FromString(s string) (*Composition, error) {
	c, ok := compositions[strings.ToLower(s)]
	if !ok {
		return nil, ErrUnknownComposition
	}

	return c, nil
}

// Composition represents a collection of ranks and suits. (e.g french)
type Composition struct {
	ranks card.Ranks
	suits card.Suits
}

// New returns a new card composition given a collection of ranks and suits.
func New(ranks card.Ranks, suits card.Suits) *Composition {
	return &Composition{
		ranks: ranks,
		suits: suits,
	}
}

// Ranks returns a collection of ranks.
func (c Composition) Ranks() card.Ranks {
	return c.ranks
}

// Suits returns a collection of suits.
func (c Composition) Suits() card.Suits {
	return c.suits
}
