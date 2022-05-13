package repo

// Card represents a card.
type Card struct {
	Rank Rank
	Suit string
	Code string
}

// Rank represents the rank of a card.
type Rank struct {
	Name string
	Code string
}
