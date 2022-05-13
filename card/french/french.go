package french

import "github.com/mgjules/deckr/card"

// Ranks
var (
	ace   = card.NewRank("ACE", "A")
	two   = card.NewRank("2", "2")
	three = card.NewRank("3", "3")
	four  = card.NewRank("4", "4")
	five  = card.NewRank("5", "5")
	six   = card.NewRank("6", "6")
	seven = card.NewRank("7", "7")
	eight = card.NewRank("8", "8")
	nine  = card.NewRank("9", "9")
	ten   = card.NewRank("10", "T")
	jack  = card.NewRank("JACK", "J")
	queen = card.NewRank("QUEEN", "Q")
	king  = card.NewRank("KING", "K")
)

// Suits
var (
	spades   card.Suit = "SPADES"   // ♠
	diamonds card.Suit = "DIAMONDS" // ♦
	clubs    card.Suit = "CLUBS"    // ♣
	hearts   card.Suit = "HEARTS"   // ♥
)

// Composition returns the french deck composition.
// NOTE: Order is important here.
// No joker in this deck.
var Composition = card.NewComposition(
	card.Ranks{ace, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king},
	card.Suits{spades, diamonds, clubs, hearts},
)
