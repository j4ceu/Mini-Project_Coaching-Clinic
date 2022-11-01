package middlewares

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("jace"))
}


