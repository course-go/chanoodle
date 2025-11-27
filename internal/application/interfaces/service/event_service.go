package service

import (
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/query"
)

type EventService interface {
	Events(q query.Events) (r query.EventsResult, err error)
	Event(q query.Event) (r query.EventResult, err error)
	CreateEvent(c command.CreateEvent) (r command.CreateEventResult, err error)
	UpdateEvent(c command.UpdateEvent) (r command.UpdateEventResult, err error)
}
