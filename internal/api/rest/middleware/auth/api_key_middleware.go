package auth

import (
	"errors"
	"net/http"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ErrInvalidAPIKeyProvided = errors.New("provided API key is invalid")

// APIKey represents API key authentication middleware.
type APIKey struct {
	config config.Auth
}

func NewAPIKey(config config.Auth) APIKey {
	return APIKey{
		config: config,
	}
}

func (ak *APIKey) KeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(
		middleware.KeyAuthConfig{
			KeyLookup:    "header:X-Api-Key",
			Validator:    ak.validator,
			ErrorHandler: ak.errorHandler,
		},
	)
}

func (ak *APIKey) validator(key string, c echo.Context) (bool, error) {
	if key != ak.config.APIKey {
		return false, ErrInvalidAPIKeyProvided
	}

	return true, nil
}

func (ak *APIKey) errorHandler(err error, c echo.Context) error {
	if errors.Is(err, ErrInvalidAPIKeyProvided) {
		_ = c.JSON(http.StatusUnauthorized,
			common.NewErrorResponse(err),
		)

		return nil
	}

	_ = c.JSON(http.StatusBadRequest,
		common.NewErrorResponse(err),
	)

	return nil
}
