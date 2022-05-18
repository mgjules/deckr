package errs

import "errors"

// ErrDeckNotFound is the error returned when a deck is not found.
var ErrDeckNotFound = errors.New("deck not found")
