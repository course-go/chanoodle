package service

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/epg"
)

type EPGService interface {
	ConstructEPG(channels []entity.Channel, events []entity.Event) (epg epg.EPG)
}
