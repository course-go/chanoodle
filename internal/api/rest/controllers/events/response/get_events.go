package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/dto"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GetEvents struct {
	Events []dto.Event `json:"events"`
}

func ParseGetEvents(qr query.EventsResult) GetEvents {
	events := make([]dto.Event, 0, len(qr.Events))
	for _, event := range qr.Events {
		events = append(events, dto.NewEventFromEntity(event))
	}

	return GetEvents{
		Events: events,
	}
}
