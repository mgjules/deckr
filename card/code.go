package card

import (
	"fmt"
)

// Code represents a card code.
type Code struct {
	r string
	s string
}

// CodeFromString returns a new card code given a string.
// The string must be 2 characters long.
// The first character must be a valid rank.
// The second character must be a valid suit.
func CodeFromString(c string) (*Code, error) {
	if len(c) != 2 {
		return nil, fmt.Errorf("card code '%s' should be 2 characters", c)
	}

	return &Code{
		r: c[0:1],
		s: c[1:2],
	}, nil
}

// CodeFromRankSuit returns a new code given a rank and suit.
func CodeFromRankSuit(r Rank, s Suit) (*Code, error) {
	return CodeFromString(r.Code() + s.Code())
}

// String implements the Stringer interface.
func (c Code) String() string {
	return c.r + c.s
}

// Rank returns the rank.
func (c Code) Rank() string {
	return c.r
}

// Suit returns the suit.
func (c Code) Suit() string {
	return c.s
}

// CodesFromStrings returns a collection of codes given a collection of string.
func CodesFromStrings(cc ...string) ([]Code, error) {
	var codes []Code
	for _, c := range cc {
		code, err := CodeFromString(c)
		if err != nil {
			return nil, fmt.Errorf("new code: %w", err)
		}

		codes = append(codes, *code)
	}

	return codes, nil
}
