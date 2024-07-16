package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/gofiber/fiber/v2"
)

func handleException(c *fiber.Ctx, errorDetail exception.ApplicationError) error {
	return c.Status(errorDetail.StatusCode).JSON(fiber.Map{
		"errorMessage": errorDetail.ErrorMessage,
	})
}

func handleBadRequest(c *fiber.Ctx, errorMessage string) {

}

func handleNotFound(c *fiber.Ctx, errorMessage string) {

}

func handleInternalServerError(c *fiber.Ctx, errorMessage string) {

}

func idValidation(c *fiber.Ctx) (string, error) {
	id := c.Params("id")
	if id == "" {
		return "", handleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "id is required",
		})
	}
	return id, nil
}
