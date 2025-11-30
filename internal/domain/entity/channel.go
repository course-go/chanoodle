package entity

import (
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type AnonymousChannel struct {
	Name   string
	Genres []Genre
}

func (ac *AnonymousChannel) ToChannel(id id.ID) Channel {
	return Channel{
		ID:     id,
		Name:   ac.Name,
		Genres: ac.Genres,
	}
}

type Channel struct {
	ID     id.ID
	Name   string
	Genres []Genre
}
