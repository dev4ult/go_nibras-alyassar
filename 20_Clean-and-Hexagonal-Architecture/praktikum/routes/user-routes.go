package routes

import (
	"clean_arch/features/users/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, handler handler.IUserController) {
	users := e.Group("/users")

	users.GET("", handler.GetAllUsers())
	users.POST("", handler.CreateUser())
}