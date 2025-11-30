package dto

import "github.com/course-go/chanoodle/internal/domain/entity"

type AnonymousGenre struct {
	Name string `json:"name,omitzero"`
}

type Genre struct {
	ID   int    `json:"id,omitzero"`
	Name string `json:"name,omitzero"`
}

func NewGenreFromEntity(genre entity.Genre) Genre {
	return Genre{
		ID:   int(genre.ID),
		Name: genre.Name,
	}
}
