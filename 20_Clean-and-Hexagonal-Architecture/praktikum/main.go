package main

import (
	"github.com/labstack/echo/v4"

	route "clean_arch/routes"

	"clean_arch/features/users"
)

var (
	userHandler = users.UserRouteHandler()
)

func main() {

	e := echo.New()

	route.UserRoutes(e, userHandler)

	e.Start(":8000")
}