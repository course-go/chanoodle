package epg

import (
	"fmt"
	"net/http"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg/request"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg/response"
	application "github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// API represents the EPG REST API subgroup.
type API struct {
	log        zerolog.Logger
	epgService application.EPGService
}

func NewAPI(log zerolog.Logger, epgService application.EPGService) API {
	return API{
		log:        log.With().Str("component", "api-rest/epg-api").Logger(),
		epgService: epgService,
	}
}

func (a *API) MountRoutes(e *echo.Group) {
	epg := e.Group("/epg")

	epg.GET("", a.getEPGController)
}

func (a *API) getEPGController(c echo.Context) error {
	ctx := c.Request().Context()

	q, err := request.ParseGetEPG(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	qr, err := a.epgService.EPG(ctx, q)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed querying genres: %w", err)
	}

	data := response.NewGetEPG(qr)

	err = c.JSON(http.StatusOK, common.NewDataResponse(data))
	if err != nil {
		return fmt.Errorf("failed sending response: %w", err)
	}

	return nil
}
