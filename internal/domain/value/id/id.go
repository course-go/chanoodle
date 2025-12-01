package id

import "errors"

var ErrNoSuchEntity = errors.New("no entity with such ID")

// ID represents a unique identifier for entities.
type ID int
