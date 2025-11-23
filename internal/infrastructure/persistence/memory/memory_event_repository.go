package memory

import "github.com/course-go/chanoodle/internal/domain/interfaces/repository"

var _ repository.EventRepository = &EventRepository{}

type EventRepository struct{}

func NewEventRepository() EventRepository {
	return EventRepository{}
}
