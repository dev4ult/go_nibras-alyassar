package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"

	constant "praktikum/constants"
	controller "praktikum/controllers"
	m "praktikum/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	m.Logger(e)
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.POST("/login", controller.Login)

	usersRoute := e.Group("/users")
	// usersRoute.Use(mid.BasicAuth(m.ImplementAuth))
	usersRoute.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	usersRoute.GET("", controller.GetUsers)
	usersRoute.POST("", controller.CreateUser)

	return e
}