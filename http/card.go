package http

// Card represents a card.
// @Description  represents a card
type Card struct {
	Value string `json:"value" example:"ACE"`
	Suit  string `json:"suit" example:"SPADES"`
	Code  string `json:"code" example:"AS"`
}

// Cards represents a collection of cards.
// @Description  represents a collection of cards
type Cards struct {
	Cards []Card `json:"cards"`
}
