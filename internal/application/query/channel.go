package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type Channel struct {
	ID id.ID
}

type ChannelResult struct {
	Channel entity.Channel
}
