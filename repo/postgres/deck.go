package postgres

import (
	"fmt"

	"github.com/mgjules/deckr/deck"
)

// Deck represents a deck of cards.
type Deck struct {
	ID          string
	Shuffled    bool
	Composition string
	Codes       []string
}

// DomainDeckToDeck transforms a domain deck to a repo deck.
func DomainDeckToDeck(d *deck.Deck) *Deck {
	var rd Deck
	rd.ID = d.ID()
	rd.Shuffled = d.IsShuffled()
	rd.Composition = d.Composition()
	for _, card := range d.Cards() {
		rd.Codes = append(rd.Codes, card.Code().String())
	}

	return &rd
}

// DeckToDomainDeck transforms a repo deck to a domain deck.
func DeckToDomainDeck(rd *Deck) (*deck.Deck, error) {
	d, err := deck.New(
		deck.WithID(rd.ID),
		deck.WithShuffled(rd.Shuffled),
		deck.WithComposition(rd.Composition),
		deck.WithCodes(rd.Codes...),
	)
	if err != nil {
		return nil, fmt.Errorf("new deck: %w", err)
	}

	return d, nil
}
