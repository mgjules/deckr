package card

import (
	"fmt"
)

// Card represents a card in a deck.
type Card struct {
	rank Rank
	suit Suit
	code Code
}

// NewCard returns a new card given a rank and suit.
func NewCard(r Rank, s Suit) (*Card, error) {
	c, err := NewCodeFromRankSuit(r, s)
	if err != nil {
		return nil, fmt.Errorf("new code: %w", err)
	}

	return &Card{
		rank: r,
		suit: s,
		code: *c,
	}, nil
}

// Rank returns the rank of the card.
func (c *Card) Rank() Rank {
	return c.rank
}

// Suit returns the suit of the card.
func (c *Card) Suit() Suit {
	return c.suit
}

// Code returns the code of the card.
func (c *Card) Code() Code {
	return c.code
}
