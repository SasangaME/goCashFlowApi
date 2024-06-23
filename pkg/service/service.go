package service

import (
	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/google/uuid"
)

func RoleFindAll() ([]entity.Role, error) {
	db := database.Database.Db
	var roles []entity.Role
	db.Find(&roles)
	return roles, nil
}

func RoleFindById(id string) (entity.Role, exception.ApplicationError) {
	db := database.Database.Db
	var role entity.Role
	db.Find(&role, "id = ?", id)
	if role.Id == uuid.Nil {
		return entity.Role{}, exception.ApplicationError{
			StatusCode:   constants.NotFound,
			ErrorMessage: "role not found for id: " + id,
		}
	}
	return role, exception.ApplicationError{}
}
