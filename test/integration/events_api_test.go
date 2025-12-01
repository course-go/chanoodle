package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/dto"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/response"
	dtogenre "github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/api/rest/middleware/auth"
	"github.com/course-go/chanoodle/test/setup"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEventsController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("GetAllEvents_ReturnsEvents", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		date := setup.Date()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/events?limit=100", nil)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetEvents `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			response.GetEvents{
				Events: []dto.Event{
					{
						ID:      1,
						Channel: 1,
						Name:    "Event A",
						From:    date,
						To:      date.Add(15 * time.Minute),
						Genres: []dtogenre.Genre{
							{
								ID:   1,
								Name: "action",
							},
						},
					},
					{
						ID:      2,
						Channel: 2,
						Name:    "Event B",
						From:    date.Add(1 * time.Hour),
						To:      date.Add(1*time.Hour + 45*time.Minute),
						Genres: []dtogenre.Genre{
							{
								ID:   2,
								Name: "romance",
							},
						},
					},
					{
						ID:      3,
						Channel: 3,
						Name:    "Event C",
						From:    date.Add(2 * time.Hour),
						To:      date.Add(2*time.Hour + 30*time.Minute),
						Genres: []dtogenre.Genre{
							{
								ID:   3,
								Name: "comedy",
							},
						},
					},
				},
			},
			resp.Data,
		)
	})

	t.Run("GetLimitedEventsWithOffset_ReturnsEvents", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		date := setup.Date()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/events?limit=2&offset=1", nil)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetEvents `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data.Events, 2)
		assert.Equal(t,
			response.GetEvents{
				Events: []dto.Event{
					{
						ID:      2,
						Channel: 2,
						Name:    "Event B",
						From:    date.Add(1 * time.Hour),
						To:      date.Add(1*time.Hour + 45*time.Minute),
						Genres: []dtogenre.Genre{
							{
								ID:   2,
								Name: "romance",
							},
						},
					},
					{
						ID:      3,
						Channel: 3,
						Name:    "Event C",
						From:    date.Add(2 * time.Hour),
						To:      date.Add(2*time.Hour + 30*time.Minute),
						Genres: []dtogenre.Genre{
							{
								ID:   3,
								Name: "comedy",
							},
						},
					},
				},
			},
			resp.Data,
		)
	})
}

func TestGetEventController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("GetEventByID_ReturnsEvent", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		date := setup.Date()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/events/1", nil)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetEvent `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			response.GetEvent{
				Event: dto.Event{
					ID:      1,
					Channel: 1,
					Name:    "Event A",
					From:    date,
					To:      date.Add(15 * time.Minute),
					Genres: []dtogenre.Genre{
						{
							ID:   1,
							Name: "action",
						},
					},
				},
			},
			resp.Data,
		)
	})

	t.Run("GetEventByInvalidID_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/events/invalid", nil)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("GetNonexistentEventByID_ReturnsNotFound", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/events/69", nil)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestPostEventsController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("CreateValidEvent_ReturnsNewEvent", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		date := setup.Date()
		from := date.Add(3 * time.Hour)
		to := from.Add(2 * time.Hour)

		reqBody := fmt.Sprintf(`{
			"data": {
				"event": {
					"name": "Test Event",
					"channel": 3,
					"from": "%s",
					"to": "%s"
				}
			}
		}`, from.Format(time.RFC3339), to.Format(time.RFC3339))

		req := httptest.NewRequest(http.MethodPost, "/api/v1/events", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.PostEvents `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			response.PostEvents{
				Event: dto.Event{
					ID:      4,
					Channel: 3,
					Name:    "Test Event",
					From:    from.Truncate(time.Second),
					To:      to.Truncate(time.Second),
				},
			},
			resp.Data,
		)
	})

	t.Run("CreateInvalidEvent_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"event": {}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/events", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp common.Response

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.NotEmpty(t, resp.Error)
		assert.Contains(t, resp.Error, "failed validating")
	})
}

func TestPutEventsController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("UpdateValidEvent_ReturnsSuccess", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		date := setup.Date()
		from := date.Add(3 * time.Hour)
		to := from.Add(1 * time.Hour)

		reqBody := fmt.Sprintf(`{
			"data": {
				"event": {
					"name": "Updated Event Name",
					"channel": 2,
					"from": "%s",
					"to": "%s"
				}
			}
		}`, from.Format(time.RFC3339), to.Format(time.RFC3339))

		putReq := httptest.NewRequest(http.MethodPut, "/api/v1/events/1", bytes.NewBufferString(reqBody))
		putReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		putReq.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		putRec := httptest.NewRecorder()

		d.Router.ServeHTTP(putRec, putReq)

		assert.Equal(t, http.StatusOK, putRec.Code)

		getReq := httptest.NewRequest(http.MethodGet, "/api/v1/events/1", nil)
		getReq.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		getRec := httptest.NewRecorder()

		d.Router.ServeHTTP(getRec, getReq)

		assert.Equal(t, http.StatusOK, getRec.Code)

		var getResp struct {
			Data response.GetEvent `json:"data"`
		}

		err := json.Unmarshal(getRec.Body.Bytes(), &getResp)
		require.NoError(t, err)

		assert.Equal(t,
			response.GetEvent{
				Event: dto.Event{
					ID:      1,
					Channel: 2,
					Name:    "Updated Event Name",
					From:    from.Truncate(time.Second),
					To:      to.Truncate(time.Second),
				},
			},
			getResp.Data,
		)
	})

	t.Run("UpdateInvalidEvent_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"event": {}
			}
		}`

		req := httptest.NewRequest(http.MethodPut, "/api/v1/events/1", bytes.NewBufferString(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(auth.HeaderAPIKey, config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp common.Response

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.NotEmpty(t, resp.Error)
		assert.Contains(t, resp.Error, "failed validating")
	})
}
