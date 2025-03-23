package middlewares

import (
	"net/http"
	"strings"

	"github.com/emersonvalentim/drobe-api/services/auth"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authService *auth.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return validateToken(c, authService, next)
		}
	}
}

func validateToken(c echo.Context, authService *auth.Service, next echo.HandlerFunc) error {
	path := c.Path()
	if path == "/api/auth/signup" || path == "/api/auth/signin" {
		return next(c)
	}

	token, err := extractToken(c)
	if err != nil {
		return err
	}

	userID, err := authService.ValidateToken(c.Request().Context(), token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	c.Set("userID", userID)
	return next(c)
}

func extractToken(c echo.Context) (string, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "No token provided")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
	}

	return token, nil
}
