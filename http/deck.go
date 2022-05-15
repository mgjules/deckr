package http

import (
	"fmt"

	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	"github.com/mgjules/deckr/deck"
	"github.com/mgjules/deckr/repo"
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

// DomainDeckToRepoDeck transforms a domain deck to a repo deck.
func DomainDeckToRepoDeck(d *deck.Deck) *repo.Deck {
	var rd repo.Deck
	rd.ID = d.ID()
	rd.Shuffled = d.IsShuffled()
	rd.Composition = d.Composition()
	for _, card := range d.Cards() {
		rd.Cards = append(rd.Cards, card.Code().String())
	}

	return &rd
}

// RepoDeckToDeckOpened transforms a repo deck to a DeckOpened.
func RepoDeckToDeckOpened(rd *repo.Deck) (*DeckOpened, error) {
	var d DeckOpened
	d.ID = rd.ID
	d.Shuffled = rd.Shuffled
	d.Remaining = len(rd.Cards)

	comp, err := composition.ParseFromString(rd.Composition)
	if err != nil {
		return nil, fmt.Errorf("parse composition: %w", err)
	}

	for _, rc := range rd.Cards {
		c, err := card.NewCode(rc)
		if err != nil {
			return nil, fmt.Errorf("new code: %w", err)
		}

		r, ok := comp.Ranks().RankFromCode(*c)
		if !ok {
			return nil, fmt.Errorf("card code '%s' has an invalid rank", c)
		}

		s, ok := comp.Suits().SuitFromCode(*c)
		if !ok {
			return nil, fmt.Errorf("card code '%s' has an invalid suit", c)
		}

		d.Cards = append(d.Cards, Card{
			Value: r.String(),
			Suit:  s.String(),
			Code:  c.String(),
		})
	}

	return &d, nil
}

// RepoDeckToDomainDeck transforms a repo deck to a domain deck.
func RepoDeckToDomainDeck(rd *repo.Deck) (*deck.Deck, error) {
	var codes []card.Code
	for _, rc := range rd.Cards {
		c, err := card.NewCode(rc)
		if err != nil {
			return nil, fmt.Errorf("new code: %w", err)
		}

		codes = append(codes, *c)
	}

	d, err := deck.New(
		deck.WithID(rd.ID),
		deck.WithShuffled(rd.Shuffled),
		deck.WithComposition(rd.Composition),
		deck.WithCodes(codes...),
	)
	if err != nil {
		return nil, fmt.Errorf("new deck: %w", err)
	}

	return d, nil
}
