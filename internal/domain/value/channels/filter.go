package channels

import (
	"slices"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/genre"
)

type Filter struct {
	Genres []genre.Genre
}

// Filter runs filtering on [entity.Channel].
// It returns true when the channel passes the filter.
func (f *Filter) Filter(channel entity.Channel) bool {
	// All genres match.
	for _, genre := range f.Genres {
		if !slices.Contains(channel.Genres, genre) {
			return false
		}
	}

	return true
}
