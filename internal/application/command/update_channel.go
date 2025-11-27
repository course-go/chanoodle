package command

import "github.com/course-go/chanoodle/internal/domain/entity"

type UpdateChannel struct {
	Channel entity.Channel
}

type UpdateChannelResult struct{}
