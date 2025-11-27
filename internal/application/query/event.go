package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type Event struct {
	ID id.ID
}

type EventResult struct {
	Event entity.Event
}
