package service

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type EventService interface {
	Events(filter events.Filter, pagination pagination.Pagination[entity.Event]) ([]entity.Event, error)
	Event(id id.ID) (entity.Event, error)
	CreateEvent(anonymousEvent entity.AnonymousEvent) (entity.Event, error)
	UpdateEvent(event entity.Event) error
}
