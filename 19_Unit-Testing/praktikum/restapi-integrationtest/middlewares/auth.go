package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	config "praktikum/config"
	model "praktikum/models"
)

func ImplementAuth(username string, password string, e echo.Context) (bool, error) {
	var user model.User

	err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func CreateToken(userId int, username string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}