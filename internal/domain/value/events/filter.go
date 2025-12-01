package events

import (
	"slices"
	"time"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

// Filter represents filtering parameters for [entity.Event].
type Filter struct {
	Channels []id.ID
	Genres   []id.ID
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
	for _, genreID := range f.Genres {
		if !slices.ContainsFunc(event.Genres,
			func(genre entity.Genre) bool {
				return genre.ID == genreID
			},
		) {
			return false
		}
	}

	if f.From != nil && f.From.After(event.From) {
		return false
	}

	if f.To != nil && f.To.Before(event.To) {
		return false
	}

	return true
}
