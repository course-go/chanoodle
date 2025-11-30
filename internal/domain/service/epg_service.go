package service

import (
	"cmp"
	"maps"
	"slices"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/epg"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/rs/zerolog"
)

var _ service.EPGService = &EPGService{}

type EPGService struct {
	log zerolog.Logger
}

func NewEPGService(log zerolog.Logger) *EPGService {
	return &EPGService{
		log: log.With().Str("component", "domain/epg-service").Logger(),
	}
}

// ConstructEPG implements [service.EPGService].
func (es *EPGService) ConstructEPG(channels []entity.Channel, events []entity.Event) epg.EPG {
	epgChannelsMap := make(map[id.ID]epg.Channel, len(channels))
	for _, channel := range channels {
		epgChannelsMap[channel.ID] = epg.Channel{
			Name: channel.Name,
		}
	}

	for _, event := range events {
		channel, ok := epgChannelsMap[event.Channel]
		if !ok {
			es.log.Warn().
				Int("channelID", int(event.Channel)).
				Int("eventID", int(event.ID)).
				Msg("event's channel not found among provided channels - skipping event")

			continue
		}

		channel.Events = append(channel.Events, epg.Event{
			Name: event.Name,
			From: event.From,
			To:   event.To,
		})
	}

	// Do not display channels without events.
	emptyChannelsIDs := make([]id.ID, 0)

	for id, channel := range epgChannelsMap {
		if len(channel.Events) == 0 {
			emptyChannelsIDs = append(emptyChannelsIDs, id)
		}
	}

	for _, emptyChannelID := range emptyChannelsIDs {
		delete(epgChannelsMap, emptyChannelID)
	}

	epgChannels := slices.Collect(maps.Values(epgChannelsMap))

	// Sort channel events based on their start.
	for _, epgChannel := range epgChannels {
		slices.SortFunc(epgChannel.Events, func(a, b epg.Event) int {
			return cmp.Compare(a.From.Unix(), b.From.Unix())
		})
	}

	// Sort channels based on their name.
	slices.SortFunc(epgChannels, func(a, b epg.Channel) int {
		return cmp.Compare(a.Name, b.Name)
	})

	return epg.EPG{
		Channels: epgChannels,
	}
}
