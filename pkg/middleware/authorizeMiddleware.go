package middleware

import (
	"strings"

	"github.com/SasangaME/goCashFlowApi/pkg/handler"
	"github.com/SasangaME/goCashFlowApi/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func Authorize(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorizeHeader := c.Get("Authorization")
		if authorizeHeader == "" {
			return handler.HandleUnauthorizedError(c)
		}

		jwtString := strings.Split(authorizeHeader, " ")[1]
		if jwtString == "" {
			return handler.HandleUnauthorizedError(c)
		}

		userId, authResponse := service.Authorize(jwtString, roles)
		if authResponse.IsError {
			return handler.HandleUnauthorizedError(c)
		}

		c.Locals("userId", userId)

		return c.Next()
	}
}
