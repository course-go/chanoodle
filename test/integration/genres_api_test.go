package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/response"
	"github.com/course-go/chanoodle/test/setup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGenresController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("GetAllGenres_ReturnsTwoGenres", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/genres", nil)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetGenres `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data.Genres, 5)
		assert.Equal(t,
			response.GetGenres{
				Genres: []dto.Genre{
					{ID: 1, Name: "action"},
					{ID: 2, Name: "romance"},
					{ID: 3, Name: "comedy"},
					{ID: 4, Name: "drama"},
					{ID: 5, Name: "thriller"},
				},
			},
			resp.Data,
		)
	})

	t.Run("GetLimitedGenresWithOffset_ReturnsTwoGenres", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/genres?limit=2&offset=2", nil)
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.GetGenres `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Len(t, resp.Data.Genres, 2)
		assert.Equal(t,
			response.GetGenres{
				Genres: []dto.Genre{
					{ID: 3, Name: "comedy"},
					{ID: 4, Name: "drama"},
				},
			},
			resp.Data,
		)
	})
}

func TestPostGenresController(t *testing.T) {
	t.Parallel()

	config := setup.Config(t)

	t.Run("CreateValidGenre_ReturnsExistingGenre", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"genre": {
					"name": "drama"
				}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/genres", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.PostGenres `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			dto.Genre{
				ID:   4,
				Name: "drama",
			},
			resp.Data.Genre,
		)
	})

	t.Run("CreateValidGenre_ReturnsNewGenre", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"genre": {
					"name": "psychological"
				}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/genres", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Api-Key", config.Auth.APIKey)

		rec := httptest.NewRecorder()

		d.Router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Data response.PostGenres `json:"data"`
		}

		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)

		assert.Equal(t,
			dto.Genre{
				ID:   6,
				Name: "psychological",
			},
			resp.Data.Genre,
		)
	})

	t.Run("CreateInvalidGenre_ReturnsBadRequest", func(t *testing.T) {
		t.Parallel()

		d := setup.NewDependencies(t, config)
		setup.Seed(t, d)

		reqBody := `{
			"data": {
				"genre": {}
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/genres", bytes.NewBufferString(reqBody))
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
