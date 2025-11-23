package entity

import "time"

type Event struct {
	ID   string
	Name string
	From time.Time
	To   time.Time
}
