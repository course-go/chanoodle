package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/dto"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/response"
	dtogenre "github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/test/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetChannelsController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("GetAllChannels_ReturnsChannels", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/channels?limit=100", nil)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetChannels `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			response.GetChannels{
				Channels: []dto.Channel{
					{
						ID:       1,
						Name:     "CT24",
						Priority: 100,
						Genres: []dtogenre.Genre{
							{
								ID:   1,
								Name: "action",
							},
						},
					},
					{
						ID:       2,
						Name:     "BBC One",
						Priority: 50,
						Genres: []dtogenre.Genre{
							{
								ID:   2,
								Name: "romance",
							},
						},
					},
					{
						ID:       3,
						Name:     "Šlágr",
						Priority: 30,
						Genres: []dtogenre.Genre{
							{
								ID:   3,
								Name: "comedy",
							},
						},
					},
					{
						ID:       4,
						Name:     "LeoTV",
						Priority: 120,
						Genres: []dtogenre.Genre{
							{
								ID:   4,
								Name: "drama",
							},
						},
					},
				},
			},
			resp.Data,
		)
	})

	t.Run("GetLimitedChannelsWithOffset_ReturnsChannels", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/channels?limit=2&offset=1", nil)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetChannels `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data.Channels, 2)
		assert.Equal(t,
			response.GetChannels{
				Channels: []dto.Channel{
					{
						ID:       2,
						Name:     "BBC One",
						Priority: 50,
						Genres: []dtogenre.Genre{
							{
								ID:   2,
								Name: "romance",
							},
						},
					},
					{
						ID:       3,
						Name:     "Šlágr",
						Priority: 30,
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

func TestGetChannelController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("GetChannelByID_ReturnsChannel", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/channels/1", nil)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetChannel `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			response.GetChannel{
				Channel: dto.Channel{
					ID:       1,
					Name:     "CT24",
					Priority: 100,
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

	t.Run("GetChannelByInvalidID_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/channels/invalid", nil)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestPostChannelsController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("CreateValidChannel_ReturnsNewChannel", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"channel": {
					"name": "BBC Two"
				}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/channels", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.PostChannels `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			response.PostChannels{
				Channel: dto.Channel{
					ID:       5,
					Name:     "BBC Two",
					Priority: 100,
				},
			},
			resp.Data,
		)
	})

	t.Run("CreateInvalidChannel_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"channel": {}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/channels", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

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

func TestPutChannelsController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("UpdateValidChannel_ReturnsSuccess", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"channel": {
					"name": "CT1",
					"priority": 80
				}
			}
		}`

		putReq := httptest.NewRequest(http.MethodPut, "/api/v1/channels/1", bytes.NewBufferString(reqBody))
		putReq.Header.Set("Content-Type", "application/json")
		putReq.Header.Set("X-Api-Key", config.Auth.APIKey)

		putRec := httptest.NewRecorder()

		d.Router.ServeHTTP(putRec, putReq)

		assert.Equal(t, http.StatusOK, putRec.Code)

		getReq := httptest.NewRequest(http.MethodGet, "/api/v1/channels/1", nil)
		getReq.Header.Set("X-Api-Key", config.Auth.APIKey)

		getRec := httptest.NewRecorder()

		d.Router.ServeHTTP(getRec, getReq)

		assert.Equal(t, http.StatusOK, getRec.Code)

		var getResp struct {
			Data response.GetChannel `json:"data"`
		}

		err := json.Unmarshal(getRec.Body.Bytes(), &getResp)
		require.NoError(t, err)

		assert.Equal(t,
			response.GetChannel{
				Channel: dto.Channel{
					ID:       1,
					Name:     "CT1",
					Priority: 80,
				},
			},
			getResp.Data,
		)
	})

	t.Run("UpdateInvalidChannel_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"channel": {}
			}
		}`

		req := httptest.NewRequest(http.MethodPut, "/api/v1/channels/1", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

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
