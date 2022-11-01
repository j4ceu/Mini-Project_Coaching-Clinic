package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UnauthorizedRole(role []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("token").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			for _, v := range role {
				if claims["role"] == v {
					return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
				}
			}

			return next(c)
		}
	}
}
