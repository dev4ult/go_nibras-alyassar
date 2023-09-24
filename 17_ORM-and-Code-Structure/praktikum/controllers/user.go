package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	config "praktikum/config"
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


func GetUsers(ctx echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, err.Error()))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users Listed!",
		"users": users,
	})
}

func CreateUser(ctx echo.Context) error {
	var user model.User

	ctx.Bind(&user)

	result := config.DB.Create(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, result.Error.Error()))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{} {
		"message": "New User Created!",
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

	result := config.DB.Table("users").Where("id", user["id"]).Updates(newUserData)

	if result.RowsAffected < 1 {
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

	result := config.DB.Delete(&model.User{}, user["id"])

	if result.RowsAffected < 1 {
		return ctx.JSON(http.StatusInternalServerError, util.Response(500, "Something Went Wrong!"))
	}

	return ctx.JSON(http.StatusOK, util.Response(200, "User Deleted!"))
}