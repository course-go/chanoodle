package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const HeaderAPIKey = "X-Api-Key" //nolint: gosec

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
			KeyLookup:    "header:" + HeaderAPIKey,
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
		err = c.JSON(http.StatusUnauthorized,
			common.NewErrorResponse(err),
		)
		if err != nil {
			return fmt.Errorf("failed sending response: %w", err)
		}

		return nil
	}

	err = c.JSON(http.StatusBadRequest,
		common.NewErrorResponse(err),
	)
	if err != nil {
		return fmt.Errorf("failed sending response: %w", err)
	}

	return nil
}
