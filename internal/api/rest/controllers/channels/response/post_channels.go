package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/dto"
	"github.com/course-go/chanoodle/internal/application/command"
)

type PostChannels struct {
	Channel dto.Channel `json:"channel,omitzero"`
}

func ParsePostChannels(cr command.CreateChannelResult) PostChannels {
	return PostChannels{
		Channel: dto.NewChannelFromEntity(cr.Channel),
	}
}
