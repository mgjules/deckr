package card

import (
	"errors"
	"strings"
)

// ErrInvalidRank is returned when a rank is invalid.
var ErrInvalidRank = errors.New("invalid rank")

// Rank represents the rank of a card.
type Rank struct {
	name string
	code string
}

// NewRank returns a new rank given a name and code.
func NewRank(name, code string) Rank {
	return Rank{
		name: name,
		code: code,
	}
}

// Code returns the code of the rank.
func (r Rank) Code() string {
	return r.code
}

// String implements the Stringer interface.
func (r Rank) String() string {
	return r.name
}

// Ranks is a collection of Rank.
type Ranks []Rank

// Ranks returns the french ranks.
func (rs Ranks) Ranks() []Rank {
	return rs
}

// RankFromCode returns the rank from the given code.
func (rs Ranks) RankFromCode(c Code) (*Rank, error) {
	for _, r := range rs.Ranks() {
		if strings.EqualFold(r.Code(), c.Rank()) {
			return &r, nil
		}
	}

	return nil, ErrInvalidRank
}
