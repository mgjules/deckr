package deck

import "github.com/mgjules/deckr/card"

// Option is a function that can be used to modify a deck.
type Option func(d *Deck)

// WithID returns an option that sets the id of the deck.
func WithID(id string) Option {
	return func(d *Deck) {
		d.id = id
	}
}

// WithShuffled returns an option that sets the shuffled state of the deck.
func WithShuffled(s bool) Option {
	return func(d *Deck) {
		d.shuffled = s
	}
}

// WithCards returns an option that sets the cards of the deck.
func WithCards(cards ...card.Card) Option {
	return func(d *Deck) {
		d.cards = cards
	}
}
