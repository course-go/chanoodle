package command

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type UpdateEvent struct {
	ID    id.ID
	Event entity.AnonymousEvent
}

type UpdateEventResult struct{}
