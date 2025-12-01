package repository

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

// EventRepository represents a repository for managing [entity.Event]s.
type EventRepository interface {
	Events(
		filter events.Filter,
		pagination *pagination.Pagination[entity.Event],
	) (events []entity.Event, err error)
	Event(id id.ID) (event entity.Event, err error)
	CreateEvent(anonymousEvent entity.AnonymousEvent) (event entity.Event, err error)
	UpdateEvent(id id.ID, anonymousEvent entity.AnonymousEvent) (err error)
}
