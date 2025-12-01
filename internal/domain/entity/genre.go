package entity

import "github.com/course-go/chanoodle/internal/domain/value/id"

// AnonymousGenre represents a [Genre] without ID.
type AnonymousGenre struct {
	Name string
}

func (ag *AnonymousGenre) ToGenre(id id.ID) Genre {
	return Genre{
		ID:   id,
		Name: ag.Name,
	}
}

// Genre represents a TV genre.
type Genre struct {
	ID   id.ID
	Name string
}
