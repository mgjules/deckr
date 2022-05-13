package card

import "strings"

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
func (rs Ranks) RankFromCode(c Code) (*Rank, bool) {
	for _, r := range rs.Ranks() {
		if strings.EqualFold(r.Code(), c.Rank()) {
			return &r, true
		}
	}

	return nil, false
}
