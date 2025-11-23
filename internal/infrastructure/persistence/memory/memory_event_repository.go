package memory

import (
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/rs/zerolog"
)

var _ repository.EventRepository = &EventRepository{}

type EventRepository struct {
	log zerolog.Logger
}

func NewEventRepository(log zerolog.Logger) EventRepository {
	return EventRepository{
		log: log.With().Str("component", "memory/event-repository").Logger(),
	}
}
