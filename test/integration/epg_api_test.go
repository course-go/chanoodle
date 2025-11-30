package integration_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg/dto"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg/response"
	"github.com/course-go/chanoodle/test/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEPGController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("GetDayEPG_ReturnsEPG", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)

		from := setup.Date()
		to := from.Add(24 * time.Hour)

		setup.Seed(t, d)

		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf(
				"/api/v1/epg?from=%d&to=%d",
				from.Unix(),
				to.Unix(),
			),
			nil,
		)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetEPG `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data.EPG.Channels, 3)
		assert.Equal(t,
			response.GetEPG{
				EPG: dto.EPG{
					Channels: []dto.Channel{
						{
							Name: "CT24",
							Events: []dto.Event{
								{
									Name: "Event A",
									From: from,
									To:   from.Add(15 * time.Minute),
								},
							},
						},
						{
							Name: "BBC One",
							Events: []dto.Event{
								{
									Name: "Event B",
									From: from.Add(1 * time.Hour),
									To:   from.Add(1*time.Hour + 45*time.Minute),
								},
							},
						},
						{
							Name: "Šlágr",
							Events: []dto.Event{
								{
									Name: "Event C",
									From: from.Add(2 * time.Hour),
									To:   from.Add(2*time.Hour + 30*time.Minute),
								},
							},
						},
					},
				},
			},
			resp.Data,
		)
	})

	t.Run("GetHourEPG_ReturnsShortEPG", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)

		from := setup.Date().Add(time.Hour)
		to := from.Add(time.Hour)

		setup.Seed(t, d)

		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf(
				"/api/v1/epg?from=%d&to=%d",
				from.Unix(),
				to.Unix(),
			),
			nil,
		)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetEPG `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data.EPG.Channels, 1)
		assert.Equal(t,
			response.GetEPG{
				EPG: dto.EPG{
					Channels: []dto.Channel{
						{
							Name: "BBC One",
							Events: []dto.Event{
								{
									Name: "Event B",
									From: from,
									To:   from.Add(45 * time.Minute),
								},
							},
						},
					},
				},
			},
			resp.Data,
		)
	})

	t.Run("GetEPGWithInvalidTimeRange_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(
			http.MethodGet,
			"/api/v1/epg?from=invalid&to=invalid",
			nil,
		)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp struct {
			Error string `json:"error"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.NotEmpty(t, resp.Error)
	})
}
