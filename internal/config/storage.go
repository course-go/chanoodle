package config

import (
	"slices"

	"github.com/course-go/chanoodle/internal/foundation/storage"
)

type Storage struct {
	Type storage.Type `yaml:"type"`
}

func defaultStorage() Storage {
	return Storage{
		Type: storage.Memory,
	}
}

func (s *Storage) validate() error {
	if !slices.Contains(storage.Types(), s.Type) {
		return storage.ErrUnknownType
	}

	return nil
}
