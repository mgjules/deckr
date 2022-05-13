package http

// DeckClosed represents a closed deck of cards.
// @Description  represents a closed deck of cards
type DeckClosed struct {
	ID        string `json:"deck_id" example:"f6afe993-9847-508e-b206-2487f1ef5a3c"`
	Shuffled  bool   `json:"shuffled" example:"true"`
	Remaining int    `json:"remaining" example:"1"`
}

// DeckOpened represents a opened deck of cards.
// @Description  represents a opened deck of cards
type DeckOpened struct {
	DeckClosed
	Cards []Card `json:"cards"`
}
