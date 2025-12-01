package epg

import "time"

// Event represents a program event in the EPG.
type Event struct {
	Name string
	From time.Time
	To   time.Time
}
