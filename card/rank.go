package card

import "strings"

// Rank represents the rank of a card.
type Rank string

// String implements the Stringer interface.
func (r Rank) String() string {
	return string(r)
}

// Ranks is a collection of Rank.
type Ranks []Rank

// Ranks returns the french ranks.
func (rs Ranks) Ranks() []Rank {
	return rs
}

// GetRankFromCode returns the rank from the given code.
func (rs Ranks) GetRankFromCode(c Code) (Rank, bool) {
	for _, r := range rs.Ranks() {
		if strings.EqualFold(r.String()[0:1], c.Rank()) {
			return r, true
		}
	}

	return "", false
}
