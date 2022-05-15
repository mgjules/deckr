package card

import (
	"errors"
	"strings"
)

// ErrInvalidSuit is returned when a suit is invalid.
var ErrInvalidSuit = errors.New("invalid suit")

// Suit represents the suit of a card.
type Suit struct {
	name string
	code string
}

// NewSuit returns a new suit given a name and code.
func NewSuit(name, code string) Suit {
	return Suit{
		name: name,
		code: code,
	}
}

// Code returns the code of the suit.
func (s Suit) Code() string {
	return s.code
}

// String implements the Stringer interface.
func (s Suit) String() string {
	return s.name
}

// Suits is a collection of Suit.
type Suits []Suit

// Suits returns the french suits.
func (ss Suits) Suits() []Suit {
	return ss
}

// SuitFromCode returns the suit from the given code.
func (ss Suits) SuitFromCode(c Code) (*Suit, error) {
	for _, s := range ss.Suits() {
		if strings.EqualFold(s.Code(), c.Suit()) {
			return &s, nil
		}
	}

	return nil, ErrInvalidSuit
}
