package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/SasangaME/goCashFlowApi/pkg/model/mapping"
	"github.com/SasangaME/goCashFlowApi/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func UserGetAll(c *fiber.Ctx) error {
	users, err := service.UserFindAll()
	if err.IsError {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		})
	}
	response := mapping.UserListToUserResponse(users)
	return c.Status(constants.Success).JSON(response)
}

func UserGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "id is required",
		})
	}

	user, err := service.UserFindById(id)
	if err.IsError {
		return HandleException(c, err)
	}
	response := mapping.UserToUserResponse(user)
	return c.Status(constants.Success).JSON(response)
}

func UserCreate(c *fiber.Ctx) error {
	request := new(dto.UserRequest)
	err := c.BodyParser(request)
	if err != nil {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "request body parse error",
		})
	}

	user := mapping.UserRequestToUser(request)
	_, serviceErr := service.UserCreate(&user)
	if serviceErr.IsError {
		return HandleException(c, serviceErr)
	}
	response := mapping.UserToUserResponse(user)
	return c.Status(constants.Success).JSON(response)
}