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

func RoleFindById(id string) (entity.Role, exception.ApplicationError) {
	db := database.Database.Db
	var role entity.Role
	db.Find(&role, "id = ?", id)
	if role.Id == uuid.Nil {
		return entity.Role{}, exception.ApplicationError{
			IsError:      true,
			StatusCode:   constants.NotFound,
			ErrorMessage: "role not found for id: " + id,
		}
	}
	return role, exception.ApplicationError{
		IsError: false,
	}
}

func RoleCreate(role *entity.Role) (*entity.Role, exception.ApplicationError) {
	db := database.Database.Db
	role.Id = uuid.New()
	role.CreatedAt = time.Now()
	err := db.Create(&role).Error
	if err != nil {
		return role, exception.ApplicationError{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return role, exception.ApplicationError{
		IsError: false,
	}
}

func RoleUpdate(id string, request *entity.Role) (*entity.Role, exception.ApplicationError) {
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
		return &role, exception.ApplicationError{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: dbErr.Error(),
		}
	}
	return &role, exception.ApplicationError{
		StatusCode: constants.Success,
	}
}

func UserFindAll() ([]entity.User, exception.ApplicationError) {
	db := database.Database.Db
	var users []entity.User
	db.Find(&users)
	return users, exception.ApplicationError{
		IsError: false,
	}
}

func UserFindById(id string) (entity.User, exception.ApplicationError) {
	db := database.Database.Db
	var user entity.User
	db.Find(&user, "id = ?", id)
	if user.Id == uuid.Nil {
		return user, exception.ApplicationError{
			IsError:      true,
			StatusCode:   constants.NotFound,
			ErrorMessage: "user not found for id: " + id,
		}
	}
	return user, exception.ApplicationError{
		IsError: false,
	}
}

func UserCreate(user *entity.User) (*entity.User, exception.ApplicationError) {
	db := database.Database.Db
	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	err := db.Create(&user).Error
	if err != nil {
		return user, exception.ApplicationError{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return user, exception.ApplicationError{
		IsError: false,
	}
}

func UserUpdate(id string, request *entity.User) (*entity.User, exception.ApplicationError) {
	user, _ := UserFindById(id)

	db := database.Database.Db
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.UpdatedAt = time.Now()
	err := db.Save(&user).Error
	if err != nil {
		return &user, exception.ApplicationError{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return &user, exception.ApplicationError{
		IsError: false,
	}
}
