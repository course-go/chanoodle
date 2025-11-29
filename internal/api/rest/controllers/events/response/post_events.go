package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/dto"
	"github.com/course-go/chanoodle/internal/application/command"
)

type PostEvents struct {
	Event dto.Event `json:"event"`
}

func ParsePostEvents(cr command.CreateEventResult) PostEvents {
	return PostEvents{
		Event: dto.NewEventFromEntity(cr.Event),
	}
}
