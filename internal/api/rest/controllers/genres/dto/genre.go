package dto

import "github.com/course-go/chanoodle/internal/domain/entity"

type AnonymousGenre struct {
	Name string `json:"name,omitzero" validate:"required"`
}

type Genre struct {
	ID   int    `json:"id,omitzero"   validate:"required"`
	Name string `json:"name,omitzero" validate:"required"`
}

func NewGenreFromEntity(genre entity.Genre) Genre {
	return Genre{
		ID:   int(genre.ID),
		Name: genre.Name,
	}
}
