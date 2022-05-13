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
func NewCard(r Rank, s Suit, c Code) *Card {
	return &Card{
		rank: r,
		suit: s,
		code: c,
	}
}

// Rank returns the rank of the card.
func (c Card) Rank() Rank {
	return c.rank
}

// Suit returns the suit of the card.
func (c Card) Suit() Suit {
	return c.suit
}

// Code returns the code of the card.
func (c Card) Code() Code {
	return c.code
}

// NewCards returns a collection of cards
// using the given ranks, suits and codes.
func NewCards(rs Ranks, ss Suits, codes ...Code) ([]Card, error) {
	var cards []Card

	// Partial cards using codes.
	if len(codes) > 0 {
		for _, c := range codes {
			r, ok := rs.RankFromCode(c)
			if !ok {
				return nil, fmt.Errorf("card code '%s' has an invalid rank", c)
			}

			s, ok := ss.SuitFromCode(c)
			if !ok {
				return nil, fmt.Errorf("card code '%s' has an invalid suit", c)
			}

			cards = append(cards, *NewCard(r, s, c))
		}

		return cards, nil
	}

	// Full cards.
	for _, s := range ss.Suits() {
		for _, r := range rs.Ranks() {
			c, err := NewCodeFromRankSuit(r, s)
			if err != nil {
				return nil, fmt.Errorf("new code: %w", err)
			}

			cards = append(cards, *NewCard(r, s, *c))
		}
	}

	return cards, nil
}
