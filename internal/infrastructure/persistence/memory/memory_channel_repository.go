package memory

import (
	"sync"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ repository.ChannelRepository = &ChannelRepository{}

type ChannelRepository struct {
	log zerolog.Logger

	mu        sync.Mutex
	channelID id.ID
	channels  map[id.ID]entity.Channel
}

func NewChannelRepository(log zerolog.Logger) *ChannelRepository {
	return &ChannelRepository{
		log:      log.With().Str("component", "memory/channel-repository").Logger(),
		channels: make(map[id.ID]entity.Channel),
	}
}

// Channels implements [repository.ChannelRepository].
func (c *ChannelRepository) Channels(
	filter channels.Filter,
	pag pagination.Pagination[entity.Channel],
) (channels []entity.Channel, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		if filter.Filter(channel) {
			channels = append(channels, channel)
		}
	}

	channels = pag.Paginate(channels)

	return channels, nil
}

// Channel implements [repository.ChannelRepository].
func (c *ChannelRepository) Channel(channelID id.ID) (channel entity.Channel, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	channel, ok := c.channels[channelID]
	if !ok {
		return entity.Channel{}, id.ErrNoSuchEntity
	}

	return channel, nil
}

// CreateChannel implements [repository.ChannelRepository].
func (c *ChannelRepository) CreateChannel(
	anonymousChannels []entity.AnonymousChannel,
) (channel entity.Channel, err error) {
	if len(anonymousChannels) == 0 {
		return entity.Channel{}, id.ErrNoSuchEntity
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// Create all channels in the slice, return the last one
	for _, anonymousChannel := range anonymousChannels {
		c.channelID++
		channel = anonymousChannel.ToChannel(c.channelID)
		c.channels[c.channelID] = channel
	}

	return channel, nil
}

// UpdateChannel implements [repository.ChannelRepository].
func (c *ChannelRepository) UpdateChannel(channel entity.Channel) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.channels[channel.ID]
	if !ok {
		return id.ErrNoSuchEntity
	}

	c.channels[channel.ID] = channel

	return nil
}
