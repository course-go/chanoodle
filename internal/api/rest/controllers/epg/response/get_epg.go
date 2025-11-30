package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg/dto"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GetEPG struct {
	EPG dto.EPG `json:"epg"`
}

func NewGetEPG(qr query.EPGResult) GetEPG {
	return GetEPG{
		EPG: dto.NewEPGFromValue(qr.EPG),
	}
}
