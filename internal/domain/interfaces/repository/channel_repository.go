package repository

import (
	"context"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

// ChannelRepository represents a repository for managing [entity.Channel]s.
type ChannelRepository interface {
	Channels(
		ctx context.Context,
		filter channels.Filter,
		pagination *pagination.Pagination[entity.Channel],
	) (channels []entity.Channel, err error)
	Channel(
		ctx context.Context,
		id id.ID,
	) (channel entity.Channel, err error)
	CreateChannel(
		ctx context.Context,
		anonymousChannel entity.AnonymousChannel,
	) (channel entity.Channel, err error)
	UpdateChannel(
		ctx context.Context,
		id id.ID,
		anonymousChannel entity.AnonymousChannel,
	) (err error)
}
