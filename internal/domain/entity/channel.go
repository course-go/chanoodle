package entity

import (
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/priority"
)

type AnonymousChannel struct {
	Name     string
	Priority priority.Priority
	Genres   []id.ID
}

func (ac *AnonymousChannel) ToChannel(id id.ID, genres []Genre) Channel {
	prio := ac.Priority
	if prio == 0 {
		prio = priority.DefaultPriority
	}

	return Channel{
		ID:       id,
		Name:     ac.Name,
		Priority: prio,
		Genres:   genres,
	}
}

type Channel struct {
	ID       id.ID
	Name     string
	Priority priority.Priority
	Genres   []Genre
}
