package service

import (
	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/google/uuid"
	"time"
)

func RoleFindAll() ([]entity.Role, error) {
	db := database.Database.Db
	var roles []entity.Role
	db.Find(&roles)
	return roles, nil
}

func RoleFindById(id string) (entity.Role, exception.ApplicationStatus) {
	db := database.Database.Db
	var role entity.Role
	db.Find(&role, "id = ?", id)
	if role.Id == uuid.Nil {
		return entity.Role{}, exception.ApplicationStatus{
			StatusCode:   constants.NotFound,
			ErrorMessage: "role not found for id: " + id,
		}
	}
	return role, exception.ApplicationStatus{
		StatusCode: constants.Success,
	}
}

func RoleCreate(role *entity.Role) (*entity.Role, exception.ApplicationStatus) {
	db := database.Database.Db
	role.Id = uuid.New()
	role.CreatedAt = time.Now()
	err := db.Create(&role).Error
	if err != nil {
		return role, exception.ApplicationStatus{
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return role, exception.ApplicationStatus{
		StatusCode: constants.Created,
	}
}

func RoleUpdate(id string, request *entity.Role) (*entity.Role, exception.ApplicationStatus) {
	role, err := RoleFindById(id)
	if err.StatusCode >= constants.Error {
		return &role, err
	}

	db := database.Database.Db
	role.Name = request.Name
	role.Description = request.Description
	role.UpdatedAt = time.Now()
	db.Save(&role)
	return &role, exception.ApplicationStatus{
		StatusCode: constants.Success,
	}
}
