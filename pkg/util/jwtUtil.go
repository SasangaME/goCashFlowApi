package util

import (
	"errors"
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

func ValidateJWT(jwTokenString string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(jwTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token is invalid")
	}

	username := claims["sub"].(string)
	return username, nil
}
