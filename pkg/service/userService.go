package service

import (
	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/SasangaME/goCashFlowApi/pkg/model/exception"
	"github.com/google/uuid"
	"time"
)

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
