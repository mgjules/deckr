package card

import "strings"

// Suit represents the suit of a card.
type Suit string

// String implements the Stringer interface.
func (s Suit) String() string {
	return string(s)
}

// Suits is a collection of Suit.
type Suits []Suit

// Suits returns the french suits.
func (ss Suits) Suits() []Suit {
	return ss
}

// SuitFromCode returns the suit from the given code.
func (ss Suits) SuitFromCode(c Code) (Suit, bool) {
	for _, s := range ss.Suits() {
		if strings.EqualFold(s.String()[0:1], c.Suit()) {
			return s, true
		}
	}

	return "", false
}
