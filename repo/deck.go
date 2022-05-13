package repo

// Deck represents a deck of cards.
type Deck struct {
	ID       string
	Shuffled bool
	Cards    []Card
}
