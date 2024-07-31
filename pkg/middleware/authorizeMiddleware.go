package middleware

import (
	"github.com/SasangaME/goCashFlowApi/pkg/handler"
	"github.com/SasangaME/goCashFlowApi/pkg/service"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthorizeMiddleware(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizeHeader := c.Get("Authorization")
		if authorizeHeader == "" {
			return handler.HandleUnauthorizedError(c)
		}

		jwtString := strings.Split(authorizeHeader, " ")[1]
		if jwtString == "" {
			return handler.HandleUnauthorizedError(c)
		}

		authResponse := service.Authorize(jwtString, roles)
		if authResponse.IsError {
			return handler.HandleUnauthorizedError(c)
		}

		return c.Next()
	}
}
