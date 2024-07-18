package service

import (
	"time"

	"github.com/SasangaME/goCashFlowApi/pkg/database"
	"github.com/SasangaME/goCashFlowApi/pkg/model/constants"
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/model/entity"
	"github.com/google/uuid"
)

func UserFindAll() ([]entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	var users []entity.User
	db.Find(&users)
	return users, dto.ApplicationResponse{
		IsError: false,
	}
}

func UserFindById(id string) (entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	var user entity.User
	db.Find(&user, "id = ?", id)
	if user.Id == uuid.Nil {
		return user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.NotFound,
			ErrorMessage: "user not found for id: " + id,
		}
	}
	return user, dto.ApplicationResponse{
		IsError: false,
	}
}

func UserCreate(user *entity.User) (*entity.User, dto.ApplicationResponse) {
	db := database.Database.Db
	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	err := db.Create(&user).Error
	if err != nil {
		return user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return user, dto.ApplicationResponse{
		IsError: false,
	}
}

func UserUpdate(id string, request *entity.User) (*entity.User, dto.ApplicationResponse) {
	user, _ := UserFindById(id)

	db := database.Database.Db
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.UpdatedAt = time.Now()
	err := db.Save(&user).Error
	if err != nil {
		return &user, dto.ApplicationResponse{
			IsError:      true,
			StatusCode:   constants.InternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return &user, dto.ApplicationResponse{
		IsError: false,
	}
}
