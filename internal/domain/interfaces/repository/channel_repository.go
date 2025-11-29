package repository

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type ChannelRepository interface {
	Channels(
		filter channels.Filter,
		pagination pagination.Pagination[entity.Channel],
	) (channels []entity.Channel, err error)
	Channel(id id.ID) (channel entity.Channel, err error)
	CreateChannel(anonymousChannel []entity.AnonymousChannel) (channel entity.Channel, err error)
	UpdateChannel(channel entity.Channel) (err error)
}
