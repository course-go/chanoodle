package memory

import (
	"cmp"
	"fmt"
	"slices"
	"sync"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ repository.MediaRepository = &MediaRepository{}

type MediaRepository struct {
	log zerolog.Logger

	mu        sync.Mutex
	channelID id.ID
	eventID   id.ID
	genreID   id.ID
	channels  map[id.ID]entity.Channel
	events    map[id.ID]entity.Event
	genres    map[id.ID]entity.Genre
}

func NewMediaRepository(log zerolog.Logger) *MediaRepository {
	return &MediaRepository{
		log:      log.With().Str("component", "memory/media-repository").Logger(),
		channels: make(map[id.ID]entity.Channel),
		events:   make(map[id.ID]entity.Event),
		genres:   make(map[id.ID]entity.Genre),
	}
}

// Channels implements [repository.ChannelRepository].
func (mr *MediaRepository) Channels(
	filter channels.Filter,
	pagination *pagination.Pagination[entity.Channel],
) (channels []entity.Channel, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	channels = make([]entity.Channel, 0, len(mr.channels))
	for _, channel := range mr.channels {
		if filter.Filter(channel) {
			channels = append(channels, channel)
		}
	}

	slices.SortFunc(channels, func(a, b entity.Channel) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if pagination != nil {
		channels = pagination.Paginate(channels)
	}

	return channels, nil
}

// Channel implements [repository.ChannelRepository].
func (mr *MediaRepository) Channel(channelID id.ID) (channel entity.Channel, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	channel, ok := mr.channels[channelID]
	if !ok {
		return entity.Channel{}, id.ErrNoSuchEntity
	}

	return channel, nil
}

// CreateChannel implements [repository.ChannelRepository].
func (mr *MediaRepository) CreateChannel(
	anonymousChannel entity.AnonymousChannel,
) (channel entity.Channel, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	genres, err := mr.genresByIDs(anonymousChannel.Genres)
	if err != nil {
		return entity.Channel{}, fmt.Errorf("failed getting channel genres: %w", err)
	}

	mr.channelID++
	channel = anonymousChannel.ToChannel(mr.channelID, genres)
	mr.channels[mr.channelID] = channel

	return channel, nil
}

// UpdateChannel implements [repository.ChannelRepository].
func (mr *MediaRepository) UpdateChannel(i id.ID, anonymousChannel entity.AnonymousChannel) (err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	_, ok := mr.channels[i]
	if !ok {
		return id.ErrNoSuchEntity
	}

	genres, err := mr.genresByIDs(anonymousChannel.Genres)
	if err != nil {
		return fmt.Errorf("failed getting event genres: %w", err)
	}

	channel := anonymousChannel.ToChannel(i, genres)

	mr.channels[i] = channel

	return nil
}

// Events implements [repository.EventRepository].
func (mr *MediaRepository) Events(
	filter events.Filter,
	pagination *pagination.Pagination[entity.Event],
) (events []entity.Event, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	events = make([]entity.Event, 0, len(mr.events))
	for _, event := range mr.events {
		if filter.Filter(event) {
			events = append(events, event)
		}
	}

	slices.SortFunc(events, func(a, b entity.Event) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if pagination != nil {
		events = pagination.Paginate(events)
	}

	return events, nil
}

// Event implements [repository.EventRepository].
func (mr *MediaRepository) Event(i id.ID) (event entity.Event, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	event, ok := mr.events[i]
	if !ok {
		return entity.Event{}, id.ErrNoSuchEntity
	}

	return event, nil
}

// CreateEvent implements [repository.EventRepository].
func (mr *MediaRepository) CreateEvent(anonymousEvent entity.AnonymousEvent) (event entity.Event, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	genres, err := mr.genresByIDs(anonymousEvent.Genres)
	if err != nil {
		return entity.Event{}, fmt.Errorf("failed getting channel genres: %w", err)
	}

	mr.eventID++
	event = anonymousEvent.ToEvent(mr.eventID, genres)
	mr.events[mr.eventID] = event

	return event, nil
}

// UpdateEvent implements [repository.EventRepository].
func (mr *MediaRepository) UpdateEvent(i id.ID, anonymousEvent entity.AnonymousEvent) (err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	_, ok := mr.events[i]
	if !ok {
		return id.ErrNoSuchEntity
	}

	genres, err := mr.genresByIDs(anonymousEvent.Genres)
	if err != nil {
		return fmt.Errorf("failed getting event genres: %w", err)
	}

	event := anonymousEvent.ToEvent(i, genres)
	mr.events[i] = event

	return nil
}

// Genres implements [repository.GenreRepository].
func (mr *MediaRepository) Genres(pagination *pagination.Pagination[entity.Genre]) (genres []entity.Genre, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	genres = make([]entity.Genre, 0, len(mr.genres))
	for _, genre := range mr.genres {
		genres = append(genres, genre)
	}

	slices.SortFunc(genres, func(a, b entity.Genre) int {
		return cmp.Compare(a.ID, b.ID)
	})

	if pagination != nil {
		genres = pagination.Paginate(genres)
	}

	return genres, nil
}

// GetOrCreateGenre implements [repository.GenreRepository].
func (mr *MediaRepository) GetOrCreateGenre(anonymousGenre entity.AnonymousGenre) (genre entity.Genre, err error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	for _, genre := range mr.genres {
		if genre.Name == anonymousGenre.Name {
			return genre, nil
		}
	}

	mr.genreID++
	genre = anonymousGenre.ToGenre(mr.genreID)
	mr.genres[mr.genreID] = genre

	return genre, nil
}

func (mr *MediaRepository) genresByIDs(genreIDs []id.ID) (genres []entity.Genre, err error) {
	genres = make([]entity.Genre, 0, len(genreIDs))
	for _, genreID := range genreIDs {
		genre, ok := mr.genres[genreID]
		if !ok {
			return nil, fmt.Errorf("wrong genre: %w", id.ErrNoSuchEntity)
		}

		genres = append(genres, genre)
	}

	return genres, nil
}
