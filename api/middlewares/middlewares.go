package middleware

import (
	"net/http"

	"github.com/JGurus/template-initial-api/auth"
	"github.com/labstack/echo/v4"
)

//Auth .
func Auth(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		_, err := auth.ValidatedToken(token)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "no permitido"})
		}
		return f(c)
	}
}
