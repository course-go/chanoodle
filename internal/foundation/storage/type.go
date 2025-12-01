package storage

import "errors"

var ErrUnknownType = errors.New("unknown storage type")

// Type represents the type of storage implementation.
type Type string

const (
	Memory Type = "memory"
)

func Types() []Type {
	return []Type{
		Memory,
	}
}
