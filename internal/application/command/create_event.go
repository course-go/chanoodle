package command

import "github.com/course-go/chanoodle/internal/domain/entity"

type CreateEvent struct {
	Event entity.AnonymousEvent
}

type CreateEventResult struct {
	Event entity.Event
}
