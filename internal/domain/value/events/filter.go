package events

import (
	"slices"
	"time"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/genre"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type Filter struct {
	Channels []id.ID
	Genres   []genre.Genre
	From     *time.Time
	To       *time.Time
}

// Filter runs filtering on [entity.Event].
// It returns true when the event passes the filter.
func (f *Filter) Filter(event entity.Event) bool {
	if len(f.Channels) != 0 && !slices.Contains(f.Channels, event.Channel) {
		return false
	}

	// All genres match.
	for _, genre := range f.Genres {
		if !slices.Contains(event.Genres, genre) {
			return false
		}
	}

	if f.From != nil && f.From.Before(event.From) {
		return false
	}

	if f.To != nil && f.To.After(event.From) {
		return false
	}

	return true
}
