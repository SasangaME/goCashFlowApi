package service

import (
	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/google/uuid"
)

func RoleFindAll() ([]entity.Role, error) {
	db := database.Database.Db
	var roles []entity.Role
	db.Find(&roles)
	return roles, nil
}

func RoleFindById(id uuid.UUID) (entity.Role, error) {
	db := database.Database.Db
	var role entity.Role
	db.Find(&role, "id = ?", id)
	if role.Id == uuid.Nil {
		// throw an exception
	}
	return role, nil
}
