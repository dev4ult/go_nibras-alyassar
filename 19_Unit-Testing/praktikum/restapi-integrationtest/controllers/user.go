package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	config "praktikum/config"
	mid "praktikum/middlewares"
	model "praktikum/models"
	util "praktikum/utils"
)

func FindUser(paramId string) map[string]interface{} {
	var user model.User

	userId, err := strconv.Atoi(paramId)

	if err != nil {
		return util.Response(400, "Bad Request!")
	}

	result := config.DB.First(&user, userId)

	if result.RowsAffected < 1 {
		return util.Response(404, "Not Found!")
	}

	return map[string]interface{} {
		"status": 200,
		"user": user,
		"id": userId,
	}
}


func GetUsers(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users Listed!",
		"users": users,
	})
}

func CreateUser(c echo.Context) error {
	var user model.User

	c.Bind(&user)

	err := config.DB.Create(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users Created!",
		"user": user,
	})
}

func GetUser(ctx echo.Context) error {
	user := FindUser(ctx.Param("id"))

	if user["status"] != http.StatusOK {
		return ctx.JSON(user["status"].(int), user)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"message": "User Found!",
		"user": user["user"],
	})
}

func UpdateUser(ctx echo.Context) error {
	// check user if exist
	user := FindUser(ctx.Param("id"))

	if user["status"] != http.StatusOK {
		return ctx.JSON(user["status"].(int), user)
	}

	var newUserData model.User

	ctx.Bind(&newUserData)

	err := config.DB.Table("users").Where("id", user["id"]).Updates(newUserData).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "User Updated!"))
}

func DeleteUser(ctx echo.Context) error {
	// check user if exist
	user := FindUser(ctx.Param("id"))

	if user["status"] != http.StatusOK {
		return ctx.JSON(user["status"].(int), user)
	}

	err := config.DB.Delete(&model.User{}, user["id"]).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "User Deleted!"))
}

func Login(c echo.Context) error {
	var user model.User
	c.Bind(&user)

	err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := mid.CreateToken(user.Id, user.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	userResponse := model.UserResponse{Username: user.Username, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Success!",
		"user": userResponse,
	})
}