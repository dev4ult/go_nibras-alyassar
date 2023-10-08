package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)


func GenerateToken(signKey string, userId int) map[string]any {
	claims, refClaims := jwt.MapClaims{}, jwt.MapClaims{}

	// access token
	claims["id"] = userId
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := sign.SignedString([]byte(signKey))

	// refresh token
	refClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	refSign := jwt.NewWithClaims(jwt.SigningMethodHS256, refClaims)
	refreshToken, errToken := refSign.SignedString([]byte(signKey))

	if err != nil || errToken != nil {
		return nil
	}

	return map[string]any {
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}
}