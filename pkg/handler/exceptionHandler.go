package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/gofiber/fiber/v2"
)

func HandleException(c *fiber.Ctx, errorDetail exception.ApplicationError) error {
	return c.Status(errorDetail.StatusCode).JSON(fiber.Map{
		"errorMessage": errorDetail.ErrorMessage,
	})
}
