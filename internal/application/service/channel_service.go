package service

import (
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/entity"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ service.ChannelService = &ChannelService{}

type ChannelService struct {
	log            zerolog.Logger
	channelService domain.ChannelService
}

func NewChannelService(log zerolog.Logger, channelService domain.ChannelService) *ChannelService {
	return &ChannelService{
		log:            log.With().Str("component", "application/channel-service").Logger(),
		channelService: channelService,
	}
}

// Channel implements [service.ChannelService].
func (cs *ChannelService) Channel(q query.Channel) (r query.ChannelResult, err error) {
	channel, err := cs.channelService.Channel(q.ID)
	if err != nil {
		return r, err
	}

	r.Channel = channel
	return r, nil
}

// Channels implements [service.ChannelService].
func (cs *ChannelService) Channels(q query.Channels) (r query.ChannelsResult, err error) {
	pag := pagination.New[entity.Channel](0, 0)
	channels, err := cs.channelService.Channels(q.Filter, pag)
	if err != nil {
		return r, err
	}

	r.Channels = channels
	return r, nil
}

// CreateChannel implements [service.ChannelService].
func (cs *ChannelService) CreateChannel(c command.CreateChannel) (r command.CreateChannelResult, err error) {
	channel, err := cs.channelService.CreateChannel(c.Channel)
	if err != nil {
		return r, err
	}

	r.Channel = channel
	return r, nil
}

// UpdateChannel implements [service.ChannelService].
func (cs *ChannelService) UpdateChannel(c command.UpdateChannel) (r command.UpdateChannelResult, err error) {
	err = cs.channelService.UpdateChannel(c.Channel)
	if err != nil {
		return r, err
	}

	return r, nil
}
