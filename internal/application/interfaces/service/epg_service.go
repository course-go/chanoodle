package service

import (
	"github.com/course-go/chanoodle/internal/application/query"
)

type EPGService interface {
	EPG(q query.EPG) (r query.EPGResult, err error)
}
