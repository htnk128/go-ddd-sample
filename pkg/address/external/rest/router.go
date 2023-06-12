package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
