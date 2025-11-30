package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/dto"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GetChannel struct {
	Channel dto.Channel `json:"channel,omitzero"`
}

func NewGetChannel(qr query.ChannelResult) GetChannel {
	return GetChannel{
		Channel: dto.NewChannelFromEntity(qr.Channel),
	}
}
