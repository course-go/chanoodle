package service

import (
	"context"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/query"
)

// EventService defines all supported Event related use-cases.
type EventService interface {
	Events(ctx context.Context, q query.Events) (r query.EventsResult, err error)
	Event(ctx context.Context, q query.Event) (r query.EventResult, err error)
	CreateEvent(ctx context.Context, c command.CreateEvent) (r command.CreateEventResult, err error)
	UpdateEvent(ctx context.Context, c command.UpdateEvent) (r command.UpdateEventResult, err error)
}
