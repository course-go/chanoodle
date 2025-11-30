package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type Events struct {
	Filter     events.Filter
	Pagination pagination.Pagination[entity.Event]
}

type EventsResult struct {
	Events []entity.Event
}
