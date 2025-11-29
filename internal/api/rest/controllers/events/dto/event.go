package dto

import (
	"time"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/domain/entity"
)

type AnonymousEvent struct {
	Name    string      `json:"name,omitzero"`
	Channel int         `json:"channel,omitzero"`
	From    time.Time   `json:"from,omitzero"`
	To      time.Time   `json:"to,omitzero"`
	Genres  []dto.Genre `json:"genres,omitempty"`
}

type Event struct {
	ID      int         `json:"id,omitzero"`
	Channel int         `json:"channel,omitzero"`
	Name    string      `json:"name,omitzero"`
	From    time.Time   `json:"from,omitzero"`
	To      time.Time   `json:"to,omitzero"`
	Genres  []dto.Genre `json:"genres,omitempty"`
}

func NewEventFromEntity(event entity.Event) Event {
	genres := make([]dto.Genre, 0, len(event.Genres))
	for _, genre := range event.Genres {
		genres = append(genres, dto.NewGenreFromEntity(genre))
	}

	return Event{
		ID:      int(event.ID),
		Channel: int(event.Channel),
		Name:    event.Name,
		From:    event.From,
		To:      event.To,
		Genres:  genres,
	}
}
