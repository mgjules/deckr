package deck

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

// WithCodes returns an option that sets the card codes of the deck.
func WithCodes(codes ...string) Option {
	return func(d *Deck) {
		d.codes = codes
	}
}

// WithComposition returns an option that sets the composition of the deck.
func WithComposition(comp string) Option {
	return func(d *Deck) {
		d.composition = comp
	}
}
