package dto

import (
	"time"

	"github.com/course-go/chanoodle/internal/domain/value/epg"
)

type Event struct {
	Name string    `json:"name"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func NewEventFromValue(value epg.Event) Event {
	return Event{
		Name: value.Name,
		From: value.From,
		To:   value.To,
	}
}
