package helper

import "github.com/golang-jwt/jwt"

func ExtractToken(t interface{}) int {
	user := t.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userID"].(float64)
		return int(userId)
	}
	return -1
}
