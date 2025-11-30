package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/rs/zerolog"
)

var _ service.EPGService = &EPGService{}

type EPGService struct {
	log               zerolog.Logger
	epgService        domain.EPGService
	channelRepository repository.ChannelRepository
	eventRepository   repository.EventRepository
}

func NewEPGService(
	log zerolog.Logger,
	epgService domain.EPGService,
	channelRepository repository.ChannelRepository,
	eventRepository repository.EventRepository,
) *EPGService {
	return &EPGService{
		log:               log.With().Str("component", "application/epg-service").Logger(),
		epgService:        epgService,
		channelRepository: channelRepository,
		eventRepository:   eventRepository,
	}
}

// EPG implements [service.EPGService].
func (es *EPGService) EPG(q query.EPG) (r query.EPGResult, err error) {
	channels, err := es.channelRepository.Channels(channels.Filter{}, nil)
	if err != nil {
		return query.EPGResult{}, fmt.Errorf("failed getting channels from repository: %w", err)
	}

	channelIDs := make([]id.ID, 0, len(channels))
	for _, channel := range channels {
		channelIDs = append(channelIDs, channel.ID)
	}

	events, err := es.eventRepository.Events(
		events.Filter{
			Channels: channelIDs,
			From:     &q.From,
			To:       &q.To,
		},
		nil,
	)
	if err != nil {
		return query.EPGResult{}, fmt.Errorf("failed getting events from repository: %w", err)
	}

	epg := es.epgService.ConstructEPG(channels, events)

	return query.EPGResult{
		EPG: epg,
	}, nil
}
