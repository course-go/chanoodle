package response

import "github.com/course-go/chanoodle/internal/application/command"

type PutEvent struct{}

func ParsePutEvent(cr command.UpdateEventResult) PutEvent {
	return PutEvent{}
}
