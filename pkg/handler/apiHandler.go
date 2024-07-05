package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/SasangaME/goCashFlowApi/pkg/model/mapping"
	"github.com/SasangaME/goCashFlowApi/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func RoleGetAll(c *fiber.Ctx) error {
	roles, err := service.RoleFindAll()
	if err != nil {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		})
	}
	response := mapping.RoleListToDto(roles)
	return c.Status(constants.Success).JSON(response)
}

func RoleGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "id is required",
		})
	}
	role, err := service.RoleFindById(id)
	if err.IsError {
		return HandleException(c, err)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Success).JSON(response)
}

func RoleCreate(c *fiber.Ctx) error {
	request := new(dto.RoleDto)
	err := c.BodyParser(request)
	if err != nil {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "request body parse error",
		})
	}
	role := mapping.DtoToRole(request)
	if err != nil {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		})
	}

	_, serviceErr := service.RoleCreate(&role)
	if serviceErr.IsError {
		return HandleException(c, serviceErr)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Created).JSON(response)
}

func RoleUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "id is required",
		})
	}

	request := new(dto.RoleDto)
	err := c.BodyParser(request)
	if err != nil {
		return HandleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "request body parse error",
		})
	}
	role := mapping.DtoToRole(request)
	_, serviceErr := service.RoleUpdate(id, &role)
	if serviceErr.IsError {
		return HandleException(c, serviceErr)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Success).JSON(response)
}
