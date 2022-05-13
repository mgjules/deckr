package repo

// Deck represents a deck of cards.
type Deck struct {
	ID       string `json:"id"`
	Shuffled bool   `json:"shuffled"`
	Cards    []Card `json:"cards"`
}
