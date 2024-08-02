package util

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func CreateJWT(username string) (string, error) {
	jwtCfg, err := getJwtConfig()
	if err != nil {
		return "", err
	}

	jwToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwToken.Claims.(jwt.MapClaims)
	claims["sub"] = username
	claims["exp"] = time.Now().Add(time.Duration(jwtCfg.expiry) * time.Hour).Unix()
	jwtSecret := []byte(jwtCfg.secret)
	jwTokenString, err := jwToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return jwTokenString, nil
}

func ValidateJWT(jwTokenString string) (string, error) {
	jwtCfg, err := getJwtConfig()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(jwTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtCfg.secret), nil
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

type jwtConfig struct {
	secret string
	expiry int
}

func getJwtConfig() (*jwtConfig, error) {
	cfg := new(jwtConfig)
	cfg.secret, _ = Config("JWT_SECRET")
	expiryStr, err := Config("JWT_EXPIRY")
	if err != nil {
		return cfg, errors.New("JWT expiry value not found")
	}
	expiry, err := strconv.Atoi(expiryStr)
	if err != nil {
		return cfg, errors.New("JWT expired")
	}
	cfg.expiry = expiry
	return cfg, nil
}
