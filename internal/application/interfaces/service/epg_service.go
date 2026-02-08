package service

import (
	"context"

	"github.com/course-go/chanoodle/internal/application/query"
)

// EPGService defines all supported EPG related use-cases.
type EPGService interface {
	EPG(ctx context.Context, q query.EPG) (r query.EPGResult, err error)
}
