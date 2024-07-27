package util

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateJWT(username string) (string, error) {
	jwToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwToken.Claims.(jwt.MapClaims)
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	jwtSecret := []byte("secretkey")
	jwTokenString, err := jwToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return jwTokenString, nil
}
