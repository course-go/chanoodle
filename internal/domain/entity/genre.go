package entity

import "github.com/course-go/chanoodle/internal/domain/value/id"

type AnonymousGenre struct {
	Name string
}

func (ag *AnonymousGenre) ToGenre(id id.ID) Genre {
	return Genre{
		ID:   id,
		Name: ag.Name,
	}
}

type Genre struct {
	ID   id.ID
	Name string
}
