package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
)

type Channels struct {
	Filter channels.Filter
}

type ChannelsResult struct {
	Channels []entity.Channel
}
