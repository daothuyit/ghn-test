package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const MDAuthorization = "Authorization"
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjQyNDkzNTF9.B-QMXFtNDckC3j6sMwbhUEnt23ER5pV51tCBd5nmAHk"

func Authorization() echo.MiddlewareFunc {
	// Add unary interceptor
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			request := c.Request()
			ctx := c.Request().Context()
			if requestToken := request.Header.Get(MDAuthorization); requestToken != token {
				return c.JSON(http.StatusUnauthorized, &echo.Map{"error": "Unauthorizedddddd"})
			}

			c.SetRequest(request.WithContext(ctx))
			// execute handler
			err = next(c)
			return
		}
	}
}
