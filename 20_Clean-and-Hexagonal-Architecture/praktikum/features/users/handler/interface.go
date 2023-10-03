package handler

import (
	"github.com/labstack/echo/v4"

	service "clean_arch/features/users/service"
)

type IUserController interface {
	GetAllUsers() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
}

type UserController struct {
	service service.IUserService
}

func NewUserController (service service.IUserService) IUserController {
	return &UserController{
		service: service,
	}
}