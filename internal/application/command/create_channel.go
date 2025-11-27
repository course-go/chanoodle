package command

import "github.com/course-go/chanoodle/internal/domain/entity"

type CreateChannel struct {
	Channel entity.AnonymousChannel
}

type CreateChannelResult struct {
	Channel entity.Channel
}
