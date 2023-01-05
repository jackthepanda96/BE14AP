package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(float64)
		return int(userId)
	}
	return -1
}
