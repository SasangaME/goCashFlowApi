package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/gofiber/fiber/v2"
)

func idValidation(c *fiber.Ctx) (string, error) {
	id := c.Params("id")
	if id == "" {
		return "", HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "id is required",
		})
	}
	return id, nil
}
