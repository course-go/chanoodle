package entity

import (
	"time"

	"github.com/course-go/chanoodle/internal/domain/value/genre"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type AnonymousEvent struct {
	Name    string
	Channel id.ID
	From    time.Time
	To      time.Time
	Genres  []genre.Genre
}

func (ae *AnonymousEvent) ToEvent(id id.ID) Event {
	return Event{
		ID:      id,
		Channel: ae.Channel,
		Name:    ae.Name,
		From:    ae.From,
		To:      ae.To,
		Genres:  ae.Genres,
	}
}

type Event struct {
	ID      id.ID
	Channel id.ID
	Name    string
	From    time.Time
	To      time.Time
	Genres  []genre.Genre
}
