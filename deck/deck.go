package deck

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/mgjules/deckr/card"
)

// Deck represents a deck of cards.
type Deck struct {
	id       string
	shuffled bool
	cards    []card.Card
}

// New creates a new deck with the given number of cards.
func New(cards ...card.Card) (*Deck, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("new deck id: %w", err)
	}

	return &Deck{
		id:       id.String(),
		shuffled: false,
		cards:    cards,
	}, nil
}

// Cards returns the cards in the deck.
func (d Deck) Cards() []card.Card {
	return d.cards
}

// Remaining returns the number of cards remaining in the deck.
func (d Deck) Remaining() int {
	return len(d.cards)
}

// IsShuffled returns true if the deck is shuffled.
func (d Deck) IsShuffled() bool {
	return d.shuffled
}

// Shuffle shuffles the cards randomly in the deck.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	d.shuffled = true
}

// Draw returns n number of cards from the deck.
func (d *Deck) Draw(n int) ([]card.Card, error) {
	if n > d.Remaining() {
		return nil, fmt.Errorf("cannot draw %d cards as only %d remaining", n, d.Remaining())
	}

	var cards []card.Card
	for i := 0; i < n; i++ {
		var c card.Card

		// Pop card from deck (considered a stack).
		c, d.cards = d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]

		cards = append(cards, c)
	}

	return cards, nil
}
