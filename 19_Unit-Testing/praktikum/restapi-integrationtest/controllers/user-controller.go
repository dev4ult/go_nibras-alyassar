package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	helper "praktikum/helpers"
	mid "praktikum/middlewares"
	model "praktikum/models"
)

type IUserController interface {
	GetUsers() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	EditUser() echo.HandlerFunc
	RemoveUser() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserController struct {
	model model.IUserModel
}

func NewUserController(model model.IUserModel) IUserController {
	return &UserController{
		model: model,
	}
}

func (uc *UserController) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := uc.model.SelectAllUser()
	
		if users == nil {
			return c.JSON(http.StatusInternalServerError, helper.Response(500, "Something went Wrong!"))
		}
	
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Users Listed!",
			"users": users,
		})
	}
}

func (uc *UserController) CreateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user model.User

		ctx.Bind(&user)

		result := uc.model.InsertUser(user) 

		if result == nil {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something went Wrong!"))
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Users Created!",
			"user": user,
		})
	}
}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		user := uc.model.FindUser(ctx.Param("id"))
	
		if user["status"] != http.StatusOK {
			return ctx.JSON(user["status"].(int), user)
		}
	
		return ctx.JSON(http.StatusOK, map[string]interface{} {
			"message": "User Found!",
			"user": user["user"],
		})
	}
}

func (uc *UserController) EditUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check user if exist
		user := uc.model.FindUser(ctx.Param("id"))
	
		if user["status"] != http.StatusOK {
			return ctx.JSON(user["status"].(int), user)
		}
	
		var newUser model.User
	
		ctx.Bind(&newUser)
	
		update := uc.model.UpdateUser(user["id"].(int), newUser)
	
		if !update {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
		}
	
		return ctx.JSON(http.StatusOK, helper.Response(200, "User Updated!"))
	}
}

func (uc *UserController) RemoveUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// check user if exist
		user := uc.model.FindUser(ctx.Param("id"))
	
		if user["status"] != http.StatusOK {
			return ctx.JSON(user["status"].(int), user)
		}
	
		delete := uc.model.DeleteUser(user["id"].(int))
	
		if !delete {
			return ctx.JSON(http.StatusInternalServerError, helper.Response(500, "Something Went Wrong!"))
		}
	
		return ctx.JSON(http.StatusOK, helper.Response(200, "User Deleted!"))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		c.Bind(&user)
	
		result := uc.model.FindUserAccount(user) 
	
		if result == nil {
			return c.JSON(http.StatusNotFound, helper.Response(404, "User Not Found"))
		}
	
		token, err := mid.CreateToken(user.Id, user.Username)
	
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response(500, "Something went Wrong!"))
		}
	
		userResponse := model.UserResponse{Username: user.Username, Email: user.Email, Token: token}
	
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login Success!",
			"user": userResponse,
		})
	}
}