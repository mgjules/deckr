package http

import (
	"fmt"

	"github.com/mgjules/deckr/card"
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

// DomainDeckToRepo transforms a domain deck to a repo deck.
func DomainDeckToRepo(d *deck.Deck) *repo.Deck {
	var rd repo.Deck
	rd.ID = d.ID()
	rd.Shuffled = d.IsShuffled()
	for _, card := range d.Cards() {
		rd.Cards = append(rd.Cards, repo.Card{
			Rank: repo.Rank{Name: card.Rank().String(), Code: card.Rank().Code()},
			Suit: repo.Suit{Name: card.Suit().String(), Code: card.Suit().Code()},
			Code: card.Code().String(),
		})
	}

	return &rd
}

// RepoDeckToDeckOpened transforms a repo deck to a DeckOpened.
func RepoDeckToDeckOpened(rd *repo.Deck) *DeckOpened {
	var d DeckOpened
	d.ID = rd.ID
	d.Shuffled = rd.Shuffled
	d.Remaining = len(rd.Cards)
	for _, card := range rd.Cards {
		d.Cards = append(d.Cards, Card{
			Value: card.Rank.Name,
			Suit:  card.Suit.Name,
			Code:  card.Code,
		})
	}

	return &d
}

// RepoDeckToDomainDeck transforms a repo deck to a domain deck.
func RepoDeckToDomainDeck(rd *repo.Deck) (*deck.Deck, error) {
	var cc []card.Card
	for _, rc := range rd.Cards {
		rank := card.NewRank(rc.Rank.Name, rc.Rank.Code)
		suit := card.NewSuit(rc.Suit.Name, rc.Suit.Code)

		var code *card.Code
		code, err := card.NewCode(rc.Code)
		if err != nil {
			continue
		}

		cc = append(cc, *card.NewCard(rank, suit, *code))
	}

	d, err := deck.New(
		deck.WithID(rd.ID),
		deck.WithShuffled(rd.Shuffled),
		deck.WithCards(cc...),
	)
	if err != nil {
		return nil, fmt.Errorf("new deck: %w", err)
	}

	return d, nil
}
