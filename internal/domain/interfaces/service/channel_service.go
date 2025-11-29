package service

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type ChannelService interface {
	Channels(filter channels.Filter, pagination pagination.Pagination[entity.Channel]) ([]entity.Channel, error)
	Channel(id id.ID) (entity.Channel, error)
	CreateChannel(anonymousChannel entity.AnonymousChannel) (entity.Channel, error)
	UpdateChannel(channel entity.Channel) error
}
