package french

import "github.com/mgjules/deckr/card"

const (
	spades   card.Suit = "SPADES"   // ♠
	diamonds card.Suit = "DIAMONDS" // ♦
	clubs    card.Suit = "CLUBS"    // ♣
	hearts   card.Suit = "HEARTS"   // ♥
)

// Suits represents the french suits.
// NOTE: Order is important here.
// No joker in this deck.
var Suits = card.Suits{spades, diamonds, clubs, hearts}
