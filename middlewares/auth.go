package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		auth, ok := sess.Values["authenticated"].(bool)

		if !ok || !auth {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Não autorizado",
			})
		}

		return next(c)
	}
}
