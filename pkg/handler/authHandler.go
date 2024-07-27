package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	request := new(dto.LoginRequest)
	if err := c.BodyParser(request); err != nil {
		return handleBadRequest(c, "bad request")
	}
	response, responseStatus := service.Login(request.Username, request.Password)
	if responseStatus.IsError {
		return handleException(c, responseStatus)
	}

	return c.Status(constants.Success).JSON(fiber.Map{
		"jwt": response,
	})
}
