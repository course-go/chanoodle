package memory

import (
	"cmp"
	"slices"
	"sync"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ repository.EventRepository = &EventRepository{}

type EventRepository struct {
	log zerolog.Logger

	mu      sync.Mutex
	eventID id.ID
	events  map[id.ID]entity.Event
}

func NewEventRepository(log zerolog.Logger) *EventRepository {
	return &EventRepository{
		log:    log.With().Str("component", "memory/event-repository").Logger(),
		events: make(map[id.ID]entity.Event),
	}
}

// Events implements [repository.EventRepository].
func (e *EventRepository) Events(
	filter events.Filter,
	pagination *pagination.Pagination[entity.Event],
) (events []entity.Event, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	events = make([]entity.Event, 0, len(e.events))
	for _, event := range e.events {
		if filter.Filter(event) {
			events = append(events, event)
		}
	}

	slices.SortFunc(events, func(a, b entity.Event) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if pagination != nil {
		events = pagination.Paginate(events)
	}

	return events, nil
}

// Event implements [repository.EventRepository].
func (e *EventRepository) Event(i id.ID) (event entity.Event, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	event, ok := e.events[i]
	if !ok {
		return entity.Event{}, id.ErrNoSuchEntity
	}

	return event, nil
}

// CreateEvent implements [repository.EventRepository].
func (e *EventRepository) CreateEvent(anonymousEvent entity.AnonymousEvent) (event entity.Event, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.eventID++
	event = anonymousEvent.ToEvent(e.eventID)
	e.events[e.eventID] = event

	return event, nil
}

// UpdateEvent implements [repository.EventRepository].
func (e *EventRepository) UpdateEvent(event entity.Event) (err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	_, ok := e.events[event.ID]
	if !ok {
		return id.ErrNoSuchEntity
	}

	e.events[event.ID] = event

	return nil
}
