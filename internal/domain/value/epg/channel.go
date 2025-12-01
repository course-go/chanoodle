package epg

import "github.com/course-go/chanoodle/internal/domain/value/priority"

// Channel represents a TV channel in the EPG.
type Channel struct {
	Name     string
	Priority priority.Priority
	Events   []Event
}
