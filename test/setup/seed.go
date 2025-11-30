package setup

import (
	"testing"
	"time"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/priority"
	"github.com/stretchr/testify/require"
)

const (
	ShortEventDuration  = 15 * time.Minute
	MediumEventDuration = 30 * time.Minute
	LongEventDuration   = 45 * time.Minute
)

func Seed(t *testing.T, d Dependencies) {
	t.Helper()

	genres := SeedGenres(t, d)
	channels := SeedChannels(t, d, genres)
	SeedEvents(t, d, genres, channels)
}

func SeedGenres(t *testing.T, d Dependencies) []entity.Genre {
	t.Helper()

	genreNames := []string{"action", "romance", "comedy", "drama", "thriller"}
	genres := make([]entity.Genre, 0, len(genreNames))

	for _, name := range genreNames {
		genre, err := d.GenreRepository.GetOrCreateGenre(entity.AnonymousGenre{
			Name: name,
		})
		require.NoError(t, err)

		genres = append(genres, genre)
	}

	return genres
}

func SeedChannels(t *testing.T, d Dependencies, genres []entity.Genre) []entity.Channel {
	t.Helper()

	channelNames := []string{"CT24", "BBC One", "Šlágr", "LeoTV"}
	channelPriorities := []priority.Priority{100, 50, 30, 120}
	channels := make([]entity.Channel, 0, len(channelNames))

	for i, name := range channelNames {
		channel, err := d.ChannelRepository.CreateChannel(entity.AnonymousChannel{
			Name:     name,
			Priority: channelPriorities[i],
			Genres:   []entity.Genre{genres[i%len(genres)]},
		})
		require.NoError(t, err)

		channels = append(channels, channel)
	}

	return channels
}

func SeedEvents(t *testing.T, d Dependencies, genres []entity.Genre, channels []entity.Channel) []entity.Event {
	t.Helper()

	date := Date()

	eventNames := []string{"Event A", "Event B", "Event C"}
	events := make([]entity.Event, 0, len(eventNames))

	for i := range 3 {
		from := date.Add(time.Duration(i) * time.Hour)

		var to time.Time

		switch {
		case i%3 == 0:
			to = from.Add(ShortEventDuration)
		case i%2 == 0:
			to = from.Add(MediumEventDuration)
		default:
			to = from.Add(LongEventDuration)
		}

		event, err := d.EventRepository.CreateEvent(entity.AnonymousEvent{
			Name:    eventNames[i],
			Channel: channels[i%len(channels)].ID,
			From:    from,
			To:      to,
			Genres:  []entity.Genre{genres[i%len(genres)]},
		})
		require.NoError(t, err)

		events = append(events, event)
	}

	return events
}
