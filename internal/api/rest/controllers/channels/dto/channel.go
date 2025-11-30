package dto

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/domain/entity"
)

type AnonymousChannel struct {
	Name     string               `json:"name,omitzero"     validate:"required"`
	Priority int                  `json:"priority,omitzero"`
	Genres   []dto.AnonymousGenre `json:"genres,omitempty"`
}

type Channel struct {
	ID       int         `json:"id,omitzero"       validate:"required"`
	Name     string      `json:"name,omitzero"     validate:"required"`
	Priority int         `json:"priority,omitzero"`
	Genres   []dto.Genre `json:"genres,omitempty"`
}

func NewChannelFromEntity(channel entity.Channel) Channel {
	genres := make([]dto.Genre, 0, len(channel.Genres))
	for _, genre := range channel.Genres {
		genres = append(genres, dto.NewGenreFromEntity(genre))
	}

	return Channel{
		ID:       int(channel.ID),
		Name:     channel.Name,
		Priority: int(channel.Priority),
		Genres:   genres,
	}
}
