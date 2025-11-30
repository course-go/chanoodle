package storage

import "errors"

var ErrUnknownType = errors.New("unknown storage type")

type Type string

const (
	Memory Type = "memory"
)

func Types() []Type {
	return []Type{
		Memory,
	}
}
