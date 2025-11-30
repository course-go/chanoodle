package entity

import (
	"time"

	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type AnonymousEvent struct {
	Name    string
	Channel id.ID
	From    time.Time
	To      time.Time
	Genres  []id.ID
}

func (ae *AnonymousEvent) ToEvent(id id.ID, genres []Genre) Event {
	return Event{
		ID:      id,
		Channel: ae.Channel,
		Name:    ae.Name,
		From:    ae.From,
		To:      ae.To,
		Genres:  genres,
	}
}

type Event struct {
	ID      id.ID
	Channel id.ID
	Name    string
	From    time.Time
	To      time.Time
	Genres  []Genre
}
