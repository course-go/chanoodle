package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func KeyAuthMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(
		middleware.KeyAuthConfig{
			KeyLookup: "header:X-Api-Key",
			Validator: apiKeyValidator,
		},
	)
}

func apiKeyValidator(key string, c echo.Context) (bool, error) {
	return true, nil
}
