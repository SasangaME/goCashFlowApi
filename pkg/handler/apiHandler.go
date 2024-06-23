package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/SasangaME/goCashFlowApi/pkg/model/mapping"
	"github.com/SasangaME/goCashFlowApi/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func RoleGetAll(c *fiber.Ctx) error {
	roles, err := service.RoleFindAll()
	if err != nil {
		_ = HandleException(c, exception.ApplicationError{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		})
	}
	response := mapping.RoleListToDto(roles)
	return c.Status(constants.Succes).JSON(response)
}

func RoleGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	role, err := service.RoleFindById(id)
	if err.StatusCode > 0 {
		_ = HandleException(c, err)
	}
	response := mapping.RoleToDto(role)
	return c.Status(constants.Succes).JSON(response)
}
