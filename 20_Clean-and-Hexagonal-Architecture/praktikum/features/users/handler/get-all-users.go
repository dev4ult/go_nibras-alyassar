package handler

import (
	helper "clean_arch/helpers"

	"github.com/labstack/echo/v4"
)

func (uc *UserController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {

		users, message := uc.service.FetchAll()

		if users == nil {
			return c.JSON(500, helper.Response(message, users))
		}

		return c.JSON(200, helper.Response(message, users))
	}
}