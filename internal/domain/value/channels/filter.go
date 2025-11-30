package channels

import (
	"slices"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type Filter struct {
	Genres []id.ID
}

// Filter runs filtering on [entity.Channel].
// It returns true when the channel passes the filter.
func (f *Filter) Filter(channel entity.Channel) bool {
	// All genres match.
	for _, genreID := range f.Genres {
		if !slices.ContainsFunc(channel.Genres,
			func(genre entity.Genre) bool {
				return genre.ID == genreID
			},
		) {
			return false
		}
	}

	return true
}
