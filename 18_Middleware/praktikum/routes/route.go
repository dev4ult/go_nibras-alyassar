package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	m "praktikum/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	m.Logger(e)
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

	return e
}