package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/events"
)

type Events struct {
	Filter events.Filter
}

type EventsResult struct {
	Events []entity.Event
}
