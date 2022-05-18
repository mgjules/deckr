package http

import (
	"github.com/mgjules/deckr/deck"
)

// DeckClosed represents a closed deck of cards.
// @Description  represents a closed deck of cards
type DeckClosed struct {
	ID        string `json:"deck_id" example:"f6afe993-9847-508e-b206-2487f1ef5a3c"`
	Shuffled  bool   `json:"shuffled" example:"true"`
	Remaining int    `json:"remaining" example:"1"`
}

// DeckOpened represents a opened deck of cards.
// @Description  represents a opened deck of cards
type DeckOpened struct {
	DeckClosed
	Cards []Card `json:"cards"`
}

// DomainDeckToDeckClosed transforms a domain deck to a DeckClosed.
func DomainDeckToDeckClosed(d *deck.Deck) *DeckClosed {
	var dc DeckClosed
	dc.ID = d.ID()
	dc.Shuffled = d.IsShuffled()
	dc.Remaining = d.Remaining()

	return &dc
}

// DomainDeckToDeckOpened transforms a domain deck to a DeckOpened.
func DomainDeckToDeckOpened(d *deck.Deck) *DeckOpened {
	var do DeckOpened
	do.DeckClosed = *DomainDeckToDeckClosed(d)
	do.Cards = DomainCardsToCards(d.Cards()).Cards

	return &do
}
