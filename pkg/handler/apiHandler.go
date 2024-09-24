package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/gofiber/fiber/v2"
)

func handleException(c *fiber.Ctx, errorDetail dto.ApplicationResponse) error {
	return c.Status(errorDetail.StatusCode).JSON(fiber.Map{
		"errorMessage": errorDetail.ErrorMessage,
	})
}

func handleBadRequest(c *fiber.Ctx, errorMessage string) error {
	return handleException(c, dto.ApplicationResponse{
		StatusCode:   constants.BadRequest,
		ErrorMessage: errorMessage,
	})
}

func handleNotFound(c *fiber.Ctx, errorMessage string) error {
	return handleException(c, dto.ApplicationResponse{
		StatusCode:   constants.NotFound,
		ErrorMessage: errorMessage,
	})
}

func HandleInternalServerError(c *fiber.Ctx, errorMessage string) error {
	return handleException(c, dto.ApplicationResponse{
		StatusCode:   constants.InternalServerError,
		ErrorMessage: errorMessage,
	})
}

func HandleUnauthorizedError(c *fiber.Ctx) error {
	return handleException(c, dto.ApplicationResponse{
		StatusCode:   constants.NotAuthorized,
		ErrorMessage: "user not authorized",
	})
}
func validateId(c *fiber.Ctx) (string, error) {
	id := c.Params("id")
	if id == "" {
		return "", handleException(c, dto.ApplicationResponse{
			StatusCode:   constants.BadRequest,
			ErrorMessage: "id is required",
		})
	}
	return id, nil
}
