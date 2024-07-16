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
		return handleException(c, exception.ApplicationError{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		})
	}
	response := mapping.RoleListToDto(roles)
	return c.Status(constants.Success).JSON(response)
}

func RoleGetById(c *fiber.Ctx) error {
	id, _ := idValidation(c)
	role, err := service.RoleFindById(id)
	if err.IsError {
		return handleException(c, err)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Success).JSON(response)
}

func RoleCreate(c *fiber.Ctx) error {
	request := new(dto.RoleDto)
	err := c.BodyParser(request)
	if err != nil {
		return handleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "request body parse error",
		})
	}
	role := mapping.DtoToRole(request)
	if err != nil {
		return handleException(c, exception.ApplicationError{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		})
	}

	serviceErr := service.RoleCreate(&role)
	if serviceErr.IsError {
		return handleException(c, serviceErr)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Created).JSON(response)
}

func RoleUpdate(c *fiber.Ctx) error {
	id, _ := idValidation(c)

	request := new(dto.RoleDto)
	err := c.BodyParser(request)
	if err != nil {
		return handleException(c, exception.ApplicationError{
			StatusCode:   constants.BabRequest,
			ErrorMessage: "request body parse error",
		})
	}
	role := mapping.DtoToRole(request)
	_, serviceErr := service.RoleUpdate(id, &role)
	if serviceErr.IsError {
		return handleException(c, serviceErr)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Success).JSON(response)
}
