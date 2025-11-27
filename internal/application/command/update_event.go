package command

import "github.com/course-go/chanoodle/internal/domain/entity"

type UpdateEvent struct {
	Event entity.Event
}

type UpdateEventResult struct{}
