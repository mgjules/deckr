package repo

// Card represents a card.
type Card struct {
	Rank Rank
	Suit Suit
	Code string
}

// Rank represents the rank of a card.
type Rank struct {
	Name string
	Code string
}

// Suit represents the suit of a card.
type Suit struct {
	Name string
	Code string
}
