package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	config "praktikum/config"
	mid "praktikum/middlewares"
	model "praktikum/models"
)

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

	result := config.DB.Create(&user)

	if result.Error != nil || result.RowsAffected < 1 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": 500,
			"message": result.Error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users Created!",
		"user": user,
	})
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