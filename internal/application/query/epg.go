package query

import (
	"time"

	"github.com/course-go/chanoodle/internal/domain/value/epg"
)

type EPG struct {
	From time.Time
	To   time.Time
}

type EPGResult struct {
	EPG epg.EPG
}
