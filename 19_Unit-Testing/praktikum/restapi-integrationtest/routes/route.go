package routes

import (
	"os"

	jwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"

	controller "praktikum/controllers"
)

func UserRoutes(e *echo.Echo, c controller.IUserController) {
	e.POST("/users", c.CreateUser())
	e.POST("/users/login", c.Login())

	users := e.Group("/users")
	users.Use(jwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	users.GET("", c.GetUsers())
	users.GET("/:id", c.GetUser())
	users.PUT("/:id", c.EditUser())
	users.DELETE("/:id", c.RemoveUser())
}

func BookRoutes(e *echo.Echo, c controller.IBookController) {
	books := e.Group("/books")
	books.Use(jwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	books.POST("", c.CreateBook())
	books.GET("/:id", c.GetBook())
	books.PUT("/:id", c.EditBook())
	books.DELETE("/:id", c.RemoveBook())
}