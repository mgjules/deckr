package deck

import (
	"errors"
	"math/rand"
	"time"

	"github.com/mgjules/deckr/card"
	uuid "github.com/satori/go.uuid"
)

// ErrNotEnoughCards is the error returned when there are not enough cards in the deck.
var ErrNotEnoughCards = errors.New("not enough cards")

// Deck represents a deck of cards.
type Deck struct {
	id       string
	shuffled bool
	cards    []card.Card
}

// New creates a new deck with the given options.
func New(opts ...Option) (*Deck, error) {
	var d Deck

	for _, applyOpt := range opts {
		applyOpt(&d)
	}

	if d.id == "" {
		d.id = uuid.NewV4().String()
	}

	return &d, nil
}

// ID returns the id of the deck.
func (d Deck) ID() string {
	return d.id
}

// Cards returns the cards in the deck.
func (d Deck) Cards() []card.Card {
	return d.cards
}

// Remaining returns the cards remaining in the deck.
func (d Deck) Remaining() int {
	return len(d.cards)
}

// IsShuffled returns true if the deck is shuffled.
func (d Deck) IsShuffled() bool {
	return d.shuffled
}

// Shuffle shuffles the cards randomly in the deck.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	d.shuffled = true
}

// Draw returns n number of cards from the deck.
func (d *Deck) Draw(n int) ([]card.Card, error) {
	if n > d.Remaining() {
		return nil, ErrNotEnoughCards
	}

	var cards []card.Card
	for i := 0; i < n; i++ {
		var c card.Card

		// Dealing from the top (KH, QH, JH, TH, 9H, ...).
		c, d.cards = d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]

		cards = append(cards, c)
	}

	return cards, nil
}
