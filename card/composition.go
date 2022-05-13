package card

// Composition represents a collection of ranks and suits. (e.g french)
type Composition struct {
	ranks Ranks
	suits Suits
}

// NewComposition returns a new card composition given a collection of ranks and suits.
func NewComposition(r Ranks, s Suits) Composition {
	return Composition{
		ranks: r,
		suits: s,
	}
}

// Ranks returns a collection of ranks.
func (c Composition) Ranks() Ranks {
	return c.ranks
}

// Suits returns a collection of suits.
func (c Composition) Suits() Suits {
	return c.suits
}
