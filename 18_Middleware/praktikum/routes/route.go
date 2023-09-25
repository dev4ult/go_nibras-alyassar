package routes

import (
	jwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"

	constant "praktikum/constants"
	controller "praktikum/controllers"
	m "praktikum/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	m.Logger(e)
	e.POST("/users", controller.CreateUser)
	e.POST("/users/login", controller.Login)

	users := e.Group("/users")
	users.Use(jwt.JWT([]byte(constant.SECRET_JWT)))
	users.GET("", controller.GetUsers)
	users.GET("/:id", controller.GetUser)
	users.PUT("/:id", controller.UpdateUser)
	users.DELETE("/:id", controller.DeleteUser)

	books := e.Group("/books")
	books.Use(jwt.JWT([]byte(constant.SECRET_JWT)))
	books.POST("", controller.CreateBook)
	books.GET("/:id", controller.GetBook)
	books.PUT("/:id", controller.UpdateBook)
	books.DELETE("/:id", controller.DeleteBook)

	return e
}