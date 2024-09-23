package service

import (
	"time"

	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
)

func RoleFindAll() ([]entity.Role, error) {
	db := database.Database.Db
	var roles []entity.Role
	db.Find(&roles)
	return roles, nil
}

func RoleFindById(id string) (entity.Role, dto.ApplicationResponse) {
	db := database.Database.Db
	var role entity.Role
	db.Find(&role, "id = ?", id)
	if role.Id == 0 {
		return entity.Role{}, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.NotFound,
			ErrorMessage: "role not found for id: " + id,
		}
	}
	return role, dto.ApplicationResponse{
		IsError: false,
	}
}

func RoleCreate(role *entity.Role) dto.ApplicationResponse {
	db := database.Database.Db
	role.CreatedAt = time.Now()
	err := db.Create(&role).Error
	if err != nil {
		return dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return dto.ApplicationResponse{
		IsError: false,
	}
}

func RoleUpdate(id string, request *entity.Role) (*entity.Role, dto.ApplicationResponse) {
	role, err := RoleFindById(id)
	if err.StatusCode >= constants.Error {
		return &role, err
	}

	db := database.Database.Db
	role.Name = request.Name
	role.Description = request.Description
	role.UpdatedAt = time.Now()
	dbErr := db.Save(&role).Error
	if dbErr != nil {
		return &role, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: dbErr.Error(),
		}
	}
	return &role, dto.ApplicationResponse{
		StatusCode: constants.Success,
	}
}
