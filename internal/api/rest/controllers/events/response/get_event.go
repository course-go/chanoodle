package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/dto"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GetEvent struct {
	Event dto.Event `json:"event"`
}

func ParseGetEvent(qr query.EventResult) GetEvent {
	return GetEvent{
		Event: dto.NewEventFromEntity(qr.Event),
	}
}
