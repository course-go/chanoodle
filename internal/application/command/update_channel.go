package command

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
)

type UpdateChannel struct {
	ID      id.ID
	Channel entity.AnonymousChannel
}

type UpdateChannelResult struct{}
