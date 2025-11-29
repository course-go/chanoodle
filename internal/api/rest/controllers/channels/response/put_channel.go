package response

import "github.com/course-go/chanoodle/internal/application/command"

type PutChannel struct{}

func ParsePutChannel(cr command.UpdateChannelResult) PutChannel {
	return PutChannel{}
}
