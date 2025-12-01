package service

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/epg"
)

// EPGService represents a service for constructing EPGs and its manipulation.
type EPGService interface {
	ConstructEPG(channels []entity.Channel, events []entity.Event) (epg epg.EPG)
}
