package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/mgjules/deckr/card"
	"github.com/mgjules/deckr/composition"
	"github.com/satori/uuid"
)

// ErrNotEnoughCards is the error returned when there are not enough cards in the deck.
var ErrNotEnoughCards = errors.New("not enough cards")

// Deck represents a deck of cards.
type Deck struct {
	id          string
	shuffled    bool
	codes       []card.Code
	composition string
	cards       []card.Card
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

	if d.composition == "" {
		d.composition = composition.French
	}

	comp, err := composition.ParseFromString(d.composition)
	if err != nil {
		return nil, fmt.Errorf("parse composition: %w", err)
	}

	// Partial cards using codes.
	if len(d.codes) > 0 {
		for _, c := range d.codes {
			r, ok := comp.Ranks().RankFromCode(c)
			if !ok {
				return nil, fmt.Errorf("card code '%s' has an invalid rank", c)
			}

			s, ok := comp.Suits().SuitFromCode(c)
			if !ok {
				return nil, fmt.Errorf("card code '%s' has an invalid suit", c)
			}

			card, err := card.NewCard(*r, *s)
			if err != nil {
				return nil, fmt.Errorf("new card: %w", err)
			}

			d.cards = append(d.cards, *card)
		}

		return &d, nil
	}

	// Full cards.
	for _, s := range comp.Suits() {
		for _, r := range comp.Ranks() {
			card, err := card.NewCard(r, s)
			if err != nil {
				return nil, fmt.Errorf("new card: %w", err)
			}

			d.cards = append(d.cards, *card)
		}
	}

	return &d, nil
}

// ID returns the id of the deck.
func (d *Deck) ID() string {
	return d.id
}

// Cards returns the cards in the deck.
func (d *Deck) Cards() []card.Card {
	return d.cards
}

// Remaining returns the cards remaining in the deck.
func (d *Deck) Remaining() int {
	return len(d.cards)
}

// Composition returns the composition of the deck.
func (d *Deck) Composition() string {
	return d.composition
}

// IsShuffled returns true if the deck is shuffled.
func (d *Deck) IsShuffled() bool {
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
