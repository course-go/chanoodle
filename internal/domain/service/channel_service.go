package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ service.ChannelService = &ChannelService{}

type ChannelService struct {
	log               zerolog.Logger
	channelRepository repository.ChannelRepository
}

func NewChannelService(log zerolog.Logger, channelRepository repository.ChannelRepository) *ChannelService {
	return &ChannelService{
		log:               log.With().Str("component", "domain/channel-service").Logger(),
		channelRepository: channelRepository,
	}
}

// Channels implements [service.ChannelService].
func (cs *ChannelService) Channels(
	filter channels.Filter,
	pagination pagination.Pagination[entity.Channel],
) ([]entity.Channel, error) {
	channels, err := cs.channelRepository.Channels(filter, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed fetching channels from repository: %w", err)
	}

	return channels, nil
}

// Channel implements [service.ChannelService].
func (cs *ChannelService) Channel(id id.ID) (entity.Channel, error) {
	channel, err := cs.channelRepository.Channel(id)
	if err != nil {
		return entity.Channel{}, fmt.Errorf("failed fetching channel from repository: %w", err)
	}

	return channel, nil
}

// CreateChannel implements [service.ChannelService].
func (cs *ChannelService) CreateChannel(anonymousChannel entity.AnonymousChannel) (entity.Channel, error) {
	channel, err := cs.channelRepository.CreateChannel([]entity.AnonymousChannel{anonymousChannel})
	if err != nil {
		return entity.Channel{}, fmt.Errorf("failed creating channel in repository: %w", err)
	}

	return channel, nil
}

// UpdateChannel implements [service.ChannelService].
func (cs *ChannelService) UpdateChannel(channel entity.Channel) error {
	err := cs.channelRepository.UpdateChannel(channel)
	if err != nil {
		return fmt.Errorf("failed updating channel in repository: %w", err)
	}

	return nil
}
