package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type Channels struct {
	Filter     channels.Filter
	Pagination pagination.Pagination[entity.Channel]
}

type ChannelsResult struct {
	Channels []entity.Channel
}
