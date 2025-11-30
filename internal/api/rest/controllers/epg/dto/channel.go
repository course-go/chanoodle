package dto

import "github.com/course-go/chanoodle/internal/domain/value/epg"

type Channel struct {
	Name   string  `json:"name,omitzero"`
	Events []Event `json:"events,omitempty"`
}

func NewChannelFromValue(value epg.Channel) Channel {
	events := make([]Event, 0, len(value.Events))
	for _, event := range value.Events {
		events = append(events, NewEventFromValue(event))
	}

	return Channel{
		Name:   value.Name,
		Events: events,
	}
}
