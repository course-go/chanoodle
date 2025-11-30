package epg

import "time"

type Event struct {
	Name string
	From time.Time
	To   time.Time
}
