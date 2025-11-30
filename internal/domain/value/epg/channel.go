package epg

import "github.com/course-go/chanoodle/internal/domain/value/priority"

type Channel struct {
	Name     string
	Priority priority.Priority
	Events   []Event
}
