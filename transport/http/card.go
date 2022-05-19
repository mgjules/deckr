package http

import "github.com/mgjules/deckr/card"

// Card represents a card.
// @Description  represents a card
type Card struct {
	Value string `json:"value" example:"ACE"`
	Suit  string `json:"suit" example:"SPADES"`
	Code  string `json:"code" example:"AS"`
}

// Cards represents a collection of cards.
// @Description  represents a collection of cards
type Cards struct {
	Cards []Card `json:"cards"`
}

// DomainCardsToCards transforms domain cards to Cards.
func DomainCardsToCards(dc []card.Card) *Cards {
	var c Cards
	for _, card := range dc {
		c.Cards = append(c.Cards, Card{
			Value: card.Rank().String(),
			Suit:  card.Suit().String(),
			Code:  card.Code().String(),
		})
	}

	return &c
}
