package french

import "github.com/mgjules/deckr/card"

const (
	ace   card.Rank = "ACE"
	two   card.Rank = "2"
	three card.Rank = "3"
	four  card.Rank = "4"
	five  card.Rank = "5"
	six   card.Rank = "6"
	seven card.Rank = "7"
	eight card.Rank = "8"
	nine  card.Rank = "9"
	ten   card.Rank = "10"
	jack  card.Rank = "JACK"
	queen card.Rank = "QUEEN"
	king  card.Rank = "KING"
)

// Ranks represents the french ranks.
// NOTE: Order is important here.
// No joker in this deck.
var Ranks = card.Ranks{ace, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king}
